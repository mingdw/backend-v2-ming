package comereducation

type ComerEducation struct {
	Id          uint64 `gorm:"column:id primary_key" json:"id"`
	ComerId     uint64 `gorm:"column:comer_id" json:"comer_id"`
	School      string `gorm:"column:school" json:"school"`
	Degree      string `gorm:"column:degree" json:"degree"`
	Major       string `gorm:"column:major" json:"major"`
	StartDate   string `gorm:"column:start_date" json:"start_date"`
	EndDate     string `gorm:"column:end_date" json:"end_date"`
	Description string `gorm:"column:description" json:"description"`
	Level       int    `gorm:"column:level" json:"level"`
	CreatedAt   string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   string `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted   int    `gorm:"column:is_deleted" json:"is_deleted"`
}

func (ComerEducation) TableName() string {
	return "comer_education"
}
