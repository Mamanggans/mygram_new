package delivery

import (
	"fmt"
	"mygram-byferdiansyah/comment/delivery/http/middleware"
	"mygram-byferdiansyah/comment/utils"
	"mygram-byferdiansyah/domain"
	"mygram-byferdiansyah/helpers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	commentUseCase domain.CommentUseCase
	imageUseCase   domain.ImageUseCase
}

func NewCommentHandler(routers *gin.Engine, commentUseCase domain.CommentUseCase, imageUseCase domain.ImageUseCase) {
	handler := &commentHandler{commentUseCase, imageUseCase}

	router := routers.Group("/comments")
	{
		router.Use(middleware.Authentication())
		router.GET("", handler.Get)
		router.POST("", handler.Create)
		router.PUT("/:commentId", middleware.Authorization(handler.commentUseCase), handler.Edit)
		router.DELETE("/:commentId", middleware.Authorization(handler.commentUseCase), handler.Delete)
	}
}

// Get godoc
// @Summary			Get all comments
// @Description	Get all comments with authentication user
// @Tags        comments
// @Accept      json
// @Produce     json
// @Success     200	{object}	utils.ResponseDataGetedComment
// @Failure     400	{object}	utils.ResponseMessage
// @Failure     401	{object}	utils.ResponseMessage
// @Security    Bearer
// @Router      /comments     [get]
func (handler *commentHandler) Get(ctx *gin.Context) {
	var (
		comments []domain.Comment

		err error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = handler.commentUseCase.Get(ctx.Request.Context(), &comments, userID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail please try again",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseData{
		Status: "congratulation its success",
		Data:   comments,
	})
}

// Create godoc
// @Summary			Add a comment
// @Description	create and create a comment with authentication user
// @Tags        comments
// @Accept      json
// @Produce     json
// @Param       json	body			utils.AddComment true  "Add Comment"
// @Success     201		{object}  utils.ResponseDataAddedComment
// @Failure     400		{object}	utils.ResponseMessage
// @Failure     401		{object}	utils.ResponseMessage
// @Security    Bearer
// @Router      /comments	[post]
func (handler *commentHandler) Create(ctx *gin.Context) {
	var (
		comment domain.Comment
		image   domain.Image
		err     error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&comment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "its fail please try again or contact the admin",
			Message: err.Error(),
		})

		return
	}

	imageID := comment.ImageID

	if err = handler.imageUseCase.GetByID(ctx.Request.Context(), &image, imageID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, helpers.ResponseMessage{
			Status:  "fail no image found",
			Message: fmt.Sprintf("image with id %s doesn't exist", imageID),
		})

		return
	}

	comment.UserID = userID

	if err = handler.commentUseCase.Create(ctx.Request.Context(), &comment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail doesnt found the comment",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, helpers.ResponseData{
		Status: "congratulation, these are your data",
		Data: utils.AddedComment{
			ID:        comment.ID,
			UserID:    comment.UserID,
			ImageID:   comment.ImageID,
			Message:   comment.Message,
			CreatedAt: comment.CreatedAt,
		},
	})
}

// Edit godoc
// @Summary			Edit a comment
// @Description	Edit a comment by id with authentication user
// @Tags        comments
// @Accept      json
// @Produce     json
// @Param       id		path			string  true  "Comment ID"
// @Param       json	body			utils.EditComment	true	"Edit Comment"
// @Success     200		{object}  utils.ResponseDataEditedComment
// @Failure     400		{object}	utils.ResponseMessage
// @Failure     401		{object}	utils.ResponseMessage
// @Failure     404		{object}	utils.ResponseMessage
// @Security    Bearer
// @Router      /comments/{id}	[put]
func (handler *commentHandler) Edit(ctx *gin.Context) {
	var (
		comment domain.Comment
		image   domain.Image
		err     error
	)

	commentID := ctx.Param("commentId")
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&comment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail to edit your data ",
			Message: err.Error(),
		})

		return
	}

	editedComment := domain.Comment{
		UserID:  userID,
		Message: comment.Message,
	}

	if image, err = handler.commentUseCase.Edit(ctx.Request.Context(), editedComment, commentID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail to edit your data",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success here your comment that you edited",
		Data: utils.EditedComment{
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
// @Summary			Delete a comment
// @Description	Delete a comment by id with authentication user
// @Tags        comments
// @Accept      json
// @Produce     json
// @Param       id  path			string	true	"Comment ID"
// @Success     200 {object}	utils.ResponseMessageDeletedComment
// @Failure     400 {object}	utils.ResponseMessage
// @Failure     401	{object}	utils.ResponseMessage
// @Failure     404	{object}	utils.ResponseMessage
// @Security    Bearer
// @Router      /comments/{id}	[delete]
func (handler *commentHandler) Delete(ctx *gin.Context) {
	commentID := ctx.Param("commentId")

	if err := handler.commentUseCase.Delete(ctx.Request.Context(), commentID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail cannot find your id",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "your comment has been successfully deleted",
	})
}
