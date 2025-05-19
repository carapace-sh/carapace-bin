package main

import (
	_ "embed"
	"encoding/json"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

type Command struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Synopsis    string `json:"synopsis"`
	Definition  struct {
	} `json:"definition"`
	Aliases   []string `json:"aliases"`
	Arguments []any    `json:"arguments"`
	Options   []struct {
		Name          string `json:"name"`
		Description   string `json:"description"`
		ValueRequired bool   `json:"value_required"`
		ValueOptional bool   `json:"value_optional"`
	} `json:"options"`
}

//go:embed artisan.json
var artisan string // https://artisan.page/api/11.x

func main() {
	var commands []Command
	err := json.Unmarshal([]byte(artisan), &commands)
	if err != nil {
		panic(err.Error())
	}

	cmd := &cobra.Command{
		Use:   "artisan",
		Short: "Artisan is the command line interface included with Laravel",
		Long:  "https://laravel.com/",
	}
	for _, command := range commands {
		cmd.AddCommand(convert(command))
	}
	carapace.Gen(cmd)

	cmd.Execute()
}

func convert(c Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:     c.Name,
		Short:   c.Description,
		Aliases: c.Aliases,
	}

	if c.Synopsis != "" {
		cmd.Use = c.Synopsis
	}

	for _, option := range c.Options {
		switch {
		case option.ValueOptional || option.ValueRequired:
			cmd.Flags().String(option.Name, "", option.Description)
		default:
			cmd.Flags().Bool(option.Name, false, option.Description)
		}
	}
	return cmd
}
