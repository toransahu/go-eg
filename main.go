//
// handle_panics.go
// Copyright (C) 2020 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package main

import "fmt"

func main() {
	// try_handle_panic()
	try_enum()
}

//---------handle panic

func getSubString(s string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering")
			fmt.Println(r)
		}
	}()
	return s[0:10], nil
}

func try_handle_panic() {
	s, err := getSubString("Hello")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(s)
	fmt.Println("End")
}

// ----------- enum

type Providers int

const (
	AWS Providers = iota
	GCP
	AZURE
)

func (p Providers) String() string {
	return [...]string{"AWS", "GCP", "AZURE"}[p]
}

var Provider Providers

func try_enum() {
	Provider = AWS
	println(Provider.String())
}
