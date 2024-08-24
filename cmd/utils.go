package cmd

import (
	"bytes"
	"encoding/binary"
	"os"
	"strconv"
)

func saveText(filename string, text string) error {
	d1 := []byte(text)
	err := os.WriteFile(filename, d1, 0644)
	return err
}

func loadText(filename string) (string, error) {
	dat, err := os.ReadFile(filename)
	return string(dat), err
}

func saveBin(filename string, data []byte) error {

	err := os.WriteFile(filename, data, 0644)
	return err
}

func loadBin(filename string) ([]byte, error) {
	dat, err := os.ReadFile(filename)
	return dat, err
}

func int16ToBytes(val uint16, littleEndian bool) []byte {
	buf := new(bytes.Buffer)

	if littleEndian {
		binary.Write(buf, binary.LittleEndian, val)
	} else {
		binary.Write(buf, binary.BigEndian, val)
	}
	return buf.Bytes()
}

func int32ToBytes(val uint32, littleEndian bool) []byte {
	buf := new(bytes.Buffer)
	if littleEndian {
		binary.Write(buf, binary.LittleEndian, val)
	} else {
		binary.Write(buf, binary.BigEndian, val)
	}
	return buf.Bytes()
}

func isHex(value string) bool {

	_, err := strconv.ParseUint(value, 16, 16)
	if err != nil {
		return false
	}
	return true
}
