package tools

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func FileLinesCount(filePath string) (int, error) {
	var count int
	var read int
	var err error

	r, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		ZapLogger("both").Error(fmt.Sprintf("error opening file: %v", err))
		return count, err
	}

	var target []byte = []byte("\n")

	buffer := make([]byte, 32*1024)

	for {
		read, err = r.Read(buffer)
		if err != nil {
			break
		}

		count += bytes.Count(buffer[:read], target)
	}

	if err == io.EOF {
		return count, nil
	}

	return count, err
}
