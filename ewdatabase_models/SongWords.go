package ewdatabasemodels

type Word struct {
	Rowid                uint32  `gorm:"primary_key;AUTO_INCREMENT"`
	SongID               uint32  `gorm:"unique"`
	Words                string  `gorm:"type:rtf"`
	SlideUids            *string `gorm:"column:slide_uids;type:text"`
	SlideLayoutRevisions *int64  `gorm:"type:int64a"`
	SlideRevisions       *int64  `gorm:"type:int64a"`
}

func (word Word) TableName() string {
	return "word"
}
