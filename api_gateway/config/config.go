package config

import (
	"local.dev/api_gateway/internal/client"
	"local.dev/api_gateway/servicemanager"
)

var routerConfig = []servicemanager.RouterInfo{
	{
		Method:      "GET",
		Path:        "/",
		HandlerFunc: servicemanager.Login,
	},
	{
		Method:      "GET",
		Path:        "/ping",
		HandlerFunc: servicemanager.Ping,
	},
	{
		Method:      "POST",
		Path:        "/login",
		HandlerFunc: servicemanager.Login,
	},
}

func GetRouterConfig() []servicemanager.RouterInfo {
	return routerConfig
}

func GetGachaRouter(client client.GachaServiceClient) []servicemanager.RouterInfo {
	return []servicemanager.RouterInfo{
		{
			Method:      "GET",
			Path:        "/gachaInfo/:categ",
			HandlerFunc: servicemanager.GetGachaInfo(client),
		},
		{
			Method:      "GET",
			Path:        "/draw",
			HandlerFunc: servicemanager.GachaDraw(client),
		},
	}
}
