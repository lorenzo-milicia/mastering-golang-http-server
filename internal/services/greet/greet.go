package greet

import "fmt"

type Service struct{}

func (s Service) Greet(name string) string {
	return fmt.Sprintf("Hi, %s!", name)
}
