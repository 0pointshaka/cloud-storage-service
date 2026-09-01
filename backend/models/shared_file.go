package models

type SharedFile struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FileID    uint   `gorm:"not null" json:"file_id"`
	File      File   `gorm:"foreignKey:FileID" json:"-"`
	SharedBy  uint   `gorm:"not null" json:"shared_by"`
	SharedTo  uint   `gorm:"not null" json:"shared_to"`
	User      User   `gorm:"foreignKey:SharedTo" json:"-"`
	CreatedAt int64  `json:"created_at"`
}
