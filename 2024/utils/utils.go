package utils

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ReadInput(filePath string) (lines []string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("The file is not apparent", err)
	}
	defer file.Close()

	lines = []string{}

	rd := bufio.NewScanner(file)
	for rd.Scan() {
		line := rd.Text()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("read file line error: %v", err)
			return nil
		}
		lines = append(lines, line)
	}
	return lines

}
