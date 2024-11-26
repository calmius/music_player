package handlers

import (
	"fmt"

	"github.com/calmius/music_player/cmd/music_player/helpers"
	"github.com/calmius/music_player/internal/model"
	"github.com/calmius/music_player/internal/platform/postgres"
	"github.com/gin-gonic/gin"
)

func CreateSong(ctx *gin.Context) {
	var body struct {
		GroupName string
		SongName  string
	}
	ctx.BindJSON(&body)
	// ctx.DefaultQuery

	song := model.Songs{GroupName: body.GroupName, SongName: body.SongName}
	fmt.Println(song)

	result := postgres.DB.Create(&song)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}
	ctx.JSON(200, gin.H{"data": song})
}

func GetSongs(ctx *gin.Context) {
	var songs []model.Songs
	// result := postgres.DB.Find(&songs)
	result := postgres.DB.Scopes(helpers.Paginate(ctx)).Find(&songs)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}
	ctx.JSON(200, gin.H{"data": songs})
}

func GetSong(ctx *gin.Context) {
	var song model.Songs
	result := postgres.DB.First(&song, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": song})
}
func UpdateSong(ctx *gin.Context) {
	var body struct {
		GroupName string
		SongName  string
	}

	ctx.BindJSON(&body)

	var song model.Songs

	result := postgres.DB.First(&song, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	postgres.DB.Model(&song).Updates(model.Songs{GroupName: body.GroupName, SongName: body.SongName})

	ctx.JSON(200, gin.H{"data": song})
}

func DeleteSong(ctx *gin.Context) {

	id := ctx.Param("id")

	postgres.DB.Delete(&model.Songs{}, id)
	ctx.JSON(200, gin.H{"data": "song has been deleted successfully"})
}
