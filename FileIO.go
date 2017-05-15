package main

import (
	"fmt"
	"os"
	"log"
	"syscall"
)

const (
		O_RDONLY int = syscall.O_RDONLY // open the file read-only.
		O_WRONLY int = syscall.O_WRONLY // open the file write-only.
		O_RDWR   int = syscall.O_RDWR   // open the file read-write.
		O_APPEND int = syscall.O_APPEND // append data to the file when writing.
		O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
		O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist
		O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
		O_TRUNC  int = syscall.O_TRUNC  // if possible, truncate file when opened.
	)

func fileOpen() {
	//os.Open will open the file with readonly as default for other modes we need to use OpenFiles
	file, err := os.OpenFile("c:\\users\\rajjame\\downloads\\file.json", O_RDONLY, 0666) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

}

func fileAppend() {
	//os.Open will open the file with readonly as default for other modes we need to use OpenFiles
	file, err := os.OpenFile("c:\\users\\rajjame\\downloads\\file.txt",  O_APPEND | O_WRONLY, 0666) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	str := "eleven"
	//count , err := file.WriteString(str)

	fmt.Fprintln(file,str)

	_ , err = file.WriteString("\r")

	str = "twelve"
	fmt.Fprintln(file,str)

	//count , err := file.WriteString(str)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Printf("Written %d bytes \n", count)
}

