package config

import (
	"github.com/inonsdn/gacha-system/api_gateway/internal/client"
	"github.com/inonsdn/gacha-system/api_gateway/servicemanager"
)

var routerConfig = []servicemanager.RouterInfo{
	{
		Method:      "GET",
		Path:        "/ping",
		HandlerFunc: servicemanager.Ping,
	},
}

func GetRouterConfig() []servicemanager.RouterInfo {
	return routerConfig
}

func GetUserRouter(client client.UserServiceClient) []servicemanager.RouterInfo {
	return []servicemanager.RouterInfo{
		{
			Method:      "POST",
			Path:        "/login",
			HandlerFunc: servicemanager.Login(client),
		},
		{
			Method:      "POST",
			Path:        "/register",
			HandlerFunc: servicemanager.Register(client),
		},
	}
}

func GetGachaRouter(client client.GachaServiceClient) []servicemanager.RouterInfo {
	return []servicemanager.RouterInfo{
		{
			Method:      "GET",
			Path:        "/gachaInfo/:categ",
			HandlerFunc: servicemanager.GetGachaInfo(client),
		},
		{
			Method:      "POST",
			Path:        "/draw",
			HandlerFunc: servicemanager.GachaDraw(client),
		},
	}
}
