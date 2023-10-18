package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

func IntToHex(numb int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, numb)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
