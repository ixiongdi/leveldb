package rocks

import (
	"os"
	"log"
)

type DB struct {
	name string
	file *os.File
}

type Options struct {
	createIfMissing bool
}

type Status struct {
	ok bool
}

func (db DB) Open(options Options, name string) DB {
	db.name = name

	_, err := os.Stat(name)

	if err != nil {
		if os.IsNotExist(err) {
			if options.createIfMissing {
				db.file, err = os.Create(name)
			} else {
				log.Fatalf("file is not exist.")
			}
		} else {
			log.Fatal(err)
		}
	} else {
		db.file, err = os.Open(name)

		if err != nil {
			log.Fatal(err)
		}
	}

	return db;
}

func (db DB) Get(key []byte) Status {
	return Status{};
}

func (db DB) Put(key []byte, value []byte) Status{
	return Status{}
}

func (db DB) Delete(key []byte) Status {
	return Status{}
}

func (db DB) NewIterator() {
}
