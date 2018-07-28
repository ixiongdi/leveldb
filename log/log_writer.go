package log

import (
	"aahframework.org/test.v0/assert"
	"hash/crc32"
	"unsafe"
)

type WritableFile interface {
	Append([]byte) bool
	Close() bool
	Flush() bool
	Sync() bool
}

type Writer struct {
	block []byte
	blockOffset int
	typeCrc     []uint32
}

func InitTypeCrc(typeCrc *[]uint32) {
	for i := 0; i < kMaxRecordType; i++ {
		//t := byte(i)

		//typeCrc[i] = crc32.ChecksumIEEE(&t)
	}
}

func (write *Writer) Append(data []byte) bool {
	append(write.block, ...data)
	return true
}

func (write *Writer) Flush() bool {
	return false
}

func (write *Writer) AddRecord(data []byte) {

	left := len(data)

	//ptr := &data

	offset := 0

	// Status
	s := true

	begin := true

	for s && left > 0 {
		leftOver := kBlockSize - write.blockOffset

		//assert.True(t, )

		if leftOver < kHeaderSize {
			if leftOver > 0 {
				write.Append(make([]byte, leftOver))
			}
			write.blockOffset = 0
		}

		avail := kBlockSize - write.blockOffset - kHeaderSize

		fragmentLength := 0

		if left < avail {
			fragmentLength = left
		} else {
			fragmentLength = avail
		}

		recordType := kZeroType

		end := left == fragmentLength

		if (begin && end) {
			recordType = kFullType
		} else if (begin) {
			recordType = kFirstType
		} else if (end) {
			recordType = kLastType
		} else {
			recordType = kMiddleType
		}

		s = write.EmitPhysicalRecord(recordType, data, offset, fragmentLength)

		offset += fragmentLength

		//p = unsafe.Pointer(uintptr(p) + offset)

		//ptr = unsafe.Pointer(uintptr(ptr) + uintptr(fragmentLength))

		left -= fragmentLength

		begin = false
	}
}
func (write *Writer) EmitPhysicalRecord(t RecordType, data []byte, offset int,  n int) bool {
	// Format the header
	buf := make([]byte, kHeaderSize)
	buf[4] = byte(n & 0xff)
	buf[5] = byte(n >> 8)
	buf[6] = byte(t)

	// Compute the crc of the record type and the payload.
	crc := crc32.ChecksumIEEE(data[offset:offset+n])

	// Write the header and the payload
	s := write.Append(buf)

	if (s) {
		s = write.Append(data[offset:offset+n])
		if (s) {
			s = write.Flush()
		}
	}

	write.blockOffset += kHeaderSize + n

	return  s
}
