package leveldb

import (
	"testing"
	"syscall"
	"os"
	"log"
	"unsafe"
	"fmt"
)

func TestPut(t *testing.T) {
	Put("hello", "world")
	t.Log("insert a record")
	value := Get("hello")
	if value == "world" {
		t.Log("get a record")
	} else {
		t.Error("insert error")
	}

	t.Log("save")
	defer Save()
	t.Log("save")
}

func TestGet(t *testing.T) {
	value := Get("hello")
	if value == "world" {
		t.Log("get a record")
	} else {
		t.Error("insert error")
	}
	defer Save()
}

func TestDelete(t *testing.T) {
	Delete("hello")
	value := Get("hello")
	if value == "world" {
		t.Error("get a record")
	} else {
		t.Log("insert error")
	}
	defer Save()
}

func TestMmap(t *testing.T)  {
	file, err := os.Create("/tmp/test.dat")

	if err != nil {
		log.Fatal(err)
	}


	mmap, err := syscall.Mmap(int(file.Fd()), 0, 1000, syscall.PROT_READ | syscall.PROT_WRITE, syscall.MAP_SHARED)

	if err != nil {
		log.Fatal(err)
	}

	map_array := (*[1000]int)(unsafe.Pointer(&mmap[0]))

	for i := 0; i < 1000; i++  {
		map_array[i] = i
	}

	fmt.Println(map_array)
}

func TestTmpDir(t *testing.T) {
	file, _ := os.Open("./")
	names, err := file.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(names)
	//file, _ := ioutil.("./", "test");
	//file.WriteString("写入字符串");
	file.Close();
}