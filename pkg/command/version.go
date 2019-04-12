package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of CMP",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("cmp v0.1 -- HEAD")
	},
}
