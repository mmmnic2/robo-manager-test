package repository

import (
	"gorm.io/gorm"
)

type RobotRepository interface {
	DB() *gorm.DB
}

type robotRepository struct {
	Db *gorm.DB
}

func (r robotRepository) DB() *gorm.DB {
	return r.Db
}

func NewRobotRepository(db *gorm.DB) RobotRepository {
	return &robotRepository{Db: db}
}
