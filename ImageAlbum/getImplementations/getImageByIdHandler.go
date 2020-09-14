/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc Get Single Image Router
 */

package getImplementations

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"ImageStorageService/interfaces"
	util "ImageStorageService/util"
)

//Constant Declarations
const (
	imageIDEmpty        string = "Image ID Cannot be empty"
	imageGetError       string = "Error in getting Image"
	imageAlbumNameEmpty string = "Image Album Name is Empty"
)

//Get Handler Binder
type GetHandler struct {
}

//GetSingleImage Router Handler to GetSingleImage Query
func (getByIdApi *GetHandler) GetSingleImage(context echo.Context) error {

	//Path Parameter Extractor
	imageId, errConv := strconv.Atoi(context.Param("imageID"))
	imageAlbumName := context.Param("imageAlbumName")

	defer context.Request().Body.Close()
	context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	if errConv != nil || imageId <= 0 {
		return util.ResponseMapper(http.StatusMethodNotAllowed, imageIDEmpty, context)
	}

	if imageAlbumName == "" {
		return util.ResponseMapper(http.StatusMethodNotAllowed, imageAlbumNameEmpty, context)
	}

	//Database call
	singleImage, rowsAffected := interfaces.DBClient.GetSingleImageQuery(imageId, imageAlbumName)
	if rowsAffected == 0 {
		return util.ResponseMapper(http.StatusMethodNotAllowed, imageGetError, context)
	}

	//Encoding and sending the recorded response
	context.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(context.Response()).Encode(singleImage)

}
