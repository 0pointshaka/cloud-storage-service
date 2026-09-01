package models

type File struct {
	ID        uint          `gorm:"primaryKey" json:"id"`
	UserID    uint          `gorm:"not null" json:"user_id"`
	User      User          `gorm:"foreignKey:UserID" json:"-"`
	Filename  string        `gorm:"not null" json:"filename"`
	Path      string        `gorm:"not null" json:"path"`
	Size      int64         `json:"size"`
	MimeType  string        `json:"mime_type"`
	Shared    []SharedFile  `gorm:"foreignKey:FileID" json:"-"`
	CreatedAt int64         `json:"created_at"`
	UpdatedAt int64         `json:"updated_at"`
}
