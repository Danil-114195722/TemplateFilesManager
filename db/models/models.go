package models


// модель для файлов
type File struct {
	ID 		uint `gorm:"primarykey"`
	Name 	string `gorm:"type:varchar(255); not null; index:idx_composite_unique_name_tag,unique"`
	Tag		string `gorm:"type:varchar(255); not null; index:idx_composite_unique_name_tag,unique"`
}
