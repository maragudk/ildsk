# ildsk! 🔥

<img src="logo.jpg" alt="Logo" width="300" align="right">

Oversætter fra dansk til ildsk og tilbage.

Made with ✨sparkles✨ by [maragu](https://www.maragu.dev/).

## Getting started

You can start the app with:

```shell
make start
```

If you make style changes, watch the CSS with:

```shell
make watch-css
```

You can run tests and linting with:

```shell
make test lint
```

### Enabling TailwindCSS auto-complete in your IDE

[TailwindCSS has auto-complete of classnames (and more) through IDE plugins](https://tailwindcss.com/docs/editor-setup).

After you've installed the TailwindCSS plugin for your IDE, it needs some configuration to work with gomponents. Here's the config for VS Code and JetBrains IDEs:

<details>
<summary>VSCode</summary>

Edit `vscode-settings.json` and add the following:

```json
{
	"tailwindCSS.includeLanguages": {
		"go": "html",
	},
	"tailwindCSS.experimental.classRegex": [
		["Class(?:es)?[({]([^)}]*)[)}]", "[\"`]([^\"`]*)[\"`]"]
	],
}
```

[See the official plugin page for more info](https://marketplace.visualstudio.com/items?itemName=bradlc.vscode-tailwindcss)
</details>

<details>
<summary>JetBrains/GoLand</summary>

Go to `Settings` -> `Languages & Frameworks` -> `Style Sheets` -> `Tailwind CSS` and add the following (don't delete the other config):

```json
{
	"includeLanguages": {
		"go": "html"
	},
	"experimental": {
		"classRegex": [
			["Class(?:es)?[({]([^)}]*)[)}]", "[\"`]([^\"`]*)[\"`]"]
		]
	}
}
```

[See the official plugin page for more info](https://plugins.jetbrains.com/plugin/15321-tailwind-css)
</details>

## Deploying

The [CD workflow](.github/workflows/cd.yml) automatically builds a multi-platform Docker image and pushes it to the Github container registry GHCR.io, tagged with the commit hash as well as `latest`.

You can try building the image locally with:

```shell
make build-docker
```

Note that [you need the containerd image store enabled](https://docs.docker.com/desktop/containerd/#enable-the-containerd-image-store) for this to work.
