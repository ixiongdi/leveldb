package log

type RecordType int

const (
	// Zero is reserved for preallocated files
	kZeroType RecordType = iota
	kFullType
	// For fragments
	kFirstType
	kMiddleType
	kLastType
)

const (
	kMaxRecordType = 4
	kBlockSize = 32768
	// Header is checksum (4 bytes), length (2 bytes), type (1 byte).
	kHeaderSize = 4 + 2 + 1
)


