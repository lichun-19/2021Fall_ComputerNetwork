package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleConnection(c net.Conn) {
	f, err := os.Create("whatever.txt")
	defer f.Close()
	check(err)
	input := bufio.NewScanner(bufio.NewReader(c))
	sz := ""
	line := 0
	sz_count := 0
	output := ""
	if input.Scan() {
		sz = input.Text()
	}
	ori_size, _ := strconv.Atoi(sz)
	//fmt.Printf("%d\n", ori_size)
	for (sz_count < ori_size) && input.Scan() {
		line++
		sz_count += len(input.Text()) + 1
		output += fmt.Sprintf("%d ", line) + input.Text() + "\n"
	}
	fmt.Printf("Upload file size: %d\n", ori_size)

	writer := bufio.NewWriter(f)
	wri_bytes, errw := writer.WriteString(output)
	fmt.Printf("Output file size: %d\n", wri_bytes)
	check(errw)
	net_writer := bufio.NewWriter(c)
	reply_message := fmt.Sprintf("%d bytes received, %d btyes file generated\n", ori_size, wri_bytes)
	_, errnw := net_writer.WriteString(reply_message)
	check(errnw)
	net_writer.Flush()
	writer.Flush()
}
func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12012")
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		defer conn.Close()

		go handleConnection(conn)
	}

}
