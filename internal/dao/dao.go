package dao

import "gorm.io/gorm"

type Dao struct {
	DB *gorm.DB
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		DB: db,
	}
}


func (d *Dao) 