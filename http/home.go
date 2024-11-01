package http

import (
	"context"
	"errors"
	"net/http"

	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx/http"
	ghttp "maragu.dev/gomponents/http"
	. "maragu.dev/goo/http"
	"maragu.dev/httph"
	"maragu.dev/snorkel"

	"app/html"
	"app/model"
)

// Home handler for the home page.
func Home(r *Router) {
	r.Get("/", func(props html.PageProps) (Node, error) {
		return html.HomePage(props), nil
	})
}

type translator interface {
	Translate(ctx context.Context, to model.Language, text string) (string, error)
}

type TranslateRequest struct {
	Dansk string
	Ildsk string
}

func Translate(r *Router, log *snorkel.Logger, llm translator) {
	r.Mux.Post("/translate", httph.FormHandler(func(w http.ResponseWriter, r *http.Request, req TranslateRequest) {
		ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
			if !hx.IsRequest(r.Header) {
				return nil, errors.New("not an htmx request")
			}

			switch hx.GetTarget(r.Header) {
			case "ildsk":
				log.Event("Translating", 1, "to", hx.GetTarget(r.Header), "text", req.Dansk)

				translated, err := llm.Translate(r.Context(), model.LanguageIldsk, req.Dansk)
				if err != nil {
					return html.ErrorPage(html.Page), err
				}

				return html.TextareaPartial("Ildsk", translated), nil

			case "dansk":
				log.Event("Translating", 1, "to", hx.GetTarget(r.Header), "text", req.Ildsk)

				translated, err := llm.Translate(r.Context(), model.LanguageDanish, req.Ildsk)
				if err != nil {
					return html.ErrorPage(html.Page), err
				}

				return html.TextareaPartial("Dansk", translated), nil
			default:
				log.Event("Unknown target", 1, "target", hx.GetTarget(r.Header))
				return nil, errors.New("unknown target")
			}
		})(w, r)
	}))
}
