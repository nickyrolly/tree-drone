package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Estate struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id"`
	Length int       `gorm:"type:int;column:length"`
	Width  int       `gorm:"type:int;column:width"`
}

func (e *Estate) TableName() string {
	return "estate"
}

// GORM Expressions
func (e *Estate) BeforeCreate(tx *gorm.DB) error {
	e.ID = uuid.New()
	return nil
}

func (e *Estate) SetLength(length int) {
	e.Length = length
}

func (e *Estate) GetLength() int {
	return e.Length
}

func (e *Estate) SetWidth(width int) {
	e.Width = width
}

func (e *Estate) GetWidth() int {
	return e.Width
}

type EstateTree struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id"`
	Height int       `gorm:"type:int;column:height"`
	X      int       `gorm:"type:int;column:x"`
	Y      int       `gorm:"type:int;column:y"`
}

func (e *EstateTree) BeforeCreate(tx *gorm.DB) error {
	e.ID = uuid.New()
	return nil
}

func (e *EstateTree) TableName() string {
	return "estate_tree"
}
