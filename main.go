package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type HelloWorld struct {
}

func (c *HelloWorld) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "HelloWorldPlugin",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 12,
			Build: 4,
		},
		Commands: []plugin.Command{
			{
				Name:     "hello-world",
				HelpText: "Command to print out \"Hello, World!\"",
				UsageDetails: plugin.Usage{
					Usage: "cf hello-world",
				},
			},
		},
	}
}

func (c *HelloWorld) Run(cliConnection plugin.CliConnection, args []string) {
	fmt.Println("Hello, World!")
}

func main() {
	plugin.Start(new(HelloWorld))
}
