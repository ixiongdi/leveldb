package leveldb

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main()  {

	defer Save()

	// 监听8080端口
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	// 死循环
	for {
		// 拿到连接
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		// 处理函数
		go handleConnection(conn)
	}


}

func handleConnection(conn net.Conn)  {

	// Response，返回数据
	send := bufio.NewWriter(conn)

	// Request，接受数据
	scanner := bufio.NewScanner(conn)

	// 这里以\n作为分隔符
	for scanner.Scan() {
		// 命令分隔符，这里以一个空格作为分隔符
		arr := strings.Split(scanner.Text(), " ")

		command := arr[0]

		fmt.Println(command)


		switch command {
		case "set":
			key, value := arr[1], arr[2]
			Put(key, value)
		    send.WriteString("Set success!\n")
		case "get":
			key := arr[1]

			value := Get(key)

			if value == "" {
				send.WriteString("Get fail, key not found!\n")
			} else {
				send.WriteString("Get success, this is result: ")
				send.WriteString(value)
				send.WriteString("\n")
			}

		case "del":
			key := arr[1]
			Delete(key)
		    send.WriteString("Del success!\n")
		default:
			fmt.Println("Not support command!")
		}

		send.Flush()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
