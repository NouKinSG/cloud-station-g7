package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 修改这些变量 来控制程序运行逻辑
var (
	// 程序内置
	endpint     = "oss-cn-beijing.aliyuncs.com"
	acessKey    = "xx"
	acessSecret = "xx"

	// 默认配置
	bucketName = "xiaojin-devcloud-station"

	//用户需要传递的参数
	//期望用户自己输入（CLI/GUI）
	uploadFile = ""
)

// 实现文件上传的函数
func upload(file_path string) error {

	// 1、实例化client
	client, err := oss.New(endpint, acessKey, acessSecret)
	if err != nil {
		return err
	}

	// 2、获取bucket对象
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}
	// 3、上传文件到该bucket
	return bucket.PutObjectFromFile(file_path, file_path)
}

// 参数合法性检查
func validate() error {
	if endpint == "" || acessKey == "" || acessSecret == "" {
		return fmt.Errorf("endpint,acessKey,acessSecret is empty")
	}

	if uploadFile == "" {
		return fmt.Errorf("upload file path required")
	}

	return nil
}

func loadParams() {
	flag.StringVar(&uploadFile, "f", "", "输入上传文件的名称")
	flag.Parse()
}

func main() {

	//参数加载
	loadParams()

	//参数验证
	if err := validate(); err != nil {
		fmt.Printf("参数校验异常：%s\n", err)
	}

	if err := upload(uploadFile); err != nil {
		fmt.Printf("上传文件异常：%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("文件：%s 上传完成\n", uploadFile)
}
