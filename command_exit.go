package main

import "os"

func commandExit(config *Config) error {
	os.Exit(1)
	return nil
}
