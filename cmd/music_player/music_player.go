package main

import (
	"github.com/calmius/music_player/cmd/music_player/handlers"
	"github.com/calmius/music_player/internal/platform/postgres"
	"github.com/gin-gonic/gin"
)

func init() {
	postgres.InitDB()
}

func main() {
	r := gin.Default()
	r.POST("/", handlers.CreateSong)
	r.GET("/", handlers.GetSongs)
	// contacts?limit=10&offset=20 ‍
	r.GET("/:id", handlers.GetSong)
	r.PUT("/:id", handlers.UpdateSong)
	r.DELETE("/:id", handlers.DeleteSong)
	r.Run()
}
