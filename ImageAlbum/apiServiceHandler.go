/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc REST API Service Router
 */
package ImageAlbum

import (
	deleteService "ImageStorageService/ImageAlbum/deleteImplementations"
	getService "ImageStorageService/ImageAlbum/getImplementations"
	postService "ImageStorageService/ImageAlbum/postImplementations"
)

//APIServiceHandler REST API Method Handler structure
type APIServiceHandler struct {
	PostHandler   *postService.PostHandler
	DeleteHandler *deleteService.DeleteHandler
	GetHandler    *getService.GetHandler
}

//Initalization initalizes APIServiceHandler
func Initalization() *APIServiceHandler {
	APIServiceHandler := new(APIServiceHandler)
	APIServiceHandler.PostHandler = new(postService.PostHandler)
	APIServiceHandler.DeleteHandler = new(deleteService.DeleteHandler)
	APIServiceHandler.GetHandler = new(getService.GetHandler)
	return APIServiceHandler
}
