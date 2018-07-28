package leveldb

type Status uint8

type Block struct {
	trailer *Record
	//data []Record
}

const (
	// status
	Ok = 0
	Fail = 1
	// record type
	ZeroType = 0
	FullType = 1
	FirstType = 2
	MiddleType = 3
	LastType = 4
	// block size
	BlockSize = 32768
	// Header is checksum (4 bytes), length (2 bytes), type (1 byte).
	HeaderSize = 4 + 2 + 1
	//
	TypeValue = 0
	TypeDeletion = 1
)
// OpLog
type Record struct {
	checksum uint32
	length uint16
	typ uint8
	data []uint8
}

// MemTable
type MemRow struct {
	keySize uint32
	key []uint8
	sequenceNum uint64
	typ uint8
	valueSize uint32
	value []uint8
}

type MemTable struct {
	count uint64
	slt *SkipList
}

func (mem MemTable) Add(key []byte, value []byte) {
	var typ = TypeValue

	if value == nil {
		typ = TypeDeletion
	}
	memRow := MemRow{
		uint32(len(key)),
		key,
		mem.count,
		uint8(typ),
		uint32(len(value)),
		value,
	}
	mem.slt.Set(float64(mem.count), memRow)
	mem.count++
}

type DB struct {
	memTable *MemTable
}

type WriteBatch struct {
	data []MemTable;
}

//func (batch WriteBatch) Put(key []byte, value []byte)  {
//	//batch.data[len(batch.data)]
//}

func (db DB) NewDB() DB {
	//db.count = 0;
	//db.slt = New()
	return db
}

func (db DB) Put(key []byte, value []byte) Status {
	//batch := WriteBatch{}
	//
	//batch.Put(key, value)

	//return db.Write(&batch)

	db.memTable.Add(key, value)

	 return Ok
}

func (db DB) Delete(key []byte)  {
	db.memTable.Add(key, nil)
}

func (db DB) Write(batch *WriteBatch) Status {
	db.MakeRoomForWrite()



	return Ok
}

func (db DB) Get(key []byte)  {
	//db.slt.Get()
}

func (db DB) MakeRoomForWrite()  {
}