package tag

import (
	"time"
)

type Category string

const (
	ComerSkill Category = "comerSkill"
	Startup    Category = "startup"
	Bounty     Category = "bounty"
)

// TypeToCategory 将 type 值转换为对应的 Category
func TypeToCategory(typeValue int) Category {
	switch typeValue {
	case 1:
		return ComerSkill
	case 2:
		return Startup
	case 3:
		return Bounty
	default:
		return ""
	}
}

// Tag  Comunion tag for startup bounty profile and other position need Tag.
type Tag struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string   `gorm:"column:name" json:"name"`
	Category  Category `gorm:"column:category" json:"category"`
	IsIndex   bool     `gorm:"column:is_index" json:"isIndex"`
	IsDeleted bool     `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName identify the table name of this model.
func (Tag) TableName() string {
	return "tag"
}

// TagTargetRel  Comunion tag for startup bounty profile and other position need TagTargetRel.
type TagTargetRel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	TargetID  uint64   `column:"target_id"`
	Target    Category `column:"target"`
	TagID     uint64   `column:"tag_id"`
}

// TableName identify the table name of this model.
func (TagTargetRel) TableName() string {
	return "tag_target_rel"
}
