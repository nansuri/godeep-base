package base

type TmBaseContext struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	EnumType string `gorm:"size:1024;not null;" json:"enumType"`
	Value    string `gorm:"size:1024;not null;" json:"value"`
}
