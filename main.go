package main

import (
	"log"
	commentDelivery "mygram-byferdiansyah/comment/delivery/http"
	commentRepository "mygram-byferdiansyah/comment/repository/postgres"
	commentUseCase "mygram-byferdiansyah/comment/usecase"
	"mygram-byferdiansyah/config/database"
	imageDelivery "mygram-byferdiansyah/image/delivery/http"
	imageRepository "mygram-byferdiansyah/image/repository/postgres"
	imageUseCase "mygram-byferdiansyah/image/usecase"
	socialMediaDelivery "mygram-byferdiansyah/socialmedia/delivery/http"
	socialMediaRepository "mygram-byferdiansyah/socialmedia/repository/postgres"
	socialMediaUseCase "mygram-byferdiansyah/socialmedia/usecase"
	userDelivery "mygram-byferdiansyah/user/delivery/http"
	userRepository "mygram-byferdiansyah/user/repository/postgres"
	userUseCase "mygram-byferdiansyah/user/usecase"
	"os"

	_ "mygram-byferdiansyah/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyGram By Ferdiansya
// @version 1.0
// @description This API was made as a primary purpose for one of the requirements (final project) in the Hack8tiv and FGA Kominfo courses. MyGram is a website similar to Instagram. On this website, users can register by login (if they are over eight years old), post images, and comments.
// @termOfService http://swagger.io/terms/
// @contact.name ferdi
// @contact.email ferdicompany@gmail.com
// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description					        Description for what is this security definition being used
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	db := database.StartDB()

	routers := gin.Default()

	routers.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	})

	userRepository := userRepository.NewUserRepository(db)
	userUseCase := userUseCase.NewUserUseCase(userRepository)

	userDelivery.NewUserHandler(routers, userUseCase)

	imageRepository := imageRepository.NewImageRepository(db)
	imageUseCase := imageUseCase.NewImageUseCase(imageRepository)

	imageDelivery.NewImageHandler(routers, imageUseCase)

	commentRepository := commentRepository.NewCommentRepository(db)
	commentUseCase := commentUseCase.NewCommentUseCase(commentRepository)

	commentDelivery.NewCommentHandler(routers, commentUseCase, imageUseCase)

	socialMediaRepository := socialMediaRepository.NewSocialMediaRepository(db)
	socialMediaUseCase := socialMediaUseCase.NewSocialMediaUseCase(socialMediaRepository)

	socialMediaDelivery.NewSocialMediaHandler(routers, socialMediaUseCase)

	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	port := os.Getenv("PORT")

	if len(os.Args) > 1 {
		reqPort := os.Args[1]

		if reqPort != "" {
			port = reqPort
		}
	}

	if port == "" {
		port = "8080"
	}

	routers.Run(":" + port)
}
