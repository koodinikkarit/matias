package ewdatabasemodels

type Word struct {
	Rowid  uint32
	SongId uint32
	Words  string
}

func (word Word) TableName() string {
	return "word"
}
