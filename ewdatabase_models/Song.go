package ewdatabasemodels

type Song struct {
	Rowid           uint32 `gorm:"AUTO_INCREMENT;primary_key"`
	SongItemUID     string `gorm:"column:song_item_uid"`
	SongRevUID      string `gorm:"column:song_rev_uid"`
	SongUID         string `gorm:"column:song_uid"`
	Title           string
	Author          string
	Copyright       string
	Administrator   string
	Description     string
	Tags            string
	ReferenceNumber string
	VendorID        int
	PresentationID  int
	LayoutRevision  int
	Revision        int
}

func (song Song) TableName() string {
	return "song"
}
