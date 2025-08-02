package models

type AdministrativeRegion struct {
	ID         int    `gorm:"column:id;primaryKey"`
	Name       string `gorm:"column:name"`
	NameEn     string `gorm:"column:name_en"`
	CodeName   string `gorm:"column:code_name"`
	CodeNameEn string `gorm:"column:code_name_en"`
}

func (AdministrativeRegion) TableName() string {
	return "administrative_regions"
}
