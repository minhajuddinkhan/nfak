package yolo

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GoGetIt() *cobra.Command {

	return &cobra.Command{
		Use:   "add",
		Short: "Add a transaction",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here

			fmt.Println("Adding!")
		},
	}
}
