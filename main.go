package main

import (
	"time"
)

func main() {
	configurations := NewConfig(5*time.Second, 5*time.Minute)
	startRepl(configurations)
}
