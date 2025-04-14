package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"math"
	"os/exec"
	"time"
)

var (
	frameSize     = image.Pt(640, 360)
	canvasSize    = image.Pt(1280, 720)
	streams       = []string{"rtsp://cam1", "rtsp://cam2", "rtsp://cam3", "rtsp://cam4"}
	perspectiveTo = [][]image.Point{
		{{50, 50}, {600, 30}, {580, 330}, {100, 350}},
		{{680, 30}, {1230, 50}, {1180, 350}, {700, 330}},
		{{100, 370}, {580, 390}, {600, 690}, {50, 670}},
		{{700, 390}, {1180, 370}, {1230, 670}, {680, 690}},
	}
)

func main() {
	// 打开 RTSP
	caps := make([]*gocv.VideoCapture, 4)
	for i, url := range streams {
		cap, err := gocv.OpenVideoCapture(url)
		if err != nil {
			panic(fmt.Sprintf("无法打开摄像头 #%d: %v", i+1, err))
		}
		caps[i] = cap
		defer cap.Close()
	}

	// 启动 FFmpeg 推流
	ffmpeg := exec.Command("ffmpeg",
		"-f", "image2pipe",
		"-r", "25",
		"-i", "pipe:0",
		"-vcodec", "libx264",
		"-preset", "ultrafast",
		"-tune", "zerolatency",
		"-f", "rtsp",
		"-rtsp_transport", "tcp",
		"rtsp://your-server-ip:8554/merged",
	)
	stdin, _ := ffmpeg.StdinPipe()
	_ = ffmpeg.Start()

	// 初始化矩阵
	mats := make([]gocv.Mat, 4)
	warped := make([]gocv.Mat, 4)
	masks := make([]gocv.Mat, 4)
	perspectives := getPerspectiveMatrices()

	for i := 0; i < 4; i++ {
		mats[i] = gocv.NewMat()
		defer mats[i].Close()
		warped[i] = gocv.NewMat()
		defer warped[i].Close()
		masks[i] = createFeatheredMask(canvasSize.X, canvasSize.Y)
		defer masks[i].Close()
	}

	canvas := gocv.NewMatWithSize(canvasSize.Y, canvasSize.X, gocv.MatTypeCV8UC3)
	defer canvas.Close()

	// 主循环
	for {
		canvas.SetTo(gocv.NewScalar(0, 0, 0, 0)) // 清空画布

		for i := 0; i < 4; i++ {
			if ok := caps[i].Read(&mats[i]); !ok || mats[i].Empty() {
				continue
			}
			gocv.Resize(mats[i], &mats[i], frameSize, 0, 0, gocv.InterpolationDefault)
			gocv.WarpPerspective(mats[i], &warped[i], perspectives[i], canvasSize)

			// 用掩码融合图像
			temp := gocv.NewMat()
			defer temp.Close()
			warped[i].CopyToWithMask(&temp, masks[i])
			gocv.Add(canvas, temp, &canvas)
		}

		// 推送帧到 FFmpeg
		buf, _ := gocv.IMEncode(".jpg", canvas)
		stdin.Write(buf.GetBytes())
		time.Sleep(40 * time.Millisecond)
	}
}

// 获取透视变换矩阵
func getPerspectiveMatrices() []gocv.Mat {
	src := gocv.NewPointVectorFromPoints([]image.Point{
		{0, 0}, {frameSize.X, 0}, {frameSize.X, frameSize.Y}, {0, frameSize.Y},
	})
	mats := []gocv.Mat{}
	for _, dst := range perspectiveTo {
		dstVec := gocv.NewPointVectorFromPoints(dst)
		mats = append(mats, gocv.GetPerspectiveTransform(src, dstVec))
	}
	return mats
}

// 创建带羽化效果的掩码
func createFeatheredMask(width, height int) gocv.Mat {
	mask := gocv.NewMatWithSize(height, width, gocv.MatTypeCV8UC1)
	centerX := width / 2
	centerY := height / 2
	maxDist := math.Sqrt(float64(centerX*centerX + centerY*centerY))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dx := float64(x - centerX)
			dy := float64(y - centerY)
			dist := math.Sqrt(dx*dx + dy*dy)
			weight := 255 * (1 - dist/maxDist)
			if weight < 0 {
				weight = 0
			}
			mask.SetUCharAt(y, x, uint8(weight))
		}
	}
	return mask
}
