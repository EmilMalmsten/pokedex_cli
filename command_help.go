package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("here is some help")
	return nil
}
