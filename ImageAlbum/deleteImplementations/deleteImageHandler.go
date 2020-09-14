/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc Delete Image Router
 */

package deleteImplementations

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"ImageStorageService/interfaces"
	util "ImageStorageService/util"
)

//Constants Declarations
const (
	successfulOperation  string = "Successful operation"
	imageAlbumNameEmpty  string = "Image Album Name is Empty"
	imageIDEmpty         string = "Image ID Cannot be empty"
	imageDeletionError   string = "Error in deleting Image"
	imageNotPresentError string = "Image Does not exist"
)

//DeleteHandler Binder
type DeleteHandler struct {
}

//DeleteImage Router
func (deleteApi *DeleteHandler) DeleteImage(context echo.Context) error {
	imageAlbumName := context.Param("imageAlbumName")
	imageId, errConv := strconv.Atoi(context.Param("imageID"))

	defer context.Request().Body.Close()
	context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	if imageAlbumName == "" {
		return util.ResponseMapper(http.StatusMethodNotAllowed, imageAlbumNameEmpty, context)
	}

	if errConv != nil || imageId <= 0 {
		return util.ResponseMapper(http.StatusMethodNotAllowed, imageAlbumNameEmpty, context)
	}

	_, rowsAffected := interfaces.DBClient.GetSingleImageQuery(imageId, imageAlbumName)
	if rowsAffected == 0 {
		return util.ResponseMapper(http.StatusNotFound, imageNotPresentError, context)
	}

	err := interfaces.DBClient.DeleteImageQuery(imageId, imageAlbumName)
	if err != nil {
		return util.ResponseMapper(http.StatusMethodNotAllowed, imageDeletionError, context)
	}

	return util.ResponseMapper(http.StatusOK, successfulOperation, context)

}
