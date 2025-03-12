package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Estate struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id"`
	length int       `gorm:"type:int;column:length"`
	width  int       `gorm:"type:int;column:width"`
}

type EstateTree struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id"`
	Height int       `gorm:"type:int;column:height"`
	X      int       `gorm:"type:int;column:x"`
	Y      int       `gorm:"type:int;column:y"`
}

func (e *Estate) BeforeCreate(tx *gorm.DB) error {
	e.ID = uuid.New()
	return nil
}

func (e *EstateTree) BeforeCreate(tx *gorm.DB) error {
	e.ID = uuid.New()
	return nil
}

func (e *Estate) TableName() string {
	return "estate"
}

func (e *EstateTree) TableName() string {
	return "estate_tree"
}
