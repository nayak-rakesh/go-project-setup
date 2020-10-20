package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nayak-rakesh/go-project-setup/config"
	"github.com/nayak-rakesh/go-project-setup/database/postgres"
	_userHandler "github.com/nayak-rakesh/go-project-setup/internal/user/handler/http"
	_userRepo "github.com/nayak-rakesh/go-project-setup/internal/user/repository/postgres"
	_userService "github.com/nayak-rakesh/go-project-setup/internal/user/service"
)

func routes() *gin.Engine {
	// Initialize database
	db, err := postgres.NewDB(config.GetDBConnString())
	if err != nil {
		fmt.Println(err)
		panic("couldn't connect to database")
	}

	router := gin.Default()

	// user domain
	userRepo := _userRepo.Init(db)
	userService := _userService.NewUserService(userRepo)
	_userHandler.NewUserHandler(router, userService)

	return router
}

// StartApp ...
func StartApp() {
	router := routes()
	router.Run(":8080")
	fmt.Println("server stated on the port 8080")

}
