package delivery

import (
	"mygram-byferdiansyah/domain"
	"mygram-byferdiansyah/helpers"
	"mygram-byferdiansyah/image/delivery/http/middleware"
	"mygram-byferdiansyah/image/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type imageHandler struct {
	imageUseCase domain.ImageUseCase
}

func NewImageHandler(routers *gin.Engine, imageUseCase domain.ImageUseCase) {
	handler := &imageHandler{imageUseCase}

	router := routers.Group("/images")
	{
		router.Use(middleware.Authentication())
		router.GET("", handler.Get)
		router.POST("", handler.Create)
		router.PUT("/:imageId", middleware.Authorization(handler.imageUseCase), handler.Edit)
		router.DELETE("/:imageId", middleware.Authorization(handler.imageUseCase), handler.Delete)
	}
}

// Get godoc
// @Summary    	Get all images
// @Description	Get all images with authentication user
// @Tags        images
// @Accept      json
// @Produce     json
// @Success     200			{object}	utils.ResponseDataGetedImage
// @Failure     400			{object}	utils.ResponseMessage
// @Failure     401			{object}	utils.ResponseMessage
// @Security    Bearer
// @Router      /images	[get]
func (handler *imageHandler) Get(ctx *gin.Context) {
	var (
		images []domain.Image
		err    error
	)

	if err = handler.imageUseCase.Get(ctx.Request.Context(), &images); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	getsImages := []*utils.GetedImage{}

	for _, image := range images {
		getsImages = append(getsImages, &utils.GetedImage{
			ID:        image.ID,
			Title:     image.Title,
			Caption:   image.Caption,
			ImageUrl:  image.ImageUrl,
			UserID:    image.UserID,
			CreatedAt: image.CreatedAt,
			UpdatedAt: image.UpdatedAt,
			User: &utils.User{
				Email:    image.User.Email,
				Username: image.User.Username,
			},
		})
	}

	ctx.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data:   getsImages,
	})
}

// Create godoc
// @Summary    	Create a image
// @Description	Create and create a image with authentication user
// @Tags        images
// @Accept      json
// @Produce     json
// @Param       json		body			utils.AddImage	true	"Add Image"
// @Success     201			{object}  utils.ResponseDataAddedImage
// @Failure     400			{object}	utils.ResponseMessage
// @Failure     401			{object}	utils.ResponseMessage
// @Security    Bearer
// @Router      /images	[post]
func (handler *imageHandler) Create(ctx *gin.Context) {
	var (
		image domain.Image
		err   error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&image); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	image.UserID = userID

	if err = handler.imageUseCase.Create(ctx.Request.Context(), &image); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, helpers.ResponseData{
		Status: "success",
		Data: utils.AddedImage{
			ID:        image.ID,
			Title:     image.Title,
			Caption:   image.Caption,
			ImageUrl:  image.ImageUrl,
			UserID:    image.UserID,
			CreatedAt: image.CreatedAt,
		},
	})
}

// Edit godoc
// @Summary     Edit a image
// @Description	Edit a image by id with authentication user
// @Tags        images
// @Accept      json
// @Produce     json
// @Param       id		path      string	true	"Image ID"
// @Param       json	body			utils.EditImage true  "Image"
// @Success     200		{object}  utils.ResponseDataEditedImage
// @Failure     400		{object}	utils.ResponseMessage
// @Failure     401		{object}	utils.ResponseMessage
// @Failure     404		{object}	utils.ResponseMessage
// @Security    Bearer
// @Router      /images/{id}		[put]
func (handler *imageHandler) Edit(ctx *gin.Context) {
	var (
		image domain.Image
		err   error
	)

	if err = ctx.ShouldBindJSON(&image); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	editedImage := domain.Image{
		Title:    image.Title,
		Caption:  image.Caption,
		ImageUrl: image.ImageUrl,
	}

	imageID := ctx.Param("imageId")

	if image, err = handler.imageUseCase.Edit(ctx.Request.Context(), editedImage, imageID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data: utils.EditedImage{
			ID:        image.ID,
			UserID:    image.UserID,
			Title:     image.Title,
			ImageUrl:  image.ImageUrl,
			Caption:   image.Caption,
			UpdatedAt: image.UpdatedAt,
		},
	})
}

// Delete godoc
// @Summary     Delete a image
// @Description	Delete a image by id with authentication user
// @Tags        images
// @Accept      json
// @Produce     json
// @Param       id	path			string	true	"Image ID"
// @Success     200	{object}	utils.ResponseMessageDeletedImage
// @Failure     400	{object}	utils.ResponseMessage
// @Failure     401	{object}	utils.ResponseMessage
// @Failure     404	{object}	utils.ResponseMessage
// @Security    Bearer
// @Router      /images/{id}	[delete]
func (handler *imageHandler) Delete(ctx *gin.Context) {
	imageID := ctx.Param("imageId")

	if err := handler.imageUseCase.Delete(ctx.Request.Context(), imageID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "your image has been successfully deleted",
	})
}
