// Package HTML holds all the common HTML components and utilities.
package html

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
	"maragu.dev/goo/html"
)

var hashOnce sync.Once
var appCSSPath, appJSPath, htmxJSPath string

// Page layout with header, footer, and container to restrict width and set base padding.
func Page(props html.PageProps, children ...Node) Node {
	// Hash the paths for easy cache busting on changes
	hashOnce.Do(func() {
		appCSSPath = getHashedPath("public/styles/app.css")
		htmxJSPath = getHashedPath("public/scripts/htmx.js")
		appJSPath = getHashedPath("public/scripts/app.js")
	})

	return HTML5(HTML5Props{
		Title:       props.Title,
		Description: props.Description,
		Language:    "da",
		Head: []Node{
			Link(Rel("stylesheet"), Href(appCSSPath)),
			Script(Src(htmxJSPath), Defer()),
			Script(Src(appJSPath), Defer()),
		},
		Body: []Node{Class("bg-gray-900 text-white font-serif"),
			Div(Class("min-h-screen flex flex-col justify-between bg-white"),
				header(),
				Div(Class("grow"),
					container(true, true,
						Group(children),
					),
				),
				footer(),
			),
		},
	})
}

// header bar with logo and navigation.
func header() Node {
	return Div(Class("bg-gray-800 text-white shadow"),
		container(true, false,
			Div(Class("h-16 flex items-center justify-between"),
				A(Href("/"), Class("inline-flex items-center text-xl font-semibold hover:text-amber-500"),
					Img(Src("/images/logo.jpg"), Alt("Logo"), Class("h-12 w-auto bg-white rounded-full mr-4")),
					Text("Forside"),
				),
			),
		),
	)
}

// container restricts the width and sets padding.
func container(padX, padY bool, children ...Node) Node {
	return Div(
		Classes{
			"max-w-7xl mx-auto":     true,
			"px-4 md:px-8 lg:px-16": padX,
			"py-4 md:py-8":          padY,
		},
		Group(children),
	)
}

// footer with a link to the gomponents website.
func footer() Node {
	return Div(Class("bg-gray-800 text-white"),
		container(true, false,
			Div(Class("h-16 flex items-center justify-center space-x-8"),
				A(Class("hover:text-amber-500"), Href("https://evaehler.dk"), Text("Sprogdesign af Eva Ehler")),
				A(Class("hover:text-amber-500"), Href("https://www.maragu.dk"), Text("Software af maragu")),
			),
		),
	)
}

func getHashedPath(path string) string {
	externalPath := strings.TrimPrefix(path, "public")
	ext := filepath.Ext(path)
	if ext == "" {
		panic("no extension found")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Sprintf("%v.x%v", strings.TrimSuffix(externalPath, ext), ext)
	}

	return fmt.Sprintf("%v.%x%v", strings.TrimSuffix(externalPath, ext), sha256.Sum256(data), ext)
}
