package models

type AdministrativeUnit struct {
	ID          int    `gorm:"column:id;primaryKey"`
	FullName    string `gorm:"column:full_name"`
	FullNameEn  string `gorm:"column:full_name_en"`
	ShortName   string `gorm:"column:short_name"`
	ShortNameEn string `gorm:"column:short_name_en"`
	CodeName    string `gorm:"column:code_name"`
	CodeNameEn  string `gorm:"column:code_name_en"`
}

func (AdministrativeUnit) TableName() string {
	return "administrative_regions"

}
