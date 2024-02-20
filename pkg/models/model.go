package models

type CharacterData struct {
	Id        int64  `gorm:"primaryKey;column:id;type:INT"`
	Character string `gorm:"column:character;type:VARCHAR(255)"`
}
type GetCharacterData struct {
	Character string `json:"character"`
}
