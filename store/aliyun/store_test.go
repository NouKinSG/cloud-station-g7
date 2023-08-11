package aliyun_test

import (
	"os"
	"testing"

	"github.com/NouKinSG/cloud-station-g7/store"
	"github.com/NouKinSG/cloud-station-g7/store/aliyun"
	"github.com/stretchr/testify/assert"
)

var (
	uploader store.Uploader
)

var (
	AccessKey       = os.Getenv("ALI_AK")
	AccessKeySecret = os.Getenv("ALI_SK")
	OssEndpoint     = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName      = os.Getenv("ALI_BUCKET_NAME")
)

// Aliyun Oss Store Upload测试用例

func TestUpload(t *testing.T) {

	// 使用asser 编写一U个测试用例的断言语句
	// 通过New 获取一个断言实例
	should := assert.New(t)

	err := uploader.Upload(BucketName, "ttt.txt", "store_test.go")
	if should.NoError(err) {
		//没有 Error 开启下一个步骤
		t.Log("upload OK")
	}
}

func TestUploadError(t *testing.T) {

	// 使用asser 编写一U个测试用例的断言语句
	// 通过New 获取一个断言实例
	should := assert.New(t)

	err := uploader.Upload(BucketName, "ttt.txt", "store_testxxx.go")
	should.Error(err, "open store_testxxx.go: The system cannot find the file specified.")
}

// 通过init 编写 upload实例化逻辑
func init() {
	ali, err := aliyun.NewDefaultNewAliOssStore()
	if err != nil {
		panic(err)
	}
	uploader = ali
}
