package tx

func NewTxStore() *TxOssStore {
	return &TxOssStore{}
}

type TxOssStore struct {
}

func (s *TxOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	return nil
}
