package cmd

import (
	"github.com/kwchang0831/fog/utli"
	"github.com/spf13/cobra"
)

var cmdRevert = &cobra.Command{
	Use:   "revert pathToLogFile",
	Short: "Revert the committed changes from log file.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utli.Revert(args[0], wet)
	},
}

func init() {
	rootCmd.AddCommand(cmdRevert)
	cmdRevert.PersistentFlags().BoolVarP(&wet, "wet", "w", false, "wet run, commit changes")
}
