package main

import (
	"fmt"
	"learn-go/minio/config"
	"log"

	"github.com/minio/minio-go/v6"
)

func main() {
	minioClient := config.Minio()

	// 创建一个叫mymusic的存储桶。
	bucketName := "mymusic"
	location := "us-east-2"

	err := minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)

	// 上传一个文件。
	objectName := "2/golden-oldies.zip"
	filePath := "./tmp/golden-oldies.zip"
	contentType := "application/zip"

	// 使用FPutObject上传一个zip文件。
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	//err = minioClient.RemoveObject(bucketName, objectName)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	doneCh := make(chan struct{})
	defer close(doneCh)
	objectCh := minioClient.ListObjectsV2(bucketName, "", false, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Println(object)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}
