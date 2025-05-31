package servicemanager

import "github.com/gin-gonic/gin"

type RouterInfo struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

type GroupRouterInfo struct {
	Method      string
	Path        string
	HandlerFunc func(c interface{ GetServiceName() })
}
