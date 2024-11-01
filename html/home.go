package html

import (
	"fmt"
	"strings"

	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

// HomePage is the front page of the app.
func HomePage(props PageProps) Node {
	props.Title = "Ildsk! 🔥"

	return Page(props,
		Div(Class("prose prose-amber prose-lg md:prose-xl"),
			H1(Text("Ildsk! 🔥")),

			Form(Method("post"), Action("/translate"),
				Div(Class("flex flex-col sm:flex-row gap-16 align-center justify-center"),
					Div(Class("flex flex-col space-y-8"),
						Div(ID("dansk"),
							TextareaPartial("Dansk", ""),
						),

						button("Oversæt til ildsk 🔥", "#ildsk"),
					),

					Div(Class("flex flex-col space-y-8"),
						Div(ID("ildsk"),
							TextareaPartial("Ildsk", ""),
						),

						button("Oversæt til dansk 🇩🇰", "#dansk"),
					),
				),
			),
		),
	)
}

func TextareaPartial(name, value string) Node {
	placeholder := fmt.Sprintf("Skriv noget på %v her…", strings.ToLower(name))

	return Group{
		Div(
			Label(For(name+"-area"), Class("block text-sm/6 font-medium text-gray-900"), Text(name)),
			Div(Class("mt-2"),
				Textarea(Rows("10"), Cols("40"), Name(name), ID(name+"-area"), Placeholder(placeholder),
					Class("block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-amber-600 sm:text-sm/6 font-sans"),
					Text(value)),
			),
		),
	}
}

func button(text, target string) Node {
	return Button(Type("submit"),
		Class("rounded-md bg-amber-600 px-2.5 py-1.5 text-sm font-semibold text-white shadow-sm hover:bg-amber-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-amber-600 font-sans"),
		Text(text), hx.Post("/translate"), hx.Target(target))
}
