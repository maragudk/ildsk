package html

import (
	. "maragu.dev/gomponents"
	goohtml "maragu.dev/goo/html"
)

type PageProps = goohtml.PageProps

type PageFunc = goohtml.PageFunc

func ErrorPage(page PageFunc) Node {
	return goohtml.ErrorPage(page)
}
