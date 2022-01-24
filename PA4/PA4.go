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

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12012")
	conn, _ := ln.Accept()
	f, err := os.Create("whatever.txt")
	check(err)
	defer conn.Close()
	defer ln.Close()
	defer f.Close()

	input := bufio.NewScanner(bufio.NewReader(conn))
	sz := ""
	line := 0
	sz_count := 0
	output := ""
	if input.Scan() {
		sz = input.Text()
	}
	ori_size, _ := strconv.Atoi(sz)//string to integer
	for (sz_count < ori_size) && input.Scan() {
		line++
		sz_count += len(input.Text()) + 1
		output += fmt.Sprintf("%d ", line) + input.Text() + "\n" //is yet another API in fmt. It Printf() to a string essentially.
	}
	fmt.Printf("Upload file size: %d\n", ori_size)

	writer := bufio.NewWriter(f)
	wri_bytes, errw := writer.WriteString(output)//write to whatever.txt
	fmt.Printf("Output file size: %d\n", wri_bytes)
	check(errw)
	net_writer := bufio.NewWriter(conn)
	reply_message := fmt.Sprintf("%d bytes received, %d btyes file generated\n", ori_size, wri_bytes)
	_, errnw := net_writer.WriteString(reply_message)// server write back to client
	check(errnw)
	net_writer.Flush()
	writer.Flush()
}
