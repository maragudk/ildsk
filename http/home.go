package http

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx/http"
	ghttp "maragu.dev/gomponents/http"

	"app/html"
)

type translator interface {
}

// Home handler for the home page.
func Home(r chi.Router) {
	r.Get("/", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return html.HomePage(html.PageProps{}), nil
	}))
}

func Translate(r chi.Router, log *slog.Logger, llm translator) {
	r.Post("/translate", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		if !hx.IsRequest(r.Header) {
			return nil, errors.New("not an htmx request")
		}

		switch hx.GetTarget(r.Header) {
		case "ildsk":
			return html.TextareaPartial("Ildsk", "BRAWAWWAW"), nil
		case "dansk":
			return html.TextareaPartial("Dansk", "Hej med dig"), nil
		default:
			log.Info("Unknown target", "target", hx.GetTarget(r.Header))
			return nil, errors.New("unknown target")
		}
	}))
}
