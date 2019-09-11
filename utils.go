package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/mr-tron/base58"
)

func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)
	return []byte(encode)
}

func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}

	return decode
}

func writeToJSON(fileName string, data interface{}) error {
	JSON, err := json.Marshal(data)
	path := filepath.Join(".", "JSONdata")
	os.MkdirAll(path, os.ModePerm)
	f, err := os.Create(filepath.Join("JSONdata", filepath.Base(fileName+".json")))
	defer f.Close()
	if err != nil {
		return err
	}
	f.Write(JSON)

	return nil
}
