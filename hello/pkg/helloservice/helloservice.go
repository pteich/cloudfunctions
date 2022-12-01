package helloservice

import (
	"fmt"
	"html"
)

type HelloService struct {
}

func New() *HelloService {
	return &HelloService{}
}

func (h *HelloService) HelloName(name string) string {
	if name == "" {
		return "Hello, World!"
	}

	return fmt.Sprintf("Hello, %s!", html.EscapeString(name))
}
