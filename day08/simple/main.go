package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	endpint        = "oss-cn-chengdu.aliyuncs.com"
	ak             = "LTAI5tMvG5NA51eiH3ENZDaa"
	sk             = "vWOGbrPKQGmLVo4CKSgmAB62vdum10"
	bucketName     = "cloud-station"
	uploadFilePath = "go.mod"
)

var (
	help bool
)

func main() {
	// 读取用户输入的参数
	loadParam()

	// 校验程序的入参
	if err := validate(); err != nil {
		fmt.Printf("validate param error, %s\n", err)
		os.Exit(1)
	}

	// 执行文件上传
	if err := upload(uploadFilePath); err != nil {
		fmt.Printf("upload file error, %s\n", err)
		os.Exit(1)
	}

	// 执行成功，反馈
	fmt.Printf("upload file: %s success\n", uploadFilePath)
}

func loadParam() error {
	flag.Parse()

	if help {
		usage()
		os.Exit(0)
	}

	return nil
}

func validate() error {
	if endpint == "" {
		return fmt.Errorf("endpoint is missed")
	}
	if ak == "" || sk == "" {
		return fmt.Errorf("ak or sk is missed")
	}
	if uploadFilePath == "" {
		return fmt.Errorf("upload file missed")
	}
	return nil
}

func upload(filePath string) error {
	client, err := oss.New(endpint, ak, sk)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 第一个参数: 上传到oss里面的文件的key(路径), go.sum --->  2021/7/21/go.sum
	// 第二个参数: 需要上传的文件的路径
	err = bucket.PutObjectFromFile(filePath, filePath)
	if err != nil {
		return err
	}

	// 打印下载URL
	// sts, 临时授权token(有效期1天)
	signedURL, err := bucket.SignURL(filePath, oss.HTTPGet, 60*60*24)
	if err != nil {
		return fmt.Errorf("sign file download url error, %s", err)
	}
	fmt.Printf("下载链接: %s\n", signedURL)
	fmt.Println("\n注意: 文件下载有效期为1天, 中转站保存时间为3天, 请及时下载")
	return nil
}

func usage() {
	fmt.Fprintf(os.Stderr, `cloud-station version: 0.0.1
Usage: cloud-station [-h] -f <uplaod_file_path>
Options:
`)
	flag.PrintDefaults()
}

func init() {
	flag.BoolVar(&help, "h", false, "help usage")
	flag.StringVar(&uploadFilePath, "f", "", "upload file path")
}
