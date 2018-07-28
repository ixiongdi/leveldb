package wal

import (
	"math/rand"
	"testing"
	"time"
	"fmt"
)
type Class struct {
	age int
}
func(class *Class) AddAge() {
	class.age++
}

func TestAddAge(t *testing.T) {
	class := Class{0}

	for i := 0; i < 1; i++  {
		class.AddAge()

		fmt.Println(class)
	}
}

func  GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func TestWriteAbleFile(t *testing.T) {
	//rand.Uint32()

	op := NewOP()

	for i := 0; i < 10; i++ {
		s := GetRandomString(100)
		data := []byte(s)

		fmt.Println(op.manager.current)

		op.WriteAbleFile(data)
	}
}
