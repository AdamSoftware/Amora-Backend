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
	"Amora/internal/events"
	"Amora/internal/events/eventService"
	"Amora/internal/events/eventHandler"
	"Amora/internal/events/eventRepository"
	"Amora/internal/attachedGifts"
	"Amora/internal/attachedGifts/attachedGiftsHandler"
	"Amora/internal/attachedGifts/attachedGiftsService"
	"Amora/internal/attachedGifts/attachedGiftsRepository"
	"Amora/internal/gifts"
	"Amora/internal/gifts/gifthandler"
	"Amora/internal/gifts/giftservice"
	"Amora/internal/gifts/giftrepository"
	"Amora/internal/restaurants"
	"Amora/internal/restaurants/restauranthandler"
	"Amora/internal/restaurants/restaurantservice"
	"Amora/internal/restaurants/restaurantrepository"
	"Amora/internal/venues"
	"Amora/internal/venues/venueservice"
	"Amora/internal/venues/venuehandler"
	"Amora/internal/venues/venuerepository"
	"Amora/internal/experiences"
	"Amora/internal/experiences/experiencehandler"
	"Amora/internal/experiences/experienceservice"
	"Amora/internal/experiences/experiencerepository"
	"Amora/internal/plans"
	"Amora/internal/plans/planhandler"
	"Amora/internal/plans/planrepository"
	"Amora/internal/plans/planservice"
	"Amora/internal/messages"
	"Amora/internal/messages/messagehandler"
	"Amora/internal/messages/messageservice"
	"Amora/internal/messages/messagerepository"
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

func initEventModule(db *gorm.DB) (routes.RouteModule, error) {
	coreRepo := CoreRepository.CoreRepositoryInit[*events.Events]{DB: db}
	coreService := CoreService.CoreServiceInit[*events.Events]{Repo: &coreRepo}
	coreHandler := CoreHandler.CoreHandlerInit[*events.Events]{Service: &coreService}

	repo, err := eventRepository.NewEventRepositoryInit(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize event repo: %w", err)
	}

	service, err := eventService.NewEventServiceInit(repo)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize event service: %w", err)
	}

	handler := eventHandler.NewEventHandlerInit(&coreHandler, service)
	return routes.NeweventsModule(handler), nil

}

func initAttachedGiftsModule(db *gorm.DB) (routes.RouteModule, error) {
	coreRepo := CoreRepository.CoreRepositoryInit[*attachedGifts.AttachedGifts]{DB: db}
	coreService := CoreService.CoreServiceInit[*attachedGifts.AttachedGifts]{Repo: &coreRepo}
	coreHandler := CoreHandler.CoreHandlerInit[*attachedGifts.AttachedGifts]{Service: &coreService}

	repo, err := attachedGiftsRepository.NewattachedgiftsRepositoryInit(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize attached gifts repo: %w", err)
	}

	service, err := attachedGiftsService.NewAttachedGiftsServiceInit(repo)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize attached gifts service: %w", err)
	}

	handler := attachedGiftsHandler.NewattachedGiftsHandlerInit(&coreHandler, service)
	return routes.NewattachedGiftsModule(handler), nil
}

func initGiftModule(db *gorm.DB) (routes.RouteModule, error) {
	coreRepo := CoreRepository.CoreRepositoryInit[*gifts.Gifts]{DB: db}
	coreService := CoreService.CoreServiceInit[*gifts.Gifts]{Repo: &coreRepo}
	coreHandler := CoreHandler.CoreHandlerInit[*gifts.Gifts]{Service: &coreService}

	repo, err := giftrepository.NewGiftsRepositoryInit(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gift repo: %w", err)
	}

	service, err := giftservice.NewGiftServiceInit(repo)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gift service: %w", err)
	}

	handler := gifthandler.NewGiftHandlerInit(&coreHandler, service)
	return routes.NewGiftModule(handler), nil
}


