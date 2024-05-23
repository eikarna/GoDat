package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	log "github.com/codecat/go-libs/log"
	"github.com/goccy/go-json"
	"io/ioutil"
	"os"
	"time"
)

type Item struct {
	DataPos80          []byte
	Data12             []byte
	Data15             []byte
	Name               string
	TexturePath        string
	ExtraFilePath      string
	PetName            string
	PetPrefix          string
	PetSuffix          string
	PetAbility         string
	ExtraOptions       string
	TexturePath2       string
	ExtraOptions2      string
	PunchOption        string
	StrData11          string
	StrData15          string
	StrData16          string
	ExtraFileHash      int32
	ItemID             int32
	TextureHash        int32
	Val1               int32
	DropChance         int32
	ExtrafileHash      int32
	AudioVolume        int32
	WeatherID          int32
	SeedColorA         int8
	SeedColorR         int8
	SeedColorG         int8
	SeedColorB         int8
	SeedOverlayColorA  int8
	SeedOverlayColorR  int8
	SeedOverlayColorG  int8
	SeedOverlayColorB  int8
	GrowTime           int32
	IntData13          int32
	IntData14          int32
	IntData17          int32
	IntData18          int32
	Rarity             int16
	Val2               int16
	IsRayman           int16
	EditableType       int8
	ItemCategory       int8
	ActionType         int16
	HitsoundType       int8
	ItemKind           int8
	TextureX           int8
	TextureY           int8
	SpreadType         int8
	CollisionType      int8
	BreakHits          int8
	ClothingType       int8
	MaxAmount          int8
	SeedBase           int8
	SeedOverlay        int8
	TreeBase           int8
	TreeLeaves         int8
	IsStripeyWallpaper int8
}

type ItemInfo struct {
	ItemVersion int16
	ItemCount   int32

	Items []Item

	//items.dat packet
	FileSize int32
	FileHash uint32
}

func getHash(str []byte, length int) uint32 {
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

func getFileHash(pathFile string) uint32 {
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

	return getHash(data, int(size))
}

func (Info *ItemInfo) GetItemHash() uint32 {
	return uint32(Info.FileHash)
}

func byteArrayToInt(byteSlice []byte) (int, error) {
	var result int
	for _, b := range byteSlice {
		if b < '0' || b > '9' {
			return 0, fmt.Errorf("Invalid byte: %c", b)
		}
		result = result*10 + int(b-'0')
	}
	return result, nil
}

func DecodeItemsDat(pathFile string, timestamp time.Time) (*ItemInfo, error) {
	itemInfo := &ItemInfo{}
	key := "PBG892FXX982ABC*"

	file, err := os.Open(pathFile)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("error getting file information: %v", err)
	}

	size := fileInfo.Size()
	if size == -1 {
		return nil, fmt.Errorf("error: File size is -1")
	}
	itemInfo.FileSize = int32(size)
	itemInfo.FileHash = getFileHash(pathFile)
	data := make([]byte, size)
	_, err = file.Read(data)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	memPos := 0
	itemInfo.ItemVersion = int16(binary.LittleEndian.Uint16(data[memPos:]))
	memPos += 2
	itemInfo.ItemCount = int32(binary.LittleEndian.Uint32(data[memPos:]))
	memPos += 4
	itemInfo.Items = make([]Item, itemInfo.ItemCount)
	for i := 0; int32(i) < itemInfo.ItemCount; i++ {
		// Items Dat Info start from 66
		itemId := binary.LittleEndian.Uint32(data[memPos:])
		if int32(itemId) < itemInfo.ItemCount {
			itemInfo.Items[i].ItemID = int32(itemId)
			memPos += 4
			itemInfo.Items[i].EditableType = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].ItemCategory = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].ActionType = int16(data[memPos])
			memPos += 1
			itemInfo.Items[i].HitsoundType = int8(data[memPos])
			memPos += 1

			strLen := int(binary.LittleEndian.Uint16(data[memPos:]))
			memPos += 2
			for j := 0; j < strLen; j++ {
				itemInfo.Items[i].Name += string(data[memPos] ^ key[(int32(j)+itemInfo.Items[i].ItemID)%int32(len(key))])
				memPos++
			}
			strLen = int(binary.LittleEndian.Uint16(data[memPos:]))
			memPos += 2
			itemInfo.Items[i].TexturePath = string(data[memPos : memPos+strLen])
			memPos += strLen
			itemInfo.Items[i].TextureHash = int32(binary.LittleEndian.Uint32(data[memPos:]))
			memPos += 4
			itemInfo.Items[i].ItemKind = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].Val1 = int32(binary.LittleEndian.Uint32(data[memPos:]))
			memPos += 4
			itemInfo.Items[i].TextureX = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].TextureY = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].SpreadType = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].IsStripeyWallpaper = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].CollisionType = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].BreakHits = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].DropChance = int32(binary.LittleEndian.Uint32(data[memPos:]))
			memPos += 4
			itemInfo.Items[i].ClothingType = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].Rarity = int16(binary.LittleEndian.Uint16(data[memPos:]))
			memPos += 2
			itemInfo.Items[i].MaxAmount = int8(data[memPos])
			memPos += 1
			strLen = int(binary.LittleEndian.Uint16(data[memPos:]))
			memPos += 2
			itemInfo.Items[i].ExtraFilePath = string(data[memPos : memPos+strLen])
			memPos += strLen
			itemInfo.Items[i].ExtraFileHash = int32(binary.LittleEndian.Uint32(data[memPos : memPos+4]))
			memPos += 4
			itemInfo.Items[i].AudioVolume = int32(binary.LittleEndian.Uint32(data[memPos : memPos+4]))
			memPos += 4
			strLen = int(binary.LittleEndian.Uint16(data[memPos:]))
			memPos += 2
			itemInfo.Items[i].PetName = string(data[memPos : memPos+strLen])
			memPos += strLen
			strLen = int(binary.LittleEndian.Uint16(data[memPos:]))
			memPos += 2
			itemInfo.Items[i].PetPrefix = string(data[memPos : memPos+strLen])
			memPos += strLen
			strLen = int(binary.LittleEndian.Uint16(data[memPos:]))
			memPos += 2
			itemInfo.Items[i].PetSuffix = string(data[memPos : memPos+strLen])
			memPos += strLen
			strLen = int(binary.LittleEndian.Uint16(data[memPos:]))
			memPos += 2
			itemInfo.Items[i].PetAbility = string(data[memPos : memPos+strLen])
			memPos += strLen
			itemInfo.Items[i].SeedBase = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].SeedOverlay = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].TreeBase = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].TreeLeaves = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].SeedColorA = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].SeedColorR = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].SeedColorG = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].SeedColorB = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].SeedOverlayColorA = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].SeedOverlayColorR = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].SeedOverlayColorG = int8(data[memPos])
			memPos += 1
			itemInfo.Items[i].SeedOverlayColorB = int8(data[memPos])
			memPos += 1
			memPos += 4
			itemInfo.Items[i].GrowTime = int32(binary.LittleEndian.Uint32(data[memPos:]))
			memPos += 4
			itemInfo.Items[i].Val2 = int16(binary.LittleEndian.Uint16(data[memPos:]))
			memPos += 2
			itemInfo.Items[i].IsRayman = int16(binary.LittleEndian.Uint16(data[memPos:]))
			memPos += 2
			strLen = int(binary.LittleEndian.Uint16(data[memPos:]))
			memPos += 2
			itemInfo.Items[i].ExtraOptions = string(data[memPos : memPos+strLen])
			memPos += strLen

			strLen = int(binary.LittleEndian.Uint16(data[memPos:]))
			memPos += 2
			itemInfo.Items[i].TexturePath2 = string(data[memPos : memPos+strLen])
			memPos += strLen
			strLen = int(binary.LittleEndian.Uint16(data[memPos:]))
			memPos += 2
			itemInfo.Items[i].ExtraOptions2 = string(data[memPos : memPos+strLen])
			memPos += strLen
			itemInfo.Items[i].DataPos80 = data[memPos : memPos+80]
			memPos += 80
			if itemInfo.ItemVersion >= 11 {
				strLen := int(binary.LittleEndian.Uint16(data[memPos:]))
				memPos += 2
				itemInfo.Items[i].PunchOption = string(data[memPos : memPos+strLen])
				memPos += strLen
			}
			if itemInfo.ItemVersion >= 12 {
				itemInfo.Items[i].Data12 = data[memPos : memPos+13]
				memPos += 13
			}
			if itemInfo.ItemVersion >= 13 {
				itemInfo.Items[i].IntData13 = int32(binary.LittleEndian.Uint32(data[memPos : memPos+4]))
				memPos += 4
			}
			if itemInfo.ItemVersion >= 14 {
				itemInfo.Items[i].IntData14 = int32(binary.LittleEndian.Uint32(data[memPos : memPos+4]))
				memPos += 4
			}
			if itemInfo.ItemVersion >= 15 {
				itemInfo.Items[i].Data15 = data[memPos : memPos+25]
				memPos += 25
				strLen := int(binary.LittleEndian.Uint16(data[memPos:]))
				memPos += 2
				itemInfo.Items[i].StrData15 = string(data[memPos : memPos+strLen])
				memPos += strLen
			}
			if itemInfo.ItemVersion >= 16 {
				strLen := int(binary.LittleEndian.Uint16(data[memPos:]))
				memPos += 2
				itemInfo.Items[i].StrData16 = string(data[memPos : memPos+strLen])
				memPos += strLen
			}
			if itemInfo.ItemVersion >= 17 {
				itemInfo.Items[i].IntData17 = int32(binary.LittleEndian.Uint32(data[memPos : memPos+4]))
				memPos += 4
			}
			if itemInfo.ItemVersion >= 18 {
				itemInfo.Items[i].IntData18 = int32(binary.LittleEndian.Uint32(data[memPos : memPos+4]))
				memPos += 4
			}

		} else {
			break
		}
	}

	log.Info("Items.dat decoded for %s. With Item Count: %d, ItemsDatVersion: %d, Item Hash: %v", time.Since(timestamp), itemInfo.ItemCount, itemInfo.ItemVersion, itemInfo.FileHash)
	return itemInfo, nil
}

func EncodeItemsDat(itemInfo *ItemInfo, pathFile string, now time.Time) error {
	key := "PBG892FXX982ABC*"

	buffer := &bytes.Buffer{}

	// Write the item version
	if err := binary.Write(buffer, binary.LittleEndian, itemInfo.ItemVersion); err != nil {
		return fmt.Errorf("error writing item version: %v", err)
	}

	// Write the item count
	if err := binary.Write(buffer, binary.LittleEndian, itemInfo.ItemCount); err != nil {
		return fmt.Errorf("error writing item count: %v", err)
	}

	for _, item := range itemInfo.Items {
		// Write item ID
		if err := binary.Write(buffer, binary.LittleEndian, item.ItemID); err != nil {
			return fmt.Errorf("error writing item ID: %v", err)
		}

		// Write other fields
		if err := binary.Write(buffer, binary.LittleEndian, item.EditableType); err != nil {
			return fmt.Errorf("error writing EditableType: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.ItemCategory); err != nil {
			return fmt.Errorf("error writing ItemCategory: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.ActionType); err != nil {
			return fmt.Errorf("error writing ActionType: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.HitsoundType); err != nil {
			return fmt.Errorf("error writing HitsoundType: %v", err)
		}

		// Encode the Name field with the key
		nameLength := uint16(len(item.Name))
		if err := binary.Write(buffer, binary.LittleEndian, nameLength); err != nil {
			return fmt.Errorf("error writing name length: %v", err)
		}

		for j, char := range item.Name {
			encodedChar := byte(char) ^ key[(int32(j)+item.ItemID)%int32(len(key))]
			if err := buffer.WriteByte(encodedChar); err != nil {
				return fmt.Errorf("error writing encoded name character: %v", err)
			}
		}

		// Encode the TexturePath
		texturePathLength := uint16(len(item.TexturePath))
		if err := binary.Write(buffer, binary.LittleEndian, texturePathLength); err != nil {
			return fmt.Errorf("error writing texture path length: %v", err)
		}

		if _, err := buffer.Write([]byte(item.TexturePath)); err != nil {
			return fmt.Errorf("error writing texture path: %v", err)
		}

		// Write other fields
		if err := binary.Write(buffer, binary.LittleEndian, item.TextureHash); err != nil {
			return fmt.Errorf("error writing TextureHash: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.ItemKind); err != nil {
			return fmt.Errorf("error writing ItemKind: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.Val1); err != nil {
			return fmt.Errorf("error writing Val1: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.TextureX); err != nil {
			return fmt.Errorf("error writing TextureX: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.TextureY); err != nil {
			return fmt.Errorf("error writing TextureY: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.SpreadType); err != nil {
			return fmt.Errorf("error writing SpreadType: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.IsStripeyWallpaper); err != nil {
			return fmt.Errorf("error writing IsStripeyWallpaper: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.CollisionType); err != nil {
			return fmt.Errorf("error writing CollisionType: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.BreakHits); err != nil {
			return fmt.Errorf("error writing BreakHits: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.DropChance); err != nil {
			return fmt.Errorf("error writing DropChance: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.ClothingType); err != nil {
			return fmt.Errorf("error writing ClothingType: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.Rarity); err != nil {
			return fmt.Errorf("error writing Rarity: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.MaxAmount); err != nil {
			return fmt.Errorf("error writing MaxAmount: %v", err)
		}

		// Encode the ExtraFilePath
		extraFilePathLength := uint16(len(item.ExtraFilePath))
		if err := binary.Write(buffer, binary.LittleEndian, extraFilePathLength); err != nil {
			return fmt.Errorf("error writing extra file path length: %v", err)
		}

		if _, err := buffer.Write([]byte(item.ExtraFilePath)); err != nil {
			return fmt.Errorf("error writing extra file path: %v", err)
		}

		/* Padding bytes
		if _, err := buffer.Write(make([]byte, )); err != nil {
			return fmt.Errorf("error writing padding bytes: %v", err)
		}*/

		// Encode PetName
		petNameLength := uint16(len(item.PetName))
		if err := binary.Write(buffer, binary.LittleEndian, petNameLength); err != nil {
			return fmt.Errorf("error writing pet name length: %v", err)
		}

		if _, err := buffer.Write([]byte(item.PetName)); err != nil {
			return fmt.Errorf("error writing pet name: %v", err)
		}

		// Encode PetPrefix
		petPrefixLength := uint16(len(item.PetPrefix))
		if err := binary.Write(buffer, binary.LittleEndian, petPrefixLength); err != nil {
			return fmt.Errorf("error writing pet prefix length: %v", err)
		}

		if _, err := buffer.Write([]byte(item.PetPrefix)); err != nil {
			return fmt.Errorf("error writing pet prefix: %v", err)
		}

		// Encode PetSuffix
		petSuffixLength := uint16(len(item.PetSuffix))
		if err := binary.Write(buffer, binary.LittleEndian, petSuffixLength); err != nil {
			return fmt.Errorf("error writing pet suffix length: %v", err)
		}

		if _, err := buffer.Write([]byte(item.PetSuffix)); err != nil {
			return fmt.Errorf("error writing pet suffix: %v", err)
		}

		// Encode PetAbility
		petAbilityLength := uint16(len(item.PetAbility))
		if err := binary.Write(buffer, binary.LittleEndian, petAbilityLength); err != nil {
			return fmt.Errorf("error writing pet ability length: %v", err)
		}

		if _, err := buffer.Write([]byte(item.PetAbility)); err != nil {
			return fmt.Errorf("error writing pet ability: %v", err)
		}

		// Write other fields
		if err := binary.Write(buffer, binary.LittleEndian, item.SeedBase); err != nil {
			return fmt.Errorf("error writing SeedBase: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.SeedOverlay); err != nil {
			return fmt.Errorf("error writing SeedOverlay: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.TreeBase); err != nil {
			return fmt.Errorf("error writing TreeBase: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.TreeLeaves); err != nil {
			return fmt.Errorf("error writing TreeLeaves: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.SeedColorA); err != nil {
			return fmt.Errorf("error writing SeedColorA: %v", err)
		}
		binary.Write(buffer, binary.LittleEndian, item.SeedColorR)
		binary.Write(buffer, binary.LittleEndian, item.SeedColorG)
		binary.Write(buffer, binary.LittleEndian, item.SeedColorB)

		if err := binary.Write(buffer, binary.LittleEndian, item.SeedOverlayColorA); err != nil {
			return fmt.Errorf("error writing SeedOverlayColor: %v", err)
		}
		binary.Write(buffer, binary.LittleEndian, item.SeedColorR)
		binary.Write(buffer, binary.LittleEndian, item.SeedColorG)
		binary.Write(buffer, binary.LittleEndian, item.SeedColorB)

		if _, err := buffer.Write(make([]byte, 4)); err != nil {
			return fmt.Errorf("error writing padding bytes: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.GrowTime); err != nil {
			return fmt.Errorf("error writing GrowTime: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.Val2); err != nil {
			return fmt.Errorf("error writing Val2: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.IsRayman); err != nil {
			return fmt.Errorf("error writing IsRayman: %v", err)
		}
		// Encode ExtraOptions
		extraOptionsLength := uint16(len(item.ExtraOptions))
		if err := binary.Write(buffer, binary.LittleEndian, extraOptionsLength); err != nil {
			return fmt.Errorf("error writing extra options length: %v", err)
		}

		if _, err := buffer.Write([]byte(item.ExtraOptions)); err != nil {
			return fmt.Errorf("error writing extra options: %v", err)
		}

		// Encode TexturePath2
		texturePath2Length := uint16(len(item.TexturePath2))
		if err := binary.Write(buffer, binary.LittleEndian, texturePath2Length); err != nil {
			return fmt.Errorf("error writing texture path2 length: %v", err)
		}

		if _, err := buffer.Write([]byte(item.TexturePath2)); err != nil {
			return fmt.Errorf("error writing texture path2: %v", err)
		}

		// Encode ExtraOptions2
		extraOptions2Length := uint16(len(item.ExtraOptions2))
		if err := binary.Write(buffer, binary.LittleEndian, extraOptions2Length); err != nil {
			return fmt.Errorf("error writing extra options2 length: %v", err)
		}

		if _, err := buffer.Write([]byte(item.ExtraOptions2)); err != nil {
			return fmt.Errorf("error writing extra options2: %v", err)
		}

		// Padding bytes (80 bytes)
		if _, err := buffer.Write(make([]byte, 80)); err != nil {
			return fmt.Errorf("error writing padding bytes: %v", err)
		}

		// Version-specific fields
		if itemInfo.ItemVersion >= 11 {
			punchOptionLength := uint16(len(item.PunchOption))
			if err := binary.Write(buffer, binary.LittleEndian, punchOptionLength); err != nil {
				return fmt.Errorf("error writing punch option length: %v", err)
			}
			if _, err := buffer.Write([]byte(item.PunchOption)); err != nil {
				return fmt.Errorf("error writing punch option: %v", err)
			}
		}
		if itemInfo.ItemVersion >= 12 {
			if _, err := buffer.Write(item.Data12); err != nil {
				return fmt.Errorf("error writing version 12 specific bytes: %v", err)
			}
		}
		if itemInfo.ItemVersion >= 13 {
			if err := binary.Write(buffer, binary.LittleEndian, item.IntData13); err != nil {
				return fmt.Errorf("error writing version 13 specific bytes: %v", err)
			}
		}
		if itemInfo.ItemVersion >= 14 {
			if err := binary.Write(buffer, binary.LittleEndian, item.IntData14); err != nil {
				return fmt.Errorf("error writing version 14 specific bytes: %v", err)
			}
		}
		if itemInfo.ItemVersion >= 15 {
			if _, err := buffer.Write(item.Data15); err != nil {
				return fmt.Errorf("error writing version 15 specific bytes: %v", err)
			}
			strLen := uint16(len(item.StrData15))
			if err := binary.Write(buffer, binary.LittleEndian, strLen); err != nil {
				return fmt.Errorf("error writing strData15 length: %v", err)
			}
			if _, err := buffer.Write([]byte(item.StrData15)); err != nil {
				return fmt.Errorf("error writing strData15: %v", err)
			}
		}
		if itemInfo.ItemVersion >= 16 {
			strLen := uint16(len(item.StrData16))
			if err := binary.Write(buffer, binary.LittleEndian, strLen); err != nil {
				return fmt.Errorf("error writing strData16 length: %v", err)
			}
			if _, err := buffer.Write([]byte(item.StrData16)); err != nil {
				return fmt.Errorf("error writing strData16: %v", err)
			}
		}
		if itemInfo.ItemVersion >= 17 {
			if err := binary.Write(buffer, binary.LittleEndian, item.IntData17); err != nil {
				return fmt.Errorf("error writing strData17: %v", err)
			}
		}
		if itemInfo.ItemVersion >= 18 {
			if err := binary.Write(buffer, binary.LittleEndian, item.IntData18); err != nil {
				return fmt.Errorf("error writing strData18: %v", err)
			}
		}
	}

	// Write the buffer to a file
	outputFile, err := os.Create(pathFile)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()

	if _, err := buffer.WriteTo(outputFile); err != nil {
		return fmt.Errorf("error writing to output file: %v", err)
	}
	log.Info("items.dat successfully encoded for %s", time.Since(now))

	return nil
}

func main() {
	now := time.Now()
	decoded, err := DecodeItemsDat("items.dat", now)
	if err != nil {
		log.Error("Error ketika mencoba decode \"items.dat\": %v", err)
	}
	data, _ := json.MarshalIndent(decoded, "", " ")
	_ = ioutil.WriteFile("items.json", data, 0644)
	now = time.Now()
	err = EncodeItemsDat(decoded, "items.encodedd.dat", now)
	if err != nil {
		log.Error("Error ketika mencoba encode \"items.dat\": %v", err)
	}
}
