package models

type Province struct {
	Code                 string `gorm:"column:code;primaryKey"`
	Name                 string `gorm:"column:name"`
	NameEn               string `gorm:"column:name_en"`
	FullName             string `gorm:"column:full_name"`
	FullNameEn           string `gorm:"column:full_name_en"`
	CodeName             string `gorm:"column:code_name"`
	AdministrativeUnitID *int   `gorm:"column:administrative_unit_id"`
}

func (Province) TableName() string {
	return "provinces"
}