func initRestuarantModule(db *gorm.DB) (routes.RouteModule, error) {
	coreRepo := CoreRepository.CoreRepositoryInit[*restaurants.Restaurants]{DB: db}
	coreService := CoreService.CoreServiceInit[*restaurants.Restaurants]{Repo: &coreRepo}
	coreHandler := CoreHandler.CoreHandlerInit[*restaurants.Restaurants]{Service: &coreService}

	repo, err := restaurantrepository.NewRestaurantsRepositoryInit(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gift repo: %w", err)
	}

	service, err := restaurantservice.NewRestaurantServiceInit(repo)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gift service: %w", err)
	}

	handler := restauranthandler.NewRestaurantHandlerInit(&coreHandler, service) 
	return routes.NewRestaurantModule(handler), nil
}

func initVenueModule(db *gorm.DB) (routes.RouteModule, error) {
	coreRepo := CoreRepository.CoreRepositoryInit[*venues.Venues]{DB: db}
	coreService := CoreService.CoreServiceInit[*venues.Venues]{Repo: &coreRepo}
	coreHandler := CoreHandler.CoreHandlerInit[*venues.Venues]{Service: &coreService}

	repo, err := venuerepository.NewVenuesRepositoryInit(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gift repo: %w", err)
	}

	service, err := venueservice.NewVenueServiceInit(repo)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gift service: %w", err)
	}

	handler := venuehandler.NewVenueHandlerInit(&coreHandler, service) 
	return routes.NewVenueModule(handler), nil
}

func initExperienceModule(db *gorm.DB) (routes.RouteModule, error) {
	coreRepo := CoreRepository.CoreRepositoryInit[*experiences.Experiences]{DB: db}
	coreService := CoreService.CoreServiceInit[*experiences.Experiences]{Repo: &coreRepo}
	coreHandler := CoreHandler.CoreHandlerInit[*experiences.Experiences]{Service: &coreService}

	repo, err := experiencerepository.NewExperiencesRepositoryInit(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize experience repo: %w", err)
	}

	service, err := experienceservice.NewExperienceServiceInit(repo)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize experience service: %w", err)
	}

	handler := experiencehandler.NewExperienceHandlerInit(&coreHandler, service)
	return routes.NewExperienceModule(handler), nil
}

func initPlanModule(db *gorm.DB) (routes.RouteModule, error) {
	coreRepo := CoreRepository.CoreRepositoryInit[*plans.Plans]{DB: db}
	coreService := CoreService.CoreServiceInit[*plans.Plans]{Repo: &coreRepo}
	coreHandler := CoreHandler.CoreHandlerInit[*plans.Plans]{Service: &coreService}

	repo, err := planrepository.NewPlansRepositoryInit(db) 
	if err != nil {
		return nil, fmt.Errorf("failed to initialize plan repo: %w", err)
	}

	service, err := planservice.NewPlanServiceInit(repo)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize plan service: %w", err)
	}

	handler := planhandler.NewPlanHandlerInit(&coreHandler, service)
	return routes.NewPlanModule(handler), nil
}

func initMessageModule(db *gorm.DB) (routes.RouteModule, error) {
	coreRepo := CoreRepository.CoreRepositoryInit[*messages.Messages]{DB: db}
	coreService := CoreService.CoreServiceInit[*messages.Messages]{Repo: &coreRepo}
	coreHandler := CoreHandler.CoreHandlerInit[*messages.Messages]{Service: &coreService}

	repo, err := messagerepository.NewMessagesRepositoryInit(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize message repo: %w", err)
	}

	service, err := messageservice.NewMessageServiceInit(repo)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize message service: %w", err)
	}

	handler := messagehandler.NewMessageHandlerInit(&coreHandler, service)
	return routes.NewMessageModule(handler), nil
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
		initEventModule,
		initAttachedGiftsModule,
		initGiftModule,
		initRestuarantModule,
		initVenueModule,
		initExperienceModule,
		initPlanModule,
		initMessageModule,
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

