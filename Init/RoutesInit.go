package Init

import (
	"Amora/internal/Logging"
	"Amora/internal/core/CoreHandler"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/core/CoreService"
	"Amora/internal/routes"
	"Amora/internal/subscriptions/subscriptionRepository"
	"Amora/internal/user"
	"Amora/internal/user/userHandler"
	"Amora/internal/user/userRepository"
	"Amora/internal/user/userService"
	"Amora/internal/subscriptions/subscriptionService"
	"Amora/internal/subscriptions/subscriptionHandler"
	"fmt"
	"net/http"
	"Amora/internal/subscriptions"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"Amora/internal/partner"
	"Amora/internal/partner/partnerHandler"
	"Amora/internal/partner/partnerService"
	"Amora/internal/partner/partnerRepository"
)


func initUserModule(db *gorm.DB) (routes.RouteModule, error) {
	coreRepo := CoreRepository.CoreRepositoryInit[*user.Users]{DB: db}
	coreService := CoreService.CoreServiceInit[*user.Users]{Repo: &coreRepo}
	coreHandler := CoreHandler.CoreHandlerInit[*user.Users]{Service: &coreService}

	userRepo := userRepository.UserRepositoryInit{Core: &coreRepo}
	userService, err := userService.NewUserServiceInit(&userRepo)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize user service: %w", err)
	}

	handler := userHandler.NewUserHandlerInit(&coreHandler, userService)
	return routes.NewUserModule(handler), nil
}

func initSubscriptionModule(db *gorm.DB) (routes.RouteModule, error) {
	coreRepo := CoreRepository.CoreRepositoryInit[*subscriptions.Subscriptions]{DB: db}
	coreService := CoreService.CoreServiceInit[*subscriptions.Subscriptions]{Repo: &coreRepo}
	coreHandler := CoreHandler.CoreHandlerInit[*subscriptions.Subscriptions]{Service: &coreService}

	repo, err := subscriptionRepository.NewSubscriptionRepositoryInit(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize subscription repo: %w", err)
	}

	service, err := subscriptionService.NewsubscriptionServiceInit(repo)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize subscription service: %w", err)
	}

	handler := subscriptionHandler.NewsubscriptionHandlerInit(&coreHandler, service)
	return routes.NewSubscriptionModule(handler), nil
}

func initPartnerModule(db *gorm.DB) (routes.RouteModule, error) {
	coreRepo := CoreRepository.CoreRepositoryInit[*partner.Partners]{DB: db}
	coreService := CoreService.CoreServiceInit[*partner.Partners]{Repo: &coreRepo}
	coreHandler := CoreHandler.CoreHandlerInit[*partner.Partners]{Service: &coreService}

	repo, err := partnerRepository.NewPartnerRepositoryInit(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize partner repo: %w", err)
	}

	service, err := partnerService.NewPartnerServiceInit(repo)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize partner service: %w", err)
	}

	handler := partnerHandler.NewpartnerHandlerInit(&coreHandler, service)
	return routes.NewpartnerModule(handler), nil

}



func RoutesInit(db *gorm.DB) {
	if db == nil {
		panic("❌ Database connection is nil")
	}
	fmt.Println("✅ Database connected")

	router := chi.NewRouter()

	modules := []func(*gorm.DB) (routes.RouteModule, error){
		initUserModule,
		initSubscriptionModule,
		initPartnerModule,
	}

	for _, initFn := range modules {
		module, err := initFn(db)
		if err != nil {
			Logging.Logs.Errorf("❌ Failed to initialize module: %v", err)
			continue
		}
		fmt.Printf("✅ Mounted: %s\n", module.Prefix())
		router.Mount(module.Prefix(), module.Router())
	}

	fmt.Println("🚀 Server starting on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		Logging.Logs.Errorf("❌ Server failed to start: %v", err)
	}
}

