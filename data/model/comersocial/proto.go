package comersocial

type ComerSocial struct {
	Id         uint64 `gorm:"column:id primary_key" json:"id"`
	ComerId    uint64 `gorm:"column:comer_id" json:"comer_id"`
	Platform   string `gorm:"column:platform" json:"platform"`
	Username   string `gorm:"column:username" json:"username"`
	Url        string `gorm:"column:url" json:"url"`
	IsVerified bool   `gorm:"column:is_verified" json:"is_verified"`
	CreatedAt  string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  string `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted  int    `gorm:"column:is_deleted" json:"is_deleted"`
}

func (ComerSocial) TableName() string {
	return "comer_social"
}
