/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc Delete Image Query
 */

package database

import (
	"fmt"
)

//DeleteImageQuery function delete the images from the database
func (dc *DBRepo) DeleteImageQuery(imageID int, imageAlbumName string) error {

	query := fmt.Sprintf(`DELETE FROM "%s" WHERE "ImageID" = %d`, imageAlbumName, imageID)

	err := dc.GormDB.Debug().Exec(query).Error

	return err

}
