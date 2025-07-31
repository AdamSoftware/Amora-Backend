package Init

import (
	"Amora/internal/Logging"
	"Amora/internal/core/CoreHandler"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/core/CoreService"
	"Amora/internal/routes"
	"Amora/internal/user"
	"Amora/internal/user/userHandler"
	"Amora/internal/user/userService"
	"Amora/internal/user/userRepository"
	"fmt"
	"net/http"
	"gorm.io/gorm"
)


func UserRoutes(db *gorm.DB) {
	
	CoreRepo := CoreRepository.CoreRepositoryInit[*user.Users]{DB: db}
	CoreService := CoreService.CoreServiceInit[*user.Users]{Repo: &CoreRepo}
	coreHandler := CoreHandler.CoreHandlerInit[*user.Users]{Service: &CoreService}

	userRepo := userRepository.UserRepositoryInit{Core: &CoreRepo} 

	userService, err := userService.NewUserServiceInit(&userRepo)

	if err != nil {
		Logging.Logs.Errorf("Failed to initialize user service: %v", err)
		http.Error(nil, fmt.Sprintf("Failed to initialize user service: %v", err), http.StatusInternalServerError)
		return
	}
	
	userHandler := userHandler.NewUserHandlerInit(&coreHandler, userService)

	routes.UserRouter(userHandler)
}


func RoutesInit(db *gorm.DB) {
	UserRoutes(db)

	fmt.Println("Routes initialized successfully")
	fmt.Println("You can now start your server and access the routes")
	fmt.Println("starting server on port 8080...")
}
