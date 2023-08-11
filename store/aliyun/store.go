package aliyun

import (
	"fmt"

	"github.com/NouKinSG/cloud-station-g7/store"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	// 对象是否实现了接口的约束
	_ store.Uploader = &AliOssStore{}
)

// AliOssStore 对象的构造
func NewAliOssStore(endpoint, accesskey, accessSercret string) (*AliOssStore, error) {
	c, err := oss.New(endpoint, accesskey, accessSercret)
	if err != nil {
		return nil, err
	}

	return &AliOssStore{
		client: c,
	}, nil
}

type AliOssStore struct {
	client *oss.Client
}

func (s *AliOssStore) Upload(bucketName string, objectKey string, fileName string) error {

	// 2、获取bucket对象
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return err
	}
	// 3、上传文件到该bucket
	if err := bucket.PutObjectFromFile(objectKey, fileName); err != nil {
		return err
	}

	// 4、打印下载链接
	downloadURL, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("文件下载URL： %s \n", downloadURL)
	fmt.Println("请在一天内下载")

	return nil
}
