/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
  "github.com/alecthomas/kong"
  "github.com/terminalPoltergeist/lux/test"
)

// Lux is the command-line interface for Lux
type Lux struct {
  Test test.Options `cmd:"" help:"This is a test"`
}
