/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc Get Single Image Query
 */

package database

import (
	model "ImageStorageService/model"
)

//GetSingleImageQuery - for single Image get
func (dc *DBRepo) GetSingleImageQuery(imageID int, imageAlbumName string) (model.ImageStorageModel, int64) {

	singleImage := model.ImageStorageModel{}

	rows := dc.GormDB.Table(imageAlbumName).Debug().Where(`"ImageID" = ?`, imageID).Find(&singleImage).RowsAffected

	return singleImage, rows

}
