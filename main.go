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
	"strconv"
	"strings"
	"time"
)

var (
	cachedItemInfo *ItemInfo
)

func main() {
	for {
		UI.ClearScreen()
		UI.PrintBanner(true, "GoDat", "Version: 1.0.0", "By: Eikarna")
		fmt.Println("Please select Method (ex. 1, or \"Decode\"):")
		fmt.Println("1. Decode")
		fmt.Println("2. Encode")
		fmt.Println("3. Search Items (Name/ID)")
		fmt.Println("4. ProtonHash")
		fmt.Println("Type 'exit' to quit the program.")

		method, err := UI.Read(false, "> ")
		if err != nil {
			log.Error(err.Error())
			continue
		}

		method = strings.TrimSpace(strings.ToLower(method))
		if method == "exit" {
			fmt.Println("Exiting program.")
			break
		}

		switch method {
		case "1", "decode":
			handleDecode()
		case "2", "encode":
			handleEncode()
		case "3", "search items":
			handleSearchItems()
		case "4", "protonhash":
			handleProtonHash()
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func handleDecode() {
	for {
		targetDecodeFile, err := UI.Read(false, "File Name to be Decoded (or 'back' to return): ")
		if err != nil {
			log.Error(err.Error())
			return
		}

		if strings.TrimSpace(strings.ToLower(targetDecodeFile)) == "back" {
			return
		}

		now := time.Now()
		decoded, err := Decoder.Decode(targetDecodeFile, now)
		if err != nil {
			log.Error("Error when trying to decode \"%s\": %v", targetDecodeFile, err)
			continue
		}

		data, _ := json.MarshalIndent(decoded, "", " ")
		_ = ioutil.WriteFile("items.json", data, 0644)
		fmt.Println("File decoded successfully. Output saved to items.json.")
		return
	}
}

func handleEncode() {
	for {
		targetEncodeFile, err := UI.Read(false, "File Name to be Encoded (Only Valid JSON) (or 'back' to return): ")
		if err != nil {
			log.Error(err.Error())
			return
		}

		if strings.TrimSpace(strings.ToLower(targetEncodeFile)) == "back" {
			return
		}

		outputNameFile, err := UI.Read(false, "Output File Name: ")
		if err != nil {
			log.Error(err.Error())
			return
		}

		b, err := os.ReadFile(targetEncodeFile)
		if err != nil {
			log.Error(err.Error())
			continue
		}

		now := time.Now()
		data := &ItemInfo{}
		err = json.Unmarshal(b, data)
		if err != nil {
			log.Error(err.Error())
			continue
		}

		err = Encoder.Encode(data, outputNameFile, now)
		if err != nil {
			log.Error("Error when trying to encode \"%s\": %v", targetEncodeFile, err)
			continue
		}

		fmt.Println("File encoded successfully. Output saved to", outputNameFile)
		return
	}
}

func handleSearchItems() {
	for {
		isId := false
		var itemId int32

		targetItem, err := UI.Read(false, "Items Name/ID (or 'back' to return): ")
		if err != nil {
			log.Error(err.Error())
			return
		}

		if strings.TrimSpace(strings.ToLower(targetItem)) == "back" {
			return
		}

		if id, err := strconv.Atoi(targetItem); err == nil {
			fmt.Printf("%q looks like a number.\n", targetItem)
			isId = true
			itemId = int32(id)
		}

		if cachedItemInfo == nil {
			data, err := os.ReadFile("items.json")
			if err != nil {
				log.Error(err.Error())
				return
			}
			err = json.Unmarshal(data, &cachedItemInfo)
			if err != nil {
				log.Error(err.Error())
				return
			}
		}

		found := false
		for _, key := range cachedItemInfo.Items {
			if isId && key.ItemID == int32(itemId) {
				fmt.Println("Got Items with ID:", strconv.Itoa(int(key.ItemID)))
				fmt.Printf("%v\n", key)
				found = true
				break
			} else if !isId && strings.Contains(strings.ToLower(key.Name), strings.ToLower(targetItem)) {
				fmt.Println("Got Items with Name:", key.Name)
				fmt.Printf("%v\n", key)
				found = true
				break
			}
		}

		if !found {
			fmt.Println("No items found. Please try again.")
		}
	}
}

func handleProtonHash() {
	for {
		targetFile, err := UI.Read(false, "Target File (or 'back' to return): ")
		if err != nil {
			log.Error(err.Error())
			return
		}

		if strings.TrimSpace(strings.ToLower(targetFile)) == "back" {
			return
		}

		fmt.Printf("ProtonHash: %d\n", ProtonHash.GetFileHash(targetFile))
	}
}
