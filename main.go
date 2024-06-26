package main

import (
	"fmt"
	"remy-aquavelo/config"
	"strings"
	"time"

	"github.com/kataras/golog"

	"github.com/kataras/iris/v12"
)

func debugLogger() iris.Handler {
	return func(c iris.Context) {
		t := time.Now()
		c.Next()
		params := []string{
			fmt.Sprint(c.GetStatusCode()),
			c.Request().Method,
			c.RequestPath(false),
			time.Since(t).String(),
		}
		golog.Debug(strings.Join(params, " | "))
	}
}

func main() {
	golog.Info("Starting server")

	router := iris.New()

	router.Logger().SetLevel(config.Cfg.App.DebugLevel)
	if config.Cfg.App.DebugLevel == "debug" {
		router.Use(debugLogger())
	}

	router.Use(iris.Compression)

	router.Get("/", func(c iris.Context) {
		c.JSON(struct{ Message string }{Message: "Welcome to the oui API (ahah t'a caté la ref)"})
	})

	err := router.Listen(":" + config.Cfg.App.Port)
	if err != nil {
		golog.Fatal(err)
	}
}
