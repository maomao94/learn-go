package config

import (
	"log"

	"github.com/minio/minio-go/v6"
)

// 初始化对象存储
func Minio() *minio.Client {
	endpoint := "localhost:32139"
	accessKeyID := "AKIAIOSFODNN7EXAMPLE"
	secretAccessKey := "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	useSSL := false

	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%#v\n###success###\n", minioClient) // minioClient初使化成功
	buckets, err := minioClient.ListBuckets()
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range buckets {
		log.Printf("###bucket :%#v###\n", v)
	}
	return minioClient
}
