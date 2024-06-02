package Encoder

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	. "github.com/eikarna/GoDat/Components/Enums"
	"os"
)

func Encode(itemInfo *ItemInfo, pathFile string) error {
	if itemInfo == nil {
		return errors.New("Please provide the items data to be encoded!")
	}
	if pathFile == "" {
		pathFile = "items.dat"
	}
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
			// if err := buffer.WriteByte(byte(item.ActionType)); err != nil {
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

		// Encode extraFileHash
		if err := binary.Write(buffer, binary.LittleEndian, item.ExtraFileHash); err != nil {
			return fmt.Errorf("error writing extra file hash: %v", err)
		}

		// Encode AudioVolume
		if err := binary.Write(buffer, binary.LittleEndian, item.AudioVolume); err != nil {
			return fmt.Errorf("error writing audio volume: %v", err)
		}

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

		if err := binary.Write(buffer, binary.LittleEndian, item.SeedColorR); err != nil {
			return fmt.Errorf("error writing SeedColorR: %v", err)
		}
		if err := binary.Write(buffer, binary.LittleEndian, item.SeedColorG); err != nil {
			return fmt.Errorf("error writing SeedColorG: %v", err)
		}
		if err := binary.Write(buffer, binary.LittleEndian, item.SeedColorB); err != nil {
			return fmt.Errorf("error writing SeedColorB: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.SeedOverlayColorA); err != nil {
			return fmt.Errorf("error writing SeedOverlayColorA: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.SeedOverlayColorR); err != nil {
			return fmt.Errorf("error writing SeedOverlayColorR: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.SeedOverlayColorG); err != nil {
			return fmt.Errorf("error writing SeedOverlayColorG: %v", err)
		}

		if err := binary.Write(buffer, binary.LittleEndian, item.SeedOverlayColorB); err != nil {
			return fmt.Errorf("error writing SeedOverlayColorB: %v", err)
		}

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
		if _, err := buffer.Write([]byte(item.DataPos80)); err != nil {
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

	if _, err := buffer.WriteTo(outputFile); err != nil {
		return fmt.Errorf("error writing to output file: %v", err)
	}

	outputFile.Close()
	itemInfo, buffer, err = nil, nil, nil
	// fmt.Printf("items.dat successfully encoded for %s\n", time.Since(now))
	return nil
}
