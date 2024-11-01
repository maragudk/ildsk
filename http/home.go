package http

import (
	"errors"
	"net/http"

	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx/http"
	ghttp "maragu.dev/gomponents/http"
	. "maragu.dev/goo/http"
	"maragu.dev/httph"
	"maragu.dev/snorkel"

	"app/html"
)

// Home handler for the home page.
func Home(r *Router) {
	r.Get("/", func(props html.PageProps) (Node, error) {
		return html.HomePage(props), nil
	})
}

type translator interface {
}

type TranslateRequest struct {
}

func Translate(r *Router, log *snorkel.Logger, llm translator) {
	r.Mux.Post("/translate", httph.FormHandler(func(w http.ResponseWriter, r *http.Request, req TranslateRequest) {
		ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
			if !hx.IsRequest(r.Header) {
				return nil, errors.New("not an htmx request")
			}

			switch hx.GetTarget(r.Header) {
			case "ildsk":
				return html.TextareaPartial("Ildsk", "BRAWAWWAW"), nil
			case "dansk":
				return html.TextareaPartial("Dansk", "Hej med dig"), nil
			default:
				log.Event("Unknown target", 1, "target", hx.GetTarget(r.Header))
				return nil, errors.New("unknown target")
			}
		})(w, r)
	}))
}
