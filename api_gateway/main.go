package main

import (
	"fmt"

	"github.com/inonsdn/gacha-system/api_gateway/config"
	"github.com/inonsdn/gacha-system/api_gateway/internal/client"
	"github.com/inonsdn/gacha-system/api_gateway/internal/middleware"
	"github.com/inonsdn/gacha-system/api_gateway/servicemanager"
)

func initRoute(sm *servicemanager.ServiceManager) {

	// init public route
	routerConfigs := config.GetRouterConfig()
	fmt.Println("Initialize route for ", len(routerConfigs))
	for _, routeInfo := range routerConfigs {
		sm.SetRoute(routeInfo)
	}

	gachaClientService := client.NewGachaServiceClient()
	// init gacha route
	gachaRouteConfigs := config.GetGachaRouter(*gachaClientService)
	fmt.Println("Initialize gacha route for ", len(gachaRouteConfigs))
	for _, routeInfo := range gachaRouteConfigs {
		sm.SetGroupRoute("/", middleware.AuthJWT(), routeInfo)
	}
}

func main() {
	serviceManager := servicemanager.Initialize()
	initRoute(serviceManager)
	serviceManager.StartService()
}
