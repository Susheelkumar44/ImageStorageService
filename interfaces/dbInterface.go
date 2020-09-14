/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc DB Interface methods
 */
package interfaces

import (
	config "ImageStorageService/config"
	model "ImageStorageService/model"
)

//DBClient - DBInteractions Interface Object
var DBClient DBInteractions

//DBInteractions Interface
type DBInteractions interface {
	DBConnect(config config.DBConfig) error
	InsertImageQuery(imageModel model.ImageStorageModel, imageAlbumName string) (model.ImageStorageModel, error)
	DeleteImageQuery(imageID int, imageAlbumName string) error
	GetSingleImageQuery(imageID int, imageAlbumName string) (model.ImageStorageModel, int64)
	GetAllImagesQuery(imageAlbumName string) ([]model.ImageStorageModel, int64)
}
