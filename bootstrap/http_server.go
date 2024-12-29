package bootstrap

import (
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
	),
	fx.Invoke(
		registerServerStartHook,
	),
)

func createAPIRoutes() api.Routes {
	//return utils.Merge()
	return nil
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
