/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc Create Image Router
 */
package postImplementations

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo"

	"ImageStorageService/interfaces"
	model "ImageStorageService/model"
	util "ImageStorageService/util"
)

//Constants Declaration
const (
	decodeJSONErr       string = "Invalid Input: Error in Decoding Input JSON"
	successfulOperation string = "Successful operation"
	imageEmpty          string = "Image can't be Empty"
	imageCreationError  string = "Error in creating Image"
)

//PostHandler binder
type PostHandler struct{}

//InsertImage router handler to insert image query
func (postApi *PostHandler) InsertImage(context echo.Context) error {
	imageStorageModel := model.ImageStorageModelStruct{}
	imageModel := model.ImageStorageModel{}

	//JSON Decoder
	errDecode := json.NewDecoder(context.Request().Body).Decode(&imageStorageModel)

	defer context.Request().Body.Close()
	context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	//validating the format of JSON provided
	if errDecode != nil {
		return util.ResponseMapper(http.StatusMethodNotAllowed, decodeJSONErr, context)
	}

	if (imageStorageModel.ImageBase64Format == "") && (imageStorageModel.ImageName == "") {
		return util.ResponseMapper(http.StatusMethodNotAllowed, imageEmpty, context)
	}

	imageModel.ImageName = imageStorageModel.ImageName
	imageModel.ImageBase64Format = imageStorageModel.ImageBase64Format
	imageModel.ImageCreatedAt = time.Now().UTC()

	//Database call
	imageData, err := interfaces.DBClient.InsertImageQuery(imageModel, imageStorageModel.ImageAlbumName)
	if err != nil {
		return util.ResponseMapper(http.StatusMethodNotAllowed, imageCreationError, context)
	}

	return util.CreateResponseMapper(http.StatusOK, successfulOperation, imageData.ImageID, context)

}
