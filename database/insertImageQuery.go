/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc Insert Image Query
 */

package database

import (
	model "ImageStorageService/model"
)

//InsertImageQuery function inserts the images
func (dc *DBRepo) InsertImageQuery(imageModel model.ImageStorageModel, imageAlbumName string) (model.ImageStorageModel, error) {

	err := dc.GormDB.Debug().Table(imageAlbumName).Create(&imageModel).Error

	return imageModel, err

}
