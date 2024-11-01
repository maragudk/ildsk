package main

import (
	"maragu.dev/goo/service"

	"app/html"
	"app/http"
	"app/llm"
)

func main() {
	log := service.NewLogger()

	c := llm.NewClient()

	service.Start(service.Options{
		HTTPRouterInjector:  http.InjectHTTPRouter(log, c),
		HTMLPage:            html.Page,
		LLMPrompterInjector: c.InjectLLMPrompter,
		Log:                 log,
		Migrate:             true,
	})
}
