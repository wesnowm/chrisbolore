package filehandler

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

func GetFileHash(f *os.File) string {
	h := md5.New()

	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetBigFileHash(f *os.File) string {
	info, _ := f.Stat()
	fileSize := info.Size()

	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))

	hash := md5.New()

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))
		buf := make([]byte, blocksize)

		file.Read(buf)
		io.WriteString(hash, string(buf)) // append into the hash
	}

	fmt.Sprintf("%x", hash.Sum(nil))
}