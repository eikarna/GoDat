package Decoder

import (
	"encoding/binary"
	"errors"
	"fmt"
	. "github.com/eikarna/GoDat/Components/Enums"
	"github.com/eikarna/GoDat/Components/ProtonHash"
	"os"
	"time"
)

func Decode(pathFile string, timestamp time.Time) (*ItemInfo, error) {
	if pathFile == "" {
		return nil, errors.New("Please provide the target file!")
	}
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
	itemInfo.FileHash = ProtonHash.GetFileHash(pathFile)
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
			itemInfo.Items[i].ActionType = uint8(data[memPos])
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
	data, err = nil, nil
	fmt.Printf("Items.dat decoded for %s. With Item Count: %d, ItemsDatVersion: %d, Item Hash: %v\n", time.Since(timestamp), itemInfo.ItemCount, itemInfo.ItemVersion, itemInfo.FileHash)
	return itemInfo, nil
}

func FileBuffer(pathFile string) ([]byte, error) {
	if pathFile == "" {
		return nil, errors.New("Please provide the target file!")
	}
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
	data := make([]byte, size)
	_, err = file.Read(data)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	FileBufferPacket := make([]byte, 60+size)
	binary.LittleEndian.PutUint32(FileBufferPacket[0:], 4)
	binary.LittleEndian.PutUint32(FileBufferPacket[4:], 16)
	binary.LittleEndian.PutUint32(FileBufferPacket[8:], ^uint32(0))
	binary.LittleEndian.PutUint32(FileBufferPacket[16:], 8)
	binary.LittleEndian.PutUint32(FileBufferPacket[56:], uint32(size))
	copy(FileBufferPacket[60:], data)
	return FileBufferPacket, nil
}
