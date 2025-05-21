package comerprofile

import (
	"gorm.io/gorm"
)

func FindComerProfile(db *gorm.DB, comerID uint64) (comerProfile *ComerProfile, err error) {
	err = db.Where("comer_id = ? and is_deleted = 0", comerID).First(&comerProfile).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func UpdateComerProfile(db *gorm.DB, comerID uint64, updates map[string]interface{}) (err error) {
	err = db.Model(&ComerProfile{}).Where("comer_id = ? and is_deleted = 0", comerID).Updates(updates).Error
	if err != nil {
		return err
	}
	return
}

func InsertComerProfile(db *gorm.DB, comerProfile *ComerProfile) (err error) {
	err = db.Create(comerProfile).Error
	return
}
