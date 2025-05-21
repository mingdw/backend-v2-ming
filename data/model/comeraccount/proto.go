package comeraccount

import "time"

type ComerAccountType int
type ComerAccount struct {
	Id        uint64           `gorm:"primarykey" db:"id"`
	ComerID   uint64           `gorm:"column:comer_id;index" json:"comer_id"`
	Oin       string           `gorm:"column:oin;uniqueIndex;not null" json:"oin"`
	IsPrimary bool             `gorm:"column:is_primary;not null" json:"is_primary"`
	Nick      string           `gorm:"column:nick;not null" json:"nick"`
	Avatar    string           `gorm:"column:avatar;not null" json:"avatar"`
	Type      ComerAccountType `gorm:"column:type;not null" json:"type"`
	IsLinked  bool             `gorm:"column:is_linked;not null" json:"is_linked"`
	CreatedAt time.Time        `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time        `gorm:"column:updated_at;not null" json:"updated_at"`
	IsDeleted bool             `gorm:"column:is_deleted;not null" json:"is_deleted"`
}

// TableName ComerAccount table name for gorm
func (ComerAccount) TableName() string {
	return "comer_account"
}
