/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc Microservice Initalizer
 */

package main

import (
	"log"

	"github.com/labstack/echo"

	apiServices "ImageStorageService/ImageAlbum"
	config "ImageStorageService/config"
	"ImageStorageService/database"
	"ImageStorageService/interfaces"
)

//-----------------------------------------------------------------------------------------------------------------------//

//Main defines the point where execution starts
func main() {
	initGormClient()
	e := echo.New()

	api := apiServices.Initalization()
	g := e.Group("/v1") // added to setup initial group of api's

	g.DELETE("/image/:imageID/:imageAlbumName", api.DeleteHandler.DeleteImage)
	g.GET("/image/:imageID/:imageAlbumName", api.GetHandler.GetSingleImage)
	g.GET("/image/:imageAlbumName", api.GetHandler.GetAllImages)
	g.POST("/image", api.PostHandler.InsertImage)

	e.Logger.Fatal(e.Start(":8051"))
}

//-----------------------------------------------------------------------------------------------------------------------//

//initGormClient initializes the database
func initGormClient() {
	config, _ := config.LoadConfig()
	interfaces.DBClient = new(database.DBRepo)
	err := interfaces.DBClient.DBConnect(config)
	if err != nil {
		log.Println(err)
	}
}

//-----------------------------------------------------------------------------------------------------------------------//
