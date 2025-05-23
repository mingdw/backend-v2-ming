package comerlanguage

type ComerLanguage struct {
	Id        uint64 `gorm:"column:id primary_key" json:"id"`
	ComerId   uint64 `gorm:"column:comer_id" json:"comer_id"`
	Language  string `gorm:"column:language" json:"language"`
	Code      string `gorm:"column:code" json:"code"`
	Level     int    `gorm:"column:level" json:"level"`
	IsNative  bool   `gorm:"column:is_native" json:"is_native"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted int    `gorm:"column:is_deleted" json:"is_deleted"`
}

func (ComerLanguage) TableName() string {
	return "comer_language"
}
