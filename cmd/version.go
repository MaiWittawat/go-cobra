package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show the version",
	Aliases: []string{"v", "V"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-cobra version 0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
