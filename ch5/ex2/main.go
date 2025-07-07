package ex2

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	count, err := fileLen(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bytes in", os.Args[1], "=", count)
}

func getFile(fileName string) (*os.File, func(), error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}

	return file, func() { file.Close() }, nil
}

func fileLen(fileName string) (int, error) {
	file, closer, err := getFile(fileName)
	defer closer()
	if err != nil {
		return 0, err
	}
	b := make([]byte, 2048)
	count, err := file.Read(b)
	if err != nil {
		if err != io.EOF {
			return 0, err
		}
	}
	return count, nil
}
