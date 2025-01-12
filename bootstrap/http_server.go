package bootstrap

import (
	"PawInHand/repositories"
	"PawInHand/resources"
	"PawInHand/services"
	"PawInHand/utils"
	"context"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"log"
	"net/http"

	"PawInHand/config"
	api "PawInHand/generated/api/PawInHand/rest"
)

var FXModule_HTTPServer = fx.Module(
	"http-server",
	fx.Provide(
		createAPIRoutes,
		createHTTPClient,
		createHttpServer,
		createMuxRouter,
		createAdController,
		createPostController,
		createShelterController,
		createVolunteerWorkController,
		createEventController,
		createRatingController,
		createUserController,
		createConverterI,
	),
	fx.Invoke(
		registerServerStartHook,
	),
)

// CreateAPIRoutes initializes and merges all API routes for the application.
func createAPIRoutes(
	ads *api.AdsAPIController,
	posts *api.PostsAPIController,
	shelters *api.ShelterAPIController,
	volunteerWorks *api.VolunteerAPIController,
	ratings *api.RatingAPIController,
	events *api.EventsAPIController,
	users *api.UserAPIController,
) api.Routes {
	return utils.Merge(
		ads.Routes(),
		posts.Routes(),
		shelters.Routes(),
		volunteerWorks.Routes(),
		ratings.Routes(),
		events.Routes(),
		users.Routes(),
	)
}

func createConverterI() resources.ConverterI {
	return resources.NewConverter()
}

// CreateAdController initializes and returns the AdAPIController with its dependencies.
func createAdController(
	db repositories.AdRepo,
	i resources.ConverterI,
) *api.AdsAPIController {
	return api.NewAdsAPIController(services.NewAdAPIService(db, i))
}

// CreatePostController initializes and returns the PostAPIController with its dependencies.
func createPostController(
	db repositories.PostRepo,
	i resources.ConverterI,
) *api.PostsAPIController {
	return api.NewPostsAPIController(services.NewPostAPIService(db, i))
}

// CreateShelterController initializes and returns the ShelterAPIController with its dependencies.
func createShelterController(
	db repositories.ShelterRepo,
	i resources.ConverterI,
) *api.ShelterAPIController {
	return api.NewShelterAPIController(services.NewShelterAPIService(db, i))
}

// CreateVolunteerWorkController initializes and returns the VolunteerWorkAPIController with its dependencies.
func createVolunteerWorkController(
	db repositories.VolunteerworkRepo,
) *api.VolunteerAPIController {
	return api.NewVolunteerAPIController(services.NewVolunteerWorkAPIService(db))
}

// CreateRatingController initializes and returns the RatingAPIController with its dependencies.
func createRatingController(
	db repositories.RatingRepo,
) *api.RatingAPIController {
	return api.NewRatingAPIController(services.NewRatingAPIService(db))
}

// CreateEventController initializes and returns the EventAPIController with its dependencies.
func createEventController(
	db repositories.EventRepo,
) *api.EventsAPIController {
	return api.NewEventsAPIController(services.NewEventAPIService(db))
}

// CreateUserController initializes and returns the UserAPIController with its dependencies.
func createUserController(
	db repositories.UserRepo,
	i resources.ConverterI,
) *api.UserAPIController {
	return api.NewUserAPIController(services.NewUserAPIService(db, i))
}

func createHTTPClient() *http.Client {
	return &http.Client{}
}

func createMuxRouter(routes api.Routes) *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		log.Printf("Registering route: %s %s", route.Method, route.Pattern)
		router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
	}
	return router
}

func createHttpServer(config *config.ApplicationConfiguration, router *mux.Router) *http.Server {
	return &http.Server{
		Addr:    ":" + config.HTTPPort,
		Handler: router,
	}
}

func registerServerStartHook(lc fx.Lifecycle, server *http.Server) {
	lc.Append(fx.StartStopHook(
		func() {
			go func() {
				if err := server.ListenAndServe(); err != nil {
					log.Fatalf("failed to start server: %v", err)
				}
			}()
		},
		func() {
			if err := server.Shutdown(context.Background()); err != nil {
				log.Fatalf("failed to shutdown server: %v", err)
			}
		},
	))
}
