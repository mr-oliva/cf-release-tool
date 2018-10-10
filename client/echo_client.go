package client

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
)

type PluginDemonstratingParams struct {
	Uppercase *bool
}

func (pluginDemo *PluginDemonstratingParams) Run(cliConnection plugin.CliConnection, args []string) {
	// Initialize flags
	echoFlagSet := flag.NewFlagSet("echo", flag.ExitOnError)
	uppercase := echoFlagSet.Bool("uppercase", false, "displayes all provided text in uppercase")

	// Parse starting from [1] because the [0]th element is the
	// name of the command
	err := echoFlagSet.Parse(args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var itemToEcho string
	for _, value := range echoFlagSet.Args() {
		if *uppercase {
			itemToEcho += strings.ToUpper(value) + " "
		} else {
			itemToEcho += value + " "
		}
	}

	fmt.Println(itemToEcho)
}

func (pluginDemo *PluginDemonstratingParams) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "EchoDemo",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 1,
			Build: 4,
		},
		Commands: []plugin.Command{
			{
				Name:     "echo",
				Alias:    "repeat",
				HelpText: "Echo text passed into the command. To obtain more information use --help",
				UsageDetails: plugin.Usage{
					Usage: "echo - print input arguments to screen\n   cf echo [-uppercase] text",
					Options: map[string]string{
						"uppercase": "If this param is passed, which ever word is passed to echo will be all capitals.",
					},
				},
			},
		},
	}
}