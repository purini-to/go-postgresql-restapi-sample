package server

import (
	"github.com/gin-gonic/gin"
)

// App is application context
type App struct {
	Engine *gin.Engine
}

// Run start application
func (a *App) Run() (err error) {
	return a.Engine.Run()
}

// ProvideApp provide App instance
func ProvideApp() *App {
	return &App{
		Engine: gin.New(),
	}
}
