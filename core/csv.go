package core

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Data type a alias for map[int][]int
type Data map[int][]int

func openFile(path string) *csv.Reader {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Não foi possível abrir o arquivo!")
	}
	return csv.NewReader(bufio.NewReader(file))
}

func parseInt(value string) (valueInt int) {
	value = strings.TrimSpace(value)
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(fmt.Sprintf("%s: Valor inválido!", value))
	}

	return
}

// ParseAndRead a csv file
func ParseAndRead(path string) (data Data) {
	reader := openFile(path)

	data = make(Data)

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		data[parseInt(line[0])] = append(data[parseInt(line[0])], parseInt(line[1]))
	}

	return
}
