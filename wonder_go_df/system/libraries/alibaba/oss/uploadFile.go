package oss

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func UploadFile(filename string) string {
	var yourEndpoint string = "http://****/"
	var readEndPoint string = "http://****/"
	var bucketStr = "****"

	// 创建OSSClient实例。
	client, err := oss.New(yourEndpoint, "api", "key")
	if err != nil {
		fmt.Println("oss new Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketStr)
	if err != nil {
		fmt.Println("oss bucket Error:", err)
		os.Exit(-1)
	}

	h := md5.New()
	h.Write([]byte(filename))
	cipherStr := h.Sum(nil)
	var objectName = hex.EncodeToString(cipherStr) + ".xlsx"

	// 上传本地文件。
	err = bucket.PutObjectFromFile(objectName, filename)
	if err != nil {
		fmt.Println("oss put Error:", err)
		os.Exit(-1)
	}

	return readEndPoint + objectName
}
