package lox

import "fmt"

type SyntaxError struct {
	line    int
	message string
}

func NewSyntaxError(line int, message string) *SyntaxError {
	return &SyntaxError{line, message}
}

func (e *SyntaxError) error() string {
	return fmt.Sprintf("[line %v] Error: %v", e.line, e.message)
}
