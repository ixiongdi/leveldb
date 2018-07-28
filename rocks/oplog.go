package rocks

import (
	"os"
	"unsafe"
	"log"
	"syscall"
	"hash/crc32"
	"bytes"
	"encoding/binary"
)

const (
	OpLogSize = 32768 * 100

	T = int(unsafe.Sizeof(0)) * OpLogSize

	FULL = 1
	FIRST = 2
	MIDDLE = 3
	LAST = 4

	RecordMaxLength = 32768
	RecordHeadLength = 4 + 2 + 1
	RecordDataLength = RecordMaxLength - RecordHeadLength
)

type OpLog struct {
	name string
	file *os.File
	mMap []byte
}

type Record struct {
	checksum uint32
	length uint16
	logType uint8
	data []byte
}

func (opLog OpLog) NewOpLog() OpLog {
	opLog.name = "/tmp/op.log"

	file, err := os.Create(opLog.name)
	if err != nil {
		log.Fatal(err)
	}
	opLog.file = file

	_, err = opLog.file.Seek(int64(T - 1), 0)
	if err != nil {
		log.Fatal(err)
	}

	_, err = opLog.file.Write([]byte(" "))
	if err != nil {
		log.Fatal(err)
	}

	mmap, err := syscall.Mmap(int(opLog.file.Fd()), 0, int(T), syscall.PROT_READ | syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Fatal(err)
	}

	//mapArray := (*[OpLogSize]byte)(unsafe.Pointer(&mmap[0]))

	opLog.mMap = mmap

	return opLog
}

func (opLog OpLog) AddRecord(data []byte) {

	length := len(data)

	if length > RecordDataLength {
		//bytes.SplitN(data, nil, RecordDataLength)


		recordCount := length / RecordDataLength

		firstData := data[:RecordDataLength]

		firstRecord := Record{
			crc32.ChecksumIEEE(firstData),
			RecordDataLength,
			FIRST,
			firstData,
		}

		opLog.WriteRecord(firstRecord)

		for i := 1; i < recordCount; i++ {
			offset := RecordDataLength * i
			data := data[offset:offset+RecordDataLength]
			record := Record{
				crc32.ChecksumIEEE(data),
				RecordDataLength,
				MIDDLE,
				data,
			}

			opLog.WriteRecord(record)
		}
	} else {
		record := Record{
			crc32.ChecksumIEEE(data),
			uint16(len(data)),
			FULL,
			data,
		}

		opLog.WriteRecord(record)
	}
}

func (opLog OpLog) WriteRecord(record Record) {
	buf := bytes.NewBuffer(opLog.mMap)

	buf.Write(uint32ToBytes(record.checksum))
	buf.Write(uint16ToBytes(record.length))
	buf.WriteByte(record.logType)
	buf.Write(record.data)
}


func uint64ToBytes(i uint64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, i)
	return buf
}

func uint32ToBytes(i uint32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, i)
	return buf
}

func uint16ToBytes(i uint16) []byte {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, i)
	return buf
}


