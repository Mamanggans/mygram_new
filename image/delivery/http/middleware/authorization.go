package middleware

import (
	"fmt"
	"mygram-byferdiansyah/domain"
	"mygram-byferdiansyah/helpers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization(imageUseCase domain.ImageUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			image domain.Image
			err   error
		)

		imageID := ctx.Param("imageId")
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := string(userData["id"].(string))

		if err = imageUseCase.GetByID(ctx.Request.Context(), &image, imageID); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helpers.ResponseMessage{
				Status:  "fail",
				Message: fmt.Sprintf("image with id %s doesn't exist", imageID),
			})

			return
		}

		if image.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.ResponseMessage{
				Status:  "unauthorized",
				Message: "you don't have permission to view or edit this image",
			})

			return
		}
	}
}
