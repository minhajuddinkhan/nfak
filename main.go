package main

import (
	"fmt"
	"os"

	"github.com/nfak/cmd"
	"github.com/spf13/cobra"
)

func main() {

	var RootCmd = &cobra.Command{
		Use:   "hugo",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
					  love by spf13 and friends in Go.
					  Complete documentation is available at http://hugo.spf13.com`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			if len(args) == 0 {
				fmt.Println("no args")
				os.Exit(1)
			}
			if args[0] == "add" {
				yoloCommand := yolo.GoGetIt()
				yoloCommand.Execute()
			}

		},
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
