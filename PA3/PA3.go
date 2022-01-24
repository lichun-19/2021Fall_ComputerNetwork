package main

import "fmt"
import "bufio"
import "net"
import "os"


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	conn, errc := net.Dial("tcp", "140.112.42.221:12000")
	check(errc)
	defer conn.Close()
	
	// select a file
	fmt.Printf("Input filename: ")
	filename := ""
	fmt.Scanf("%s", &filename)
	f, erro := os.Open(filename)
	check(erro)
	defer f.Close()

	scan := bufio.NewScanner(f)
	content := ""
	for scan.Scan() {
		content += scan.Text() + "\n"
	}

	// sending the file
	fmt.Printf("Send the file size first: %d bytes\n", len(content))
	writer := bufio.NewWriter(conn)
	message := ""
	message = fmt.Sprintf("%d\n%s", len(content),  content)
	_, errw := writer.WriteString(message)
	check(errw)
	writer.Flush()

	// get reply from server	
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
	fmt.Printf("Server replies: %s\n", scanner.Text())
	}
}
