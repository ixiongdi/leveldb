package wal

import (
	"bytes"
	"hash/crc32"
	"encoding/binary"
	"fmt"
)

const (
	BlockSize = 32

	RecordHeadSize = 4 + 2 + 1

	RecordMaxDataSize = BlockSize - RecordHeadSize

	Full = 1
	First = 2
	Middle = 3
	Last = 4
)
//var manger = Manager{0}
//var currentBlock = Block{}

type OP struct {
	manager *Manager
	currentBlock *Block
}

func NewOP() OP {
	return OP{
		&Manager {
			0,
		},
		&Block{},
	}
}

func (op *OP) WriteAbleFile(data []byte)  {
	manger := op.manager

	currentBlock := op.currentBlock

	//fmt.Println(manger.current)

	length := len(data)

	free := currentBlock.GetFreeSize()

	//fmt.Println(free)

	// 当前块剩余空间小于6，直接写到硬盘，并生成新的空块
	if free <= 6 {
		// 用空填充剩余空间
		currentBlock.FillPadding(free)
		// 写入磁盘
		manger.Writer(currentBlock.GetBytes())
		// 生成新的块
		currentBlock.Remove()

		free = currentBlock.GetFreeSize()
	}

	fmt.Println(free)


	// 当前块剩余空间等于7并且数据的长度为0，可以直接写入硬盘，并生成新的空块
	if free == 7 && length == 0 {
		// 新增记录，刚好写满Block
		currentBlock.AddRecord(NewRecord(Full, data))
		// 写入磁盘
		manger.Writer(currentBlock.GetBytes())
		// 生成新块
		currentBlock.Remove()
	} else if RecordHeadSize + length <= free {
		// 不需要生成新的块，直接写入当前Block
		currentBlock.AddRecord(NewRecord(Full, data))
	} else if RecordHeadSize + length > free {
		// 当前块剩余空间不足以写下记录，用剩余空间写下First
		offset := free-RecordHeadSize
		firstRecord := NewRecord(First, data[:offset])
		currentBlock.AddRecord(firstRecord)
		manger.Writer(currentBlock.GetBytes())
		currentBlock.Remove()
		// 还需要多少块
		blockCount := (length - offset) / RecordMaxDataSize + 1

		// 还需要一块就够了
		if blockCount == 1 {
			currentBlock.AddRecord(NewRecord(Last, data[offset:]))
		} else {
			// 插入中间块
			for i := 0; i < blockCount - 1; i++ {
				//start := free - RecordHeadSize + RecordMaxDataSize * i

				offset += RecordMaxDataSize * i

				fmt.Println(offset)
				fmt.Println(length)

				currentBlock.AddRecord(NewRecord(Middle, data[offset:offset+RecordMaxDataSize]))
				manger.Writer(currentBlock.GetBytes())
				currentBlock.Remove()
			}

			// 插入最后块
			currentBlock.AddRecord(NewRecord(Last, data[offset:]))
		}
	}


}

func SequentialFile() {

}

/**
+---------+-----------+-----------+--- ... ---+
|CRC (4B) | Size (2B) | Type (1B) | Payload   |
+---------+-----------+-----------+--- ... ---+

CRC = 32bit hash computed over the payload using CRC
Size = Length of the payload data
Type = Type of record
       (kZeroType, kFullType, kFirstType, kLastType, kMiddleType )
       The type is used to group a bunch of records together to represent
       blocks that are larger than kBlockSize
Payload = Byte stream as long as specified by the payload size
 */
type Record struct {
	CRC uint32
	Size uint16
	Type uint8
	Payload []byte
}
// 新建记录，根据类型和数据
func NewRecord(typ uint8, payload []byte) Record {
	// 新建数据缓冲区，长度是type的长度加上数据的长度，type一个字节长度
	buf := bytes.NewBuffer(make([]byte, 1 + len(payload)));

	buf.WriteByte(typ)
	buf.Write(payload)

	// 计算crc32
	crc := crc32.ChecksumIEEE(buf.Bytes())

	return Record{
		crc,
		uint16(len(payload)),
		typ,
		payload,
	}
}

func (record *Record) GetBytes() []byte {
	buf := bytes.NewBuffer(make([]byte, RecordHeadSize + len(record.Payload)))

	buf.Write(uint32ToBytes(record.CRC))
	buf.Write(uint16ToBytes(record.Size))
	buf.WriteByte(record.Type)
	buf.Write(record.Payload)

	return buf.Bytes()
}

/**
      +-----+-------------+--+----+----------+------+-- ... ----+
 File  | r0  |        r1   |P | r2 |    r3    |  r4  |           |
       +-----+-------------+--+----+----------+------+-- ... ----+
       <--- kBlockSize ------>|<-- kBlockSize ------>|

  rn = variable size records
  P = Padding
 */
type Block struct {
	records []Record
	padding []byte
}

func (block *Block) GetFreeSize() int {
	var free = BlockSize

	for _, record := range block.records {
		free -= RecordHeadSize + len(record.Payload)
	}

	return free
}

func (block *Block) GetBytes() []byte {
	buf := bytes.NewBuffer(make([]byte, BlockSize))

	for i := 0; i < len(block.records); i++ {
		buf.Write(block.records[i].GetBytes())
	}
	buf.Write(block.padding)

	return buf.Bytes()
}

func (block *Block) Remove()  {
	block.records = nil
	block.padding = nil
}

func (block *Block) AddRecord(record Record)  {
	//block.records[len(block.records)] = record
	block.records = append(block.records, record)

	//fmt.Println(len(block.records))
}

func (block *Block) FillPadding(free int) {
	block.padding = make([]byte, free)
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


