package delivery

import (
	"mygram-byferdiansyah/domain"
	"mygram-byferdiansyah/helpers"
	"mygram-byferdiansyah/socialmedia/delivery/http/middleware"
	"mygram-byferdiansyah/socialmedia/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type socialMediaHandler struct {
	socialMediaUseCase domain.SocialMediaUseCase
}

func NewSocialMediaHandler(routers *gin.Engine, socialMediaUseCase domain.SocialMediaUseCase) {
	handler := &socialMediaHandler{socialMediaUseCase}

	router := routers.Group("/socialmedias")
	{
		router.Use(middleware.Authentication())
		router.GET("", handler.Get)
		router.POST("", handler.Create)
		router.PUT("/:socialMediaId", middleware.Authorization(handler.socialMediaUseCase), handler.Edit)
		router.DELETE("/:socialMediaId", middleware.Authorization(handler.socialMediaUseCase), handler.Delete)
	}
}

// Get godoc
// @Summary    	Get all social media
// @Description	Get all social media with authentication user
// @Tags        socialmedias
// @Accept      json
// @Produce     json
// @Success     200	{object}	utils.ResponseDataGetedSocialMedia
// @Failure     400	{object}	utils.ResponseMessage
// @Failure     401	{object}	utils.ResponseMessage
// @Security    Bearer
// @Router      /socialmedias	[get]
func (handler *socialMediaHandler) Get(ctx *gin.Context) {
	var (
		socialMedias []domain.SocialMedia
		err          error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = handler.socialMediaUseCase.Get(ctx.Request.Context(), &socialMedias, userID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail cant find the id",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseData{
		Status: "congratulation you have successfully send your data, heres your data ",
		Data: utils.GetedSocialMedia{
			SocialMedias: socialMedias,
		},
	})
}

// Create godoc
// @Summary    	Add a social media
// @Description	Create and create a social media with authentication user
// @Tags        socialmedias
// @Accept      json
// @Produce     json
// @Param       json	body			utils.AddSocialMedia true  "Add Social Media"
// @Success     201		{object}  utils.ResponseDataAddedSocialMedia
// @Failure     400		{object}	utils.ResponseMessage
// @Failure     401		{object}	utils.ResponseMessage
// @Security    Bearer
// @Router      /socialmedias		[post]
func (handler *socialMediaHandler) Create(ctx *gin.Context) {
	var (
		socialMedia domain.SocialMedia
		err         error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "error to find the id",
			Message: "failed to find the id",
		})

		return
	}

	socialMedia.UserID = userID

	if err = handler.socialMediaUseCase.Create(ctx.Request.Context(), &socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, helpers.ResponseData{
		Status: "success",
		Data: utils.AddedSocialMedia{
			ID:             socialMedia.ID,
			UserID:         socialMedia.UserID,
			Name:           socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
			CreatedAt:      socialMedia.CreatedAt,
		},
	})
}

// Edit godoc
// @Summary     Edit a social media
// @Description	Edit a social media by id with authentication user
// @Tags        socialmedias
// @Accept      json
// @Produce     json
// @Param       id		path      string	true	"SocialMedia ID"
// @Param				json	body			utils.EditSocialMedia	true	"Edit Social Media"
// @Success     200		{object}	utils.ResponseDataEditedSocialMedia
// @Failure     400		{object}	utils.ResponseMessage
// @Failure     401		{object}	utils.ResponseMessage
// @Failure     404		{object}	utils.ResponseMessage
// @Security    Bearer
// @Router      /socialmedias/{id} [put]
func (handler *socialMediaHandler) Edit(ctx *gin.Context) {
	var (
		socialMedia domain.SocialMedia
		err         error
	)

	socialMediaID := ctx.Param("socialMediaId")
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	editedSocialMedia := domain.SocialMedia{
		UserID:         userID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
	}

	if socialMedia, err = handler.socialMediaUseCase.Edit(ctx.Request.Context(), editedSocialMedia, socialMediaID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, utils.EditedSocialMedia{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
		UpdatedAt:      socialMedia.UpdatedAt,
	})
}

// Delete godoc
// @Summary     Delete a social media
// @Description	Delete a social media by id with authentication user
// @Tags        socialmedias
// @Accept      json
// @Produce     json
// @Param       id   path     string  true  "SocialMedia ID"
// @Success     200  {object}	utils.ResponseMessageDeletedSocialMedia
// @Failure     400  {object}	utils.ResponseMessage
// @Failure     401  {object}	utils.ResponseMessage
// @Failure     404  {object}	utils.ResponseMessage
// @Security    Bearer
// @Router      /socialmedias/{id} [delete]
func (handler *socialMediaHandler) Delete(ctx *gin.Context) {
	socialMediaID := ctx.Param("socialMediaId")

	if err := handler.socialMediaUseCase.Delete(ctx.Request.Context(), socialMediaID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "your social media has been successfully deleted",
	})
}
