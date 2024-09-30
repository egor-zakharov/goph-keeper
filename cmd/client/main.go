package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/egor-zakharov/goph-keeper/internal/cli"
	"github.com/egor-zakharov/goph-keeper/internal/client"
	"github.com/egor-zakharov/goph-keeper/internal/config"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)
	conf := config.NewConfig()
	conf.ParseFlag()
	c := client.New(conf.FlagRunGRPCAddr)
	err := c.Connect()
	if err != nil {
		fmt.Printf("Server is unavailable")
	}
	executor := cli.NewExecutor(c)
	client := prompt.New(
		executor.HandleCommands,
		executor.ShowPrompts,
		prompt.OptionShowCompletionAtStart(),
		prompt.OptionPrefix("> "),
		prompt.OptionMaxSuggestion(21),
		prompt.OptionInputTextColor(prompt.Green),
	)
	client.Run()
}
