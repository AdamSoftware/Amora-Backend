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
	"github.com/go-chi/chi/v5"
)


func UserRoutes(db *gorm.DB) *chi.Mux{

	if db == nil {
		panic("Database connection is nil")
	}else {
		fmt.Println("Database connection is established")
	}

	fmt.Println("Initializing corerepo...")
	CoreRepo := CoreRepository.CoreRepositoryInit[*user.Users]{DB: db}
	fmt.Println("Init coreservice")
	CoreService := CoreService.CoreServiceInit[*user.Users]{Repo: &CoreRepo}
	fmt.Println("Init corehandler")
	coreHandler := CoreHandler.CoreHandlerInit[*user.Users]{Service: &CoreService}

	fmt.Println("Initializing user repository...")
	userRepo := userRepository.UserRepositoryInit{Core: &CoreRepo} 

	fmt.Println("Initializing user service...")
	userService, err := userService.NewUserServiceInit(&userRepo)

	if err != nil {
		Logging.Logs.Errorf("Failed to initialize user service: %v", err)
		http.Error(nil, fmt.Sprintf("Failed to initialize user service: %v", err), http.StatusInternalServerError)
		return nil
	}
	
	fmt.Println("Initializing user handler...")
	userHandler := userHandler.NewUserHandlerInit(&coreHandler, userService)

	fmt.Println("Creating user routes...")
	return routes.UserRouter(userHandler)
}


func RoutesInit(db *gorm.DB) {
	router := UserRoutes(db)

	if router == nil {
		Logging.Logs.Error("Failed to initialize user routes")
		return
	}

	fmt.Println("User routes initialized successfully")
	fmt.Println("starting the server on port 8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		Logging.Logs.Errorf("Failed to start server: %v", err)
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
}
