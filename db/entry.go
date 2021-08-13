package db

const (
	CommonFileLength = 10
)

type Entry struct {
	KeySize   uint32
	ValueSize uint32
	Mark      uint16 //	该entry的操作标记
	Key       []byte
	Value     []byte
}
