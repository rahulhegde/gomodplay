package main

import (
	"fmt"

	sampler "github.com/rahulhegde/gomodplay"
	finance "github.com/rahulhegde/gomodplay/finance"
	greeting "github.com/rahulhegde/gomodplay/greetings"
)

func main() {
	fmt.Println("finance: version: ", finance.GetFinance())
	fmt.Println("greeting: version: ", greeting.GetGreetings())
	fmt.Println("sampler: version: ", sampler.GetSampler())
}
