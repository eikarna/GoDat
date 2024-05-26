package Enums

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
	ActionType         uint8
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
	FileHash int32
}

const (
	key = "PBG892FXX982ABC*"
)

func (Info *ItemInfo) GetItemHash() int32 {
	return int32(Info.FileHash)
}
