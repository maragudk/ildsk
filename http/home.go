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

			var to model.Language
			var in, name string

			target := hx.GetTarget(r.Header)
			switch target {
			case "ildsk":
				to = model.LanguageIldsk
				name = "Ildsk"
				in = req.Dansk

			case "dansk":
				to = model.LanguageDanish
				name = "Dansk"
				in = req.Ildsk

			default:
				log.Event("Unknown target", 1, "target", target)
				return nil, errors.New("unknown target " + target)
			}

			log.Event("Translating", 1, "to", to, "in", in)
			out, err := llm.Translate(r.Context(), to, in)
			if err != nil {
				return html.ErrorPage(html.Page), err
			}
			log.Event("Translated", 1, "to", to, "in", in, "out", out)

			return html.TextareaPartial(name, out), nil
		})(w, r)
	}))
}
