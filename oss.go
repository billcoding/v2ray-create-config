package main

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

var (
	endpoint        = ""
	accessKeyId     = ""
	accessKeySecret = ""
	bucketName      = ""
	objectName      = ""
)

func init() {
	if n := os.Getenv("ENDPOINT"); n != "" {
		endpoint = n
	}
	if n := os.Getenv("ACCESS_KEY_ID"); n != "" {
		accessKeyId = n
	}
	if n := os.Getenv("ACCESS_KEY_SECRET"); n != "" {
		accessKeySecret = n
	}
	if n := os.Getenv("BUCKET_NAME"); n != "" {
		bucketName = n
	}
	if n := os.Getenv("OBJECT_NAME"); n != "" {
		objectName = n
	}
}

func uploadConfig(conf string) {
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	_, err = buf.WriteString(conf)
	if err != nil {
		panic(err)
	}
	err = bucket.PutObject(objectName, buf)
	if err != nil {
		panic(err)
	}
}
