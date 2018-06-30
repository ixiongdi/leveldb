package main

import (
	"encoding/csv"
	"os"
	"log"
	"io/ioutil"
	"bytes"
)

const FILE = "./db.csv"

// 声明一个map
var m map[string]string


func init()  {
	// 初始化map
	m = make(map[string]string)

	// 打开文件
	file, err := os.Open(FILE)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// 读取CSV文件
	reader := csv.NewReader(file)

	// 读取数据，这里data的类型是[][]string
	data, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	// 循环数组
	for i := 0; i < len(data); i++  {
		// 获取行
		row := data[i]

		// 第一个为key
		key := row[0]
		// 第二个为value
		value := row[1]

		// 保持到内存
		m[key] = value
	}
}
// 新建，修改
func Put(key string, value string) {
	m[key] = value

	defer Save()
}
// 删除
func Delete(key string)  {
	delete(m, key)
}
// 读取
func Get(key string) string {
	return m[key]
}

// 写入文件
func Save()  {

	// 新建写入流
	buf := new(bytes.Buffer)

	writer := csv.NewWriter(buf)

	log.Print(m)
	for k, v := range m {
		log.Printf("%s: %s", k, v)
		// 写入数据
		writer.Write([]string{k, v})
		writer.Flush()
	}

	log.Println(buf)

	// 保存到文件
	ioutil.WriteFile(FILE, buf.Bytes(), 0777 )
}

