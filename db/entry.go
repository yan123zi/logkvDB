package db

import "encoding/binary"

const (
	CommonFileLength = 10
)

const (
	PUT = iota
	DEL
)

type Entry struct {
	KeySize   uint32
	ValueSize uint32
	Mark      uint16 //	该entry的操作标记
	Key       []byte
	Value     []byte
}

func NewEntry(key string, value string, Mark uint16) *Entry {
	return &Entry{
		KeySize:   uint32(len(key)),
		ValueSize: uint32(len(value)),
		Mark:      Mark,
		Key:       []byte(key),
		Value:     []byte(value),
	}
}

func (e *Entry) GetLen() int64 {
	return int64(CommonFileLength + e.KeySize + e.ValueSize)
}

func EncodeEntry(entry *Entry) []byte {
	buf := make([]byte, entry.GetLen())
	binary.BigEndian.PutUint32(buf[0:4], entry.KeySize)
	binary.BigEndian.PutUint32(buf[4:8], entry.ValueSize)
	binary.BigEndian.PutUint16(buf[8:10], entry.Mark)
	copy(buf[10:CommonFileLength+entry.KeySize], entry.Key)
	copy(buf[CommonFileLength+entry.ValueSize:], entry.Value)
	return buf
}

func DecodeEntry(buf []byte) *Entry {
	entry := &Entry{}
	keySize := binary.BigEndian.Uint32(buf[0:4])
	valueSize := binary.BigEndian.Uint32(buf[4:8])
	mark := binary.BigEndian.Uint16(buf[8:10])
	key := buf[CommonFileLength:entry.KeySize]
	value := buf[CommonFileLength+entry.KeySize:]
	return &Entry{
		KeySize:   keySize,
		ValueSize: valueSize,
		Mark:      mark,
		Key:       key,
		Value:     value,
	}
}
