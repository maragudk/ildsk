package llm

import (
	"maragu.dev/goo/llm"
)

type Client struct {
	p llm.Prompter
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) InjectLLMPrompter(p llm.Prompter) {
	c.p = p
}
