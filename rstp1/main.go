package main

import (
	"bytes"
	"image"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"gocv.io/x/gocv"
)

var frameSize = image.Pt(640, 360)

func main() {
	streams := []string{
		"rtmp://10.10.1.213:1935/avplive/test?sign=41db35390ddad33f83944f44b8b75ded",
		"rtmp://10.10.1.213:1935/avplive/test?sign=41db35390ddad33f83944f44b8b75ded",
		"rtmp://10.10.1.213:1935/avplive/test?sign=41db35390ddad33f83944f44b8b75ded",
		"rtmp://10.10.1.213:1935/avplive/test?sign=41db35390ddad33f83944f44b8b75ded",
	}

	frameChans := make([]chan gocv.Mat, len(streams))
	for i, url := range streams {
		frameChans[i] = make(chan gocv.Mat, 50) // 增加缓冲区大小
		go launchWithRetry(i, url, frameChans[i])
	}

	canvas := gocv.NewMatWithSize(720, 1280, gocv.MatTypeCV8UC3)
	defer canvas.Close()

	perspectiveMats := getPerspectiveMatrices()

	warped := make([]gocv.Mat, 4)
	for i := range warped {
		warped[i] = gocv.NewMat()
		defer warped[i].Close()
	}

	pushCmd := exec.Command("ffmpeg",
		"-f", "image2pipe",
		"-r", "25",
		"-i", "pipe:0",
		"-vcodec", "libx264",
		"-preset", "ultrafast",
		"-tune", "zerolatency",
		"-f", "tee",
		"-map", "0:v",
		"[f=rtsp|rtsp_transport=tcp]rtsp://your-server-ip:8554/merged|[f=flv]rtmp://10.10.1.213:1935/avplive/test11111?sign=41db35390ddad33f83944f44b8b75ded",
	)
	stdin, _ := pushCmd.StdinPipe()
	_ = pushCmd.Start()

	for {
		canvas.SetTo(gocv.NewScalar(0, 0, 0, 0))
		var wg sync.WaitGroup
		for i := 0; i < 4; i++ {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				select {
				case frame := <-frameChans[idx]:
					defer frame.Close()
					// 调整图像大小
					gocv.Resize(frame, &frame, frameSize, 0, 0, gocv.InterpolationDefault)
					// 进行透视变换
					gocv.WarpPerspective(frame, &warped[idx], perspectiveMats[idx], image.Pt(1280, 720))
				case <-time.After(200 * time.Millisecond): // 超时等待
					log.Printf("[cam %d] no frame available within timeout", idx)
				}
			}(i)
		}
		wg.Wait()

		for i := 0; i < 4; i++ {
			gocv.AddWeighted(canvas, 1.0, warped[i], 0.5, 0.0, &canvas)
		}

		buf, _ := gocv.IMEncode(".jpg", canvas)
		stdin.Write(buf.GetBytes())
		time.Sleep(40 * time.Millisecond)
	}
}

func launchWithRetry(camIndex int, url string, frameChan chan<- gocv.Mat) {
	for {
		cmd, pipe, err := startFFmpegPipe(url)
		if err != nil {
			log.Printf("[cam %d] failed to start FFmpeg: %v", camIndex, err)
			time.Sleep(3 * time.Second)
			continue
		}

		done := make(chan struct{})
		go func() {
			err := cmd.Wait()
			if err != nil {
				log.Printf("[cam %d] FFmpeg exited with error: %v", camIndex, err)
			} else {
				log.Printf("[cam %d] FFmpeg exited normally", camIndex)
			}
			close(done)
		}()

		readMJPEGStream(camIndex, pipe, frameChan)
		<-done
		log.Printf("[cam %d] restarting FFmpeg in 3s...", camIndex)
		time.Sleep(3 * time.Second)
	}
}

func startFFmpegPipe(streamURL string) (*exec.Cmd, io.ReadCloser, error) {
	inputArgs := []string{"-fflags", "nobuffer", "-flags", "low_delay", "-an", "-i", streamURL}
	if strings.HasPrefix(streamURL, "rtsp") {
		inputArgs = append([]string{"-rtsp_transport", "tcp", "-stimeout", "5000000"}, inputArgs...)
	}

	cmd := exec.Command("ffmpeg",
		append(inputArgs,
			"-f", "image2pipe",
			"-vcodec", "mjpeg",
			"-q:v", "5",
			"pipe:1",
		)...,
	)
	cmd.Stderr = os.Stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, nil, err
	}
	return cmd, stdout, nil
}

func readMJPEGStream(camIndex int, pipe io.Reader, frameChan chan<- gocv.Mat) {
	buf := make([]byte, 0)
	tmp := make([]byte, 4096)
	start := []byte{0xFF, 0xD8}
	end := []byte{0xFF, 0xD9}

	for {
		n, err := pipe.Read(tmp)
		if err != nil {
			if err == io.EOF {
				log.Printf("[cam %d] stream ended (EOF)", camIndex)
			} else {
				log.Printf("[cam %d] pipe read error: %v", camIndex, err)
			}
			break
		}
		log.Printf("[cam %d] Read %d bytes", camIndex, n) // 增加日志，确认是否读取到数据
		buf = append(buf, tmp[:n]...)

		for {
			s := bytes.Index(buf, start)
			e := bytes.Index(buf, end)
			if s >= 0 && e > s {
				e += 2
				img := buf[s:e]
				buf = buf[e:]
				mat, err := gocv.IMDecode(img, gocv.IMReadColor)
				if err == nil && !mat.Empty() {
					frameChan <- mat
				} else {
					log.Printf("[cam %d] failed to decode JPEG", camIndex)
				}
			} else {
				break
			}
		}
	}
}

func getPerspectiveMatrices() []gocv.Mat {
	src := gocv.NewPointVectorFromPoints([]image.Point{{0, 0}, {640, 0}, {640, 360}, {0, 360}})
	return []gocv.Mat{
		gocv.GetPerspectiveTransform(src, gocv.NewPointVectorFromPoints([]image.Point{{50, 50}, {600, 30}, {580, 330}, {100, 350}})),
		gocv.GetPerspectiveTransform(src, gocv.NewPointVectorFromPoints([]image.Point{{680, 30}, {1230, 50}, {1180, 350}, {700, 330}})),
		gocv.GetPerspectiveTransform(src, gocv.NewPointVectorFromPoints([]image.Point{{100, 370}, {580, 390}, {600, 690}, {50, 670}})),
		gocv.GetPerspectiveTransform(src, gocv.NewPointVectorFromPoints([]image.Point{{700, 390}, {1180, 370}, {1230, 670}, {680, 690}})),
	}
}
