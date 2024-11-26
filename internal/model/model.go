package model

type Songs struct {
	ID        uint `gorm:"primarykey"`
	GroupName string
	SongName  string
	Lyrics    string
	Link      string
}
