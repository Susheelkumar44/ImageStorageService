/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc JSON-to-DB-Interaction Models
 */

package model

import "time"

//ImageStorageModel Image DB Model
type ImageStorageModel struct {
	ImageID           int       `json:"imageId" gorm:"column:ImageID;primary_key;AUTO_INCREMENT"`
	ImageName         string    `json:"imageName" gorm:"column:ImageName"`
	ImageBase64Format string    `json:"imageBase64Format" gorm:"column:ImageBase64Format"`
	ImageCreatedAt    time.Time `json:"createdAt" gorm:"column:CreatedTime"`
}

//ImageStorageModelStruct Input Bound JSON Structre
type ImageStorageModelStruct struct {
	ImageAlbumName    string    `json:imageAlbumName`
	ImageName         string    `json:"imageName"`
	ImageBase64Format string    `json:"imageBase64Format"`
	ImageCreatedAt    time.Time `json:"createdAt"`
}

//APIResponse struct defines structure for the responses generated
type APIResponse struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

//APIResponseCreate struct defines structure for the responses generated for post Operation
type APIResponseCreate struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

//Data struct defines structure for the responses generated for create Operation containing the ImageID
type Data struct {
	ImageID int `json:"imageID"`
}
