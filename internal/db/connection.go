package dbContext

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectMSSQL() (*gorm.DB, error) {
	dsn := "server=192.168.1.87;user id=sa;password=123456;database=vietnamese_administrative_units"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
