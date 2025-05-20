package comer

import (
	"time"
)

type Comer struct {
	Id        uint64    `gorm:"primarykey" db:"id"`
	Address   string    `gorm:"address" db:"address"` // comer could save some useful info on block chain with this address
	CreatedAt time.Time `gorm:"created_at" db:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" db:"updated_at"`
	IsDeleted int64     `gorm:"is_deleted" db:"is_deleted"` // Is Deleted
}

// TableName Startup table name for gorm
func (Comer) TableName() string {
	return "comer"
}
