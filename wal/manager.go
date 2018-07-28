package wal

import (
	"io/ioutil"
	"log"
	"strconv"
)

const (
	Dir = "/tmp/log/"
)

type Manager struct {
	current int
}

func (manager *Manager) Writer(data []byte) {

	//fmt.Println(data)

	err := ioutil.WriteFile(Dir + strconv.Itoa(manager.current) + ".log", data, 0666)

	if err != nil {
		log.Fatal(err)
	} else {
		manager.current++
	}
}

func (manager Manager) Reader() {

}


