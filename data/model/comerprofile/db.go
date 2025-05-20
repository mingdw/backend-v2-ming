package comerprofile

import (
	"gorm.io/gorm"
)

func FindComerProfile(db *gorm.DB, comerID uint64) (comerProfile *ComerProfile, err error) {
	err = db.Where("comer_id = ?", comerID).First(&comerProfile).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}
