package comerskill

import (
	"time"
)

type ComerSkill struct {
	Id          uint64    `gorm:"column:id;primary_key;auto_increment"  json:"id"`
	ComerId     uint64    `gorm:"column:comer_id" json:"comer_id"`
	SkillName   string    `gorm:"column:skill_name" json:"skill_name"`
	Level       int       `gorm:"column:level" json:"level"`
	Years       int       `gorm:"column:years" json:"years"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted   bool      `gorm:"column:is_deleted" json:"is_deleted"`
}

func (ComerSkill) TableName() string {
	return "comer_skill"
}
