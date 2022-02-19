package behavior_strategy

import "testing"

func TestNewStorageStrategy(t *testing.T) {
	var s StorageStrategyImpl

	s = NewStorageStrategy("")
	_ = s.Save("./test_storage.json", []byte("test storage"))

	s = NewStorageStrategy("enc")
	_ = s.Save("./test_enc_storage.json", []byte("test enc storage"))
}
