package servicemanager

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ServiceManager struct {
	engine *gin.Engine
}

func Initialize() *ServiceManager {
	// TODO: support to custom host and port
	serviceManager := ServiceManager{
		engine: gin.Default(),
	}

	return &serviceManager
}

func (sm *ServiceManager) SetGroupRoute(groupPath string, middleware gin.HandlerFunc, routeInfo RouterInfo) error {
	routerGroup := sm.engine.Group(groupPath)
	routerGroup.Use(middleware)
	var err error
	if routeInfo.Method == "GET" {
		routerGroup.GET(routeInfo.Path, routeInfo.HandlerFunc)
	} else if routeInfo.Method == "POST" {
		routerGroup.POST(routeInfo.Path, routeInfo.HandlerFunc)
	} else {
		err = errors.New(fmt.Sprintf("Unsupported Method %s", routeInfo.Method))
	}
	return err
}

func (sm *ServiceManager) SetRoute(routeInfo RouterInfo) error {
	var err error
	if routeInfo.Method == "GET" {
		sm.engine.GET(routeInfo.Path, routeInfo.HandlerFunc)
	} else if routeInfo.Method == "POST" {
		sm.engine.POST(routeInfo.Path, routeInfo.HandlerFunc)
	} else {
		err = errors.New(fmt.Sprintf("Unsupported Method %s", routeInfo.Method))
	}
	return err
}

func (sm *ServiceManager) StartService() {
	// TODO: add validation function before run
	sm.engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
