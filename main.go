/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
  "os"
  "fmt"
  "errors"
  "github.com/alecthomas/kong"
)

func main() {
  lux := &Lux{}
  ctx := kong.Parse(
    lux,
    kong.Description(fmt.printf("A LUXurious tmux tool.")),
    kong.UsageOnError(),
    kong.ConfigureHelp(kong.HelpOptions{
      Compact: true,
      Summary: false,
    }),
  )
  if err := ctx.Run(); err != nil {
    if errors.Is(err, exit.ErrAborted) {
      os.Exit(exit.StatusAborted)
    }
    fmt.Println(err)
    os.Exit(1)
  }
}
