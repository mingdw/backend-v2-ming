package comeraccount

import "gorm.io/gorm"

func FindComerAccount(db *gorm.DB, comerAccountID uint64) (comerAccount *ComerAccount, err error) {
	err = db.Where("id = ? and is_deleted = 0", comerAccountID).First(&comerAccount).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func FindComerAccountByComerID(db *gorm.DB, comerID uint64) (comerAccount *ComerAccount, err error) {
	err = db.Where("comer_id = ? and is_deleted = 0", comerID).First(&comerAccount).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func UpdateComerAccount(db *gorm.DB, comerAccount *ComerAccount) (err error) {
	err = db.Model(&ComerAccount{}).Where("id = ?", comerAccount.Id).Updates(comerAccount).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return
}
