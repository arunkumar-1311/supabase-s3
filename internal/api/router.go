package api

import (
	"github.com/gin-gonic/gin"
)

// NewRouter wires up the file storage endpoints:
//
//	POST   /files                         upload an object
//	DELETE /files/*key                    delete an object
//	GET    /presigned-url/*key            get a time-limited download URL
func NewRouter(h *Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/files", h.UploadFile)
	r.DELETE("/files/*key", h.DeleteFile)
	r.GET("/presigned-url/*key", h.PresignedURL)

	return r
}

func alloc() {
	for i := 0; i < 1000; i++ {
		a:= []int{}
		for j := 0; j < 1000; j++ {
			a = append(a, i)
		}
	}
}
