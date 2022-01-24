package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func check(e error) {
 if e != nil {
 panic(e)
 }
}

func main() {
 fmt.Printf("Input filename:\n")
 inFileName := ""
 fmt.Scanf("%s", &inFileName)

 fmt.Printf("Output filename:\n")
 outFileName := ""
 fmt.Scanf("%s", &outFileName) 

 inFile,err := os.Open(inFileName)
 check(err)
 defer inFile.Close()//execute at the end

 outFile,err := os.Create(outFileName)
 check(err)
 defer outFile.Close()//execute at the end

 scanner := bufio.NewScanner(inFile)
 writer := bufio.NewWriter(outFile)
 i:=1
 for scanner.Scan(){
    writer.WriteString(strconv.Itoa(i)+" "+scanner.Text()+"\n")
    writer.Flush()
    i=i+1
}

 
}