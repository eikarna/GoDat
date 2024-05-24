package main

import (
	"fmt"
	log "github.com/codecat/go-libs/log"
	"github.com/eikarna/GoDat/Components/Decoder"
	"github.com/eikarna/GoDat/Components/Encoder"
	. "github.com/eikarna/GoDat/Components/Enums"
	"github.com/eikarna/GoDat/Components/ProtonHash"
	"github.com/eikarna/GoDat/Components/UI"
	"github.com/goccy/go-json"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	UI.PrintBanner(true, "GoDat", "Version: 1.0.0", "By: Eikarna")
	fmt.Println("Please select Method (ex. 1, or \"Decode\"):")
	fmt.Println("1. Decode")
	fmt.Println("2. Encode")
	fmt.Println("3. Search Items (Name/ID)")
	fmt.Println("4. ProtonHash")
	method, err := UI.Read(false, "> ")
	if err != nil {
		log.Error(err.Error())
		return
	}
	method = strings.TrimSpace(method)
	if method == "1" || strings.Contains(strings.ToLower(method), "decode") {
		targetDecodeFile, err := UI.Read(false, "File Name to be Decoded: ")
		if err != nil {
			log.Error(err.Error())
			return
		}
		now := time.Now()
		decoded, err := Decoder.Decode(targetDecodeFile, now)
		if err != nil {
			log.Error("Error ketika mencoba decode \"items.dat\": %v", err)
			return
		}
		data, _ := json.MarshalIndent(decoded, "", " ")
		_ = ioutil.WriteFile("items.json", data, 0644)
	} else if method == "2" || strings.Contains(strings.ToLower(method), "encode") {
		targetEncodeFile, err := UI.Read(false, "File Name to be Encoded (Only Valid JSON): ")
		if err != nil {
			log.Error(err.Error())
			return
		}
		outputNameFile, err := UI.Read(false, "Output File Name: ")
		if err != nil {
			log.Error(err.Error())
			return
		}
		b, err := os.ReadFile(targetEncodeFile) // just pass the file name
		if err != nil {
			log.Error(err.Error())
			return
		}
		now := time.Now()
		data := &ItemInfo{}
		err = json.Unmarshal(b, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
		err = Encoder.Encode(data, outputNameFile, now)
		if err != nil {
			log.Error("Error ketika mencoba encode \"items.dat\": %v", err)
			return
		}
	} else if method == "3" || strings.Contains(strings.ToLower(method), "search items") {
		log.Warn("This Feature will added Soon!")
		return
	} else if method == "4" || strings.Contains(strings.ToLower(method), "protonhash") {
		targetFile, err := UI.Read(false, "Target File: ")
		if err != nil {
			log.Error(err.Error())
			return
		}
		fmt.Printf("ProtonHash: %d\n", ProtonHash.GetFileHash(targetFile))
	} else {
		log.Fatal("Please input Correctly!")
		return
	}
}
