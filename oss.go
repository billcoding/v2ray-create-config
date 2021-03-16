package main

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	endpoint        = "http://oss-cn-hangzhou.aliyuncs.com"
	accessKeyId     = "LTAI4Fq3jJLvC1K7jfKWmrKM"
	accessKeySecret = "a4TAV7wUDknhEgQcRgfsvQRHcxjHuY"
	bucketName      = "v2ray-files"
)

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
	err = bucket.PutObject("custom-config/config-20200824.json", buf)
	if err != nil {
		panic(err)
	}
}
