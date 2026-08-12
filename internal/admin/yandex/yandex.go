package yandex

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	yandexStaticEndpoint = "https://storage.yandexcloud.net"
	yandexRegion = "ru-central1-b"
	backetName = "my-task-images"
	accessKeyID = "YCAJEwKJoLee4euAx7bBMQLMU"
	secretAccessKey = "YCM4SIqwu2lRoWnnt0spIFPG7k0zs1AaDfd8-AqO"
)

func InitS3() *s3.Client{
	creds := credentials.NewStaticCredentialsProvider(accessKeyID,
													secretAccessKey,
													"")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
					config.WithRegion(yandexRegion),
					config.WithCredentialsProvider(creds),
				)
	if err != nil{
		log.Fatal(err)
	}
	
	yse := yandexStaticEndpoint

	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = &yse
	})

	return s3Client
}

func DownLoadFile(fileName string, file multipart.File) (string, error){
	bns := backetName
	
	_, err :=  InitS3().PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bns,
		Key: &fileName,
		Body: file,
	})

	if err != nil {
		return "", err
	}

	fileURL := fmt.Sprintf("%s/%s/%s", yandexStaticEndpoint, backetName, fileName)

	return fileURL, nil
}

