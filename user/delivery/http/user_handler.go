package delivery

import (
	"mygram-byferdiansyah/domain"
	"mygram-byferdiansyah/helpers"
	"mygram-byferdiansyah/user/delivery/http/middleware"
	"mygram-byferdiansyah/user/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUseCase domain.UserUseCase
}

func NewUserHandler(routers *gin.Engine, userUseCase domain.UserUseCase) {
	handler := &userHandler{userUseCase}

	router := routers.Group("/users")
	{
		router.POST("/register", handler.Register)
		router.POST("/login", handler.Login)
		router.PUT("", middleware.Authentication(), handler.Edit)
		router.DELETE("", middleware.Authentication(), handler.Delete)
	}
}

// Register godoc
// @Summary			Register a user
// @Description	create and create a user
// @Tags				users
// @Accept			json
// @Produce			json
// @Param				json	body			utils.RegisterUser	true	"Register User"
// @Success			201		{object}	utils.ResponseDataRegisteredUser
// @Failure			400  	{object}	utils.ResponseMessage
// @Failure			409  	{object}	utils.ResponseMessage
// @Router			/users/register	[post]
func (handler *userHandler) Register(ctx *gin.Context) {
	var (
		user domain.User
		err  error
	)

	if err = ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	if err = handler.userUseCase.Register(ctx.Request.Context(), &user); err != nil {
		if strings.Contains(err.Error(), "idx_users_username") {
			ctx.AbortWithStatusJSON(http.StatusConflict, helpers.ResponseMessage{
				Status:  "fail",
				Message: "the username you entered has been used",
			})

			return
		}

		if strings.Contains(err.Error(), "idx_users_email") {
			ctx.AbortWithStatusJSON(http.StatusConflict, helpers.ResponseMessage{
				Status:  "fail",
				Message: "the email you entered has been used",
			})

			return
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, helpers.ResponseData{
		Status: "success",
		Data: utils.RegisteredUser{
			Age:      user.Age,
			Email:    user.Email,
			ID:       user.ID,
			Username: user.Username,
		},
	})
}

// Login godoc
// @Summary			Login a user
// @Description	Authentication a user and retrieve a token
// @Tags				users
// @Accept			json
// @Produce			json
// @Param				json	body			utils.LoginUser	true	"Login User"
// @Success			200		{object}	utils.ResponseDataLoggedinUser
// @Failure			400		{object}	utils.ResponseMessage
// @Failure			401		{object}	utils.ResponseMessage
// @Router			/users/login		[post]
func (handler *userHandler) Login(ctx *gin.Context) {
	var (
		user  domain.User
		err   error
		token string
	)

	if err = ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	if err = handler.userUseCase.Login(ctx.Request.Context(), &user); err != nil {
		if strings.Contains(err.Error(), "the credential you entered are wrong") {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
				Status:  "unauthenticated",
				Message: err.Error(),
			})

			return
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "unauthenticated",
			Message: err.Error(),
		})

		return
	}

	if token = helpers.GenerateToken(user.ID, user.Email); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "unauthenticated",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data: utils.LoggedinUser{
			Token: token,
		},
	})
}

// Edit godoc
// @Summary			Edit a user
// @Description	Edit a user with authentication user
// @Tags				users
// @Accept			json
// @Produce			json
// @Param				json		body			utils.EditUser   true  "Edit User"
// @Success			200			{object}  utils.ResponseDataEditedUser
// @Failure			400			{object}	utils.ResponseMessage
// @Failure			401			{object}	utils.ResponseMessage
// @Failure			409			{object}	utils.ResponseMessage
// @Security		Bearer
// @Router			/users	[put]
func (handler *userHandler) Edit(ctx *gin.Context) {
	var (
		user domain.User
		err  error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	_ = string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	editedUser := domain.User{
		Username: user.Username,
		Email:    user.Email,
	}

	if user, err = handler.userUseCase.Edit(ctx.Request.Context(), editedUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data: utils.EditedUser{
			ID:        user.ID,
			Email:     user.Email,
			Username:  user.Username,
			Age:       user.Age,
			UpdatedAt: user.UpdatedAt,
		},
	})
}

// Delete godoc
// @Summary			Delete a user
// @Description	Delete a user with authentication user
// @Tags				users
// @Accept			json
// @Produce			json
// @Success			200			{object}	utils.ResponseMessageDeletedUser
// @Failure			400			{object}	utils.ResponseMessage
// @Failure			401			{object}	utils.ResponseMessage
// @Failure			404			{object}	utils.ResponseMessage
// @Security		Bearer
// @Router			/users	[delete]
func (handler *userHandler) Delete(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err := handler.userUseCase.Delete(ctx, userID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: "account not found",
		})

		return
	}

	ctx.JSON(
		http.StatusOK,
		helpers.ResponseMessage{
			Status:  "success",
			Message: "your account has been successfully deleted",
		},
	)
}
