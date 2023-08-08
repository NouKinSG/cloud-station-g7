package example_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	// 全局client实例，在包加载的时候初始化(init)
	client *oss.Client
)

var (
	AccessKey       = os.Getenv("ALI_AK")
	AccessKeySecret = os.Getenv("ALI_SK")
	OssEndpoint     = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName      = os.Getenv("ALI_BUCET_NAME")
)

// 测试阿里云OssSDK BucketList接口
func TestBucketList(t *testing.T) {

	lsRes, err := client.ListBuckets()
	if err != nil {
		// HandleError(err)
		t.Log(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}

// 测试阿里云OssSDK PutObjectFromFile接口
func TestUploadFile(t *testing.T) {

	bucket, err := client.Bucket("my-bucket")
	if err != nil {
		// HandleError(err)
		t.Log(err)
	}

	err = bucket.PutObjectFromFile("my-object", "LocalFile")
	if err != nil {
		// HandleError(err)
		t.Log(err)
	}
}

// 初始化一个Oss Client，等下给其他所有测试用例使用
func init() {

	c, err := oss.New(OssEndpoint, AccessKey, AccessKeySecret)

	if err != nil {
		// HandleError(err)
		panic(err)
	}
	client = c
}
