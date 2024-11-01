package http

import (
	"maragu.dev/goo/http"
	"maragu.dev/snorkel"

	"app/llm"
)

func InjectHTTPRouter(log *snorkel.Logger, c *llm.Client) func(*http.Router) {
	return func(r *http.Router) {
		// Group for HTML
		r.Group(func(r *http.Router) {
			Home(r)
			Translate(r, log, c)
		})
	}
}
