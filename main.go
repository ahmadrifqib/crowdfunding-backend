package main

import (
	"crowdfunding-backend/handler"
	"crowdfunding-backend/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Tes Simpan dari Service"
	// userInput.Email = "contoh@email.com"
	// userInput.Occupation = "orang"
	// userInput.Password = "password"

	// userService.RegisterUser(userInput)
	// user := user.User{
	// 	Name: "Rio",
	// }

	// userRepository.Save(user)

	userHandler := handler.NewUserHandler(userService)
	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	router.Run()
}
