/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc Get All Images Router
 */

package getImplementations

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"

	"ImageStorageService/interfaces"
	util "ImageStorageService/util"
)

//GetAllImages Router Handler
func (getAllApi *GetHandler) GetAllImages(context echo.Context) error {
	imageAlbumName := context.Param("imageAlbumName")

	defer context.Request().Body.Close()
	context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	if imageAlbumName == "" {
		return util.ResponseMapper(http.StatusMethodNotAllowed, imageAlbumNameEmpty, context)
	}

	allImages, rowsAffected := interfaces.DBClient.GetAllImagesQuery(imageAlbumName)
	if rowsAffected == 0 {
		return util.ResponseMapper(http.StatusMethodNotAllowed, imageGetError, context)
	}

	//Encoding and sending the recorded response
	context.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(context.Response()).Encode(allImages)

}
