package llm

import (
	"context"
	_ "embed"
	"strings"

	"maragu.dev/goo/llm"

	"app/model"
)

//go:embed prompts/system.txt
var system string

func (c *Client) Translate(ctx context.Context, to model.Language, text string) (string, error) {
	var b strings.Builder
	m := llm.Message{
		Content: "Oversæt følgende tekst til " + to.String() + ":\n\n" + text,
		Role:    llm.MessageRoleUser,
	}
	if err := c.p.Prompt(ctx, system, []llm.Message{m}, &b); err != nil {
		return "", err
	}
	return b.String(), nil
}
