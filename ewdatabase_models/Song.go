package ewdatabasemodels

type Song struct {
	Rowid           uint32 `gorm:"AUTO_INCREMENT;primary_key"`
	SongItemUID     string `gorm:"unique"`
	SongRevUID      string
	SongUID         string
	Title           string
	Author          string
	Copyright       string
	Administrator   string
	Description     string
	Tags            string
	ReferenceNumber int
	VendorID        int
	PresentationID  int
	LayoutRevision  int
	Revision        int
}

func (song Song) TableName() string {
	return "song"
}
