package aws

func NewAwsOssStore() *AwsOssStore {
	return &AwsOssStore{}
}

type AwsOssStore struct {
}

func (s *AwsOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	return nil
}
