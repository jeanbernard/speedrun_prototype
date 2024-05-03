package models

type Run struct {
	Id         string `gorm:"primaryKey"`
	GameID     string
	CategoryID string
	PlayerId   string
	//Variables  []Variable
	Level    *string
	VideoURI string
	Runtime  float64
}
