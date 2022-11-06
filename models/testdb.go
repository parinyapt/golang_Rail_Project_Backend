package models

type Language struct {
	LanguageID   uint `gorm:"column:language_id" json:"language_id"`
	LanguageCode string `gorm:"column:language_code" json:"language_code"`
	LanguageName string `gorm:"column:language_name" json:"language_name"`
}

func (Language) TableName() string {
	return "crp_language"
}

// type ResponseLanguage struct {
// 	LanguageID   uint `json:"language_id"`
// 	LanguageCode string `json:"language_code"`
// 	LanguageName string `json:"language_name"`
// }