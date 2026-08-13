package yandex

import (
	"context"
	"course/internal/tasks/models"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct{
	conf *models.YandexConfig
}

func NewS3Client() *S3Client{
	return &S3Client{
		conf: NewConfig(),
	}
}



func(s3c *S3Client) InitS3() *s3.Client{
	creds := credentials.NewStaticCredentialsProvider(s3c.conf.AccessKeyID,
													s3c.conf.SecretAccessKey,
													"")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
					config.WithRegion(s3c.conf.YandexRegion),
					config.WithCredentialsProvider(creds),
				)
	if err != nil{
		log.Fatal(err)
	}
	

	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = &s3c.conf.YandexStaticEndpoint
	})

	return s3Client
}

func (s3c *S3Client)DownLoadFile(fileName string, file multipart.File) (string, error){
	
	
	_, err :=  s3c.InitS3().PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &s3c.conf.BacketName,
		Key: &fileName,
		Body: file,
	})

	if err != nil {
		return "", err
	}

	fileURL := fmt.Sprintf("%s/%s/%s",
	 			s3c.conf.YandexStaticEndpoint,
				s3c.conf.BacketName,
				fileName)

	return fileURL, nil
}

