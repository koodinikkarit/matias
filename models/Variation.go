package models

type Variation struct {
	ID          uint32
	ServerID    uint32
	Name        string
	Text        string
	Version     uint32
	LanguageID  uint32
	AuthorID    uint32
	CopyrightID uint32
	SongID      uint32
}
