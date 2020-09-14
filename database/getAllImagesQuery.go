/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc Get All Images Query
 */

package database

import (
	model "ImageStorageService/model"
)

//GetAllImagesQuery function for get all Images
func (dc *DBRepo) GetAllImagesQuery(imageAlbumName string) ([]model.ImageStorageModel, int64) {

	allImages := []model.ImageStorageModel{}

	rows := dc.GormDB.Table(imageAlbumName).Debug().Find(&allImages).RowsAffected

	return allImages, rows

}
