package behavior_strategy

import "log"

// StorageStrategyImpl 存储策略
type StorageStrategyImpl interface {
	Save(name string, data []byte) error
}

func NewStorageStrategy(t string) StorageStrategyImpl {
	switch t {
	case "enc":
		return &fileEncStorage{}
	default:
		return &fileStorage{}
	}
}

type fileStorage struct {
}

func (fs *fileStorage) Save(s string, data []byte) error {

	// file storage
	log.Println("file storage")
	return nil
}

type fileEncStorage struct {
}

func (fes *fileEncStorage) Save(s string, data []byte) error {

	// file enc storage
	log.Println("file enc storage")
	return nil
}
