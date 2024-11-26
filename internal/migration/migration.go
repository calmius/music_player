package main

import (
	"github.com/calmius/music_player/internal/model"
	"github.com/calmius/music_player/internal/platform/postgres"
)

func init() {
	postgres.InitDB()
}

func main() {
	postgres.DB.AutoMigrate(&model.Songs{})
}
