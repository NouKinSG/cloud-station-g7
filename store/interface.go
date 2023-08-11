package store

//定义如何上传文件到 bucket
// 做了抽象 ，并不关心我们需要上传到哪个OSS的bucket
type Uploader interface {
	Upload(bucketName string, objectKey string, fileName string) error
}
