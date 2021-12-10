package handler

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"strings"
)


func HandleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func URL() {
	// 获取STS临时凭证后，您可以通过其中的安全令牌（SecurityToken）和临时访问密钥（AccessKeyId和AccessKeySecret）生成OSSClient。
	client, err := oss.New("yourEndpoint", "yourAccessKeyId", "yourAccessKeySecret", oss.SecurityToken("yourSecurityToken"))
	if err != nil {
		HandleError(err)
	}
	// 填写Bucket名称，例如examplebucket。
	bucketName := "examplebucket"
	// 填写文件完整路径，例如exampledir/exampleobject.txt。文件完整路径中不能包含Bucket名称。
	objectName := "exampledir/exampleobject.txt"
	// 填写本地文件完整路径，例如D:\\localpath\\examplefile.txt，其中localpath为本地文件examplefile.txt所在本地路径。
	localFilename := "D:\\localpath\\examplefile.txt"

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// 签名直传。1
	signedURL, err := bucket.SignURL(objectName, oss.HTTPPut, 60)
	if err != nil {
		HandleError(err)
	}

	var val = "上云就上阿里云"
	err = bucket.PutObjectWithURL(signedURL, strings.NewReader(val))
	if err != nil {
		HandleError(err)
	}

	// 带可选参数的签名直传。请确保设置的ContentType值与在前端使用时设置的ContentType值一致。
	options := []oss.Option{
		oss.Meta("myprop", "mypropval"),
		oss.ContentType("text/plain"),
	}

	signedURL, err = bucket.SignURL(objectName, oss.HTTPPut, 60, options...)
	if err != nil {
		HandleError(err)
	}

	err = bucket.PutObjectFromFileWithURL(signedURL, localFilename, options...)
	if err != nil {
		HandleError(err)
	}
}

