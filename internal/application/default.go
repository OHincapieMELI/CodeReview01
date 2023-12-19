package application

import (
	"app/internal/handler"
	"app/internal/loader"
	"app/internal/repository"
	"app/internal/service"

	"github.com/gin-gonic/gin"
)

// ConfigDefaultInMemory is an struct that contains the configuration for the default application settings.
type ConfigDefaultInMemory struct {
	// FileLoader is the path to the file that contains the vehicles.
	FileLoader string
	// Addr is the address where the application will be listening.
	Addr string
}

// NewDefaultInMemory returns a new instance of a default application.
func NewDefaultInMemory(c *ConfigDefaultInMemory) *DefaultInMemory {
	// default config
	defaultCfg := &ConfigDefaultInMemory{
		FileLoader: "vehicles.json",
		Addr:       ":8080",
	}
	if c != nil {
		if c.FileLoader != "" {
			defaultCfg.FileLoader = c.FileLoader
		}
		if c.Addr != "" {
			defaultCfg.Addr = c.Addr
		}
	}

	return &DefaultInMemory{
		fileLoader: defaultCfg.FileLoader,
		addr:       defaultCfg.Addr,
	}
}

// DefaultInMemory is an struct that contains the default application settings.
type DefaultInMemory struct {
	// fileLoader is the path to the file that contains the vehicles.
	fileLoader string
	// addr is the address where the application will be listening.
	addr string
}

// Run starts the application.
func (d *DefaultInMemory) Run() (err error) {
	// dependencies initialization
	// loader
	ld := loader.NewVehicleJSON(d.fileLoader)
	data, err := ld.Load()
	if err != nil {
		return
	}

	// repository
	rp := repository.NewVehicleSlice(data.Data, data.LastId)

	// service
	sv := service.NewDefault(rp)

	// handler
	hd := handler.NewVehicleDefault(sv)

	// router
	rt := gin.New()
	// - middlewares
	rt.Use(gin.Logger())
	rt.Use(gin.Recovery())
	// - endpoints
	gr := rt.Group("/vehicles")
	{
		gr.GET("", hd.GetAll())
		gr.POST("", hd.Save())
		gr.PUT("/:id", hd.UpdateMaxSpeed())
	}

	// run application
	err = rt.Run(d.addr)
	if err != nil {
		return
	}

	return
}
