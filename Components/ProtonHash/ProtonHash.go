package ProtonHash

import (
	"fmt"
	"os"
)

func GetHash(str []byte, length int) uint32 {
	n := str
	acc := uint32(0x55555555)
	if length == 0 {
		for _, c := range n {
			acc = (acc >> 27) + (acc << 5) + uint32(c)
		}
	} else {
		for i := 0; i < length; i++ {
			acc = (acc >> 27) + (acc << 5) + uint32(n[i])
		}
	}
	return acc
}

func GetFileHash(pathFile string) uint32 {
	file, err := os.Open(pathFile)
	if err != nil {
		fmt.Errorf("error opening file: %v", err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Errorf("error getting file information: %v", err)
	}

	size := fileInfo.Size()
	if size == -1 {
		fmt.Errorf("error: File size is -1")
	}

	data := make([]byte, size)
	_, err = file.Read(data)
	if err != nil {
		fmt.Errorf("error reading file: %v", err)
	}

	return GetHash(data, int(size))
}
