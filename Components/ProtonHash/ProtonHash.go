package ProtonHash

import (
	"fmt"
	"os"
)

func GetHash(str []byte, length int) int32 {
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
	return int32(acc)
}

func GetFileHash(pathFile string) int32 {
	file, err := os.Open(pathFile)
	if err != nil {
		fmt.Errorf("error opening file: %v", err)
		return 0
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Errorf("error getting file information: %v", err)
		return 0
	}

	size := fileInfo.Size()
	if size == -1 {
		fmt.Errorf("error: File size is -1")
		return 0
	}

	data := make([]byte, size)
	_, err = file.Read(data)
	if err != nil {
		fmt.Errorf("error reading file: %v", err)
		return 0
	}

	return GetHash(data, int(size))
}

func main() {
	path := "example.txt"
	hash := GetFileHash(path)
	fmt.Printf("Hash of the file %s: %d\n", path, hash)
}
