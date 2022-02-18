package cmd

import (
	"github.com/kwchang0831/fog/utli"
	"github.com/spf13/cobra"
)

var cmdRmEmptyDir = &cobra.Command{
	Use:   "rmemptydir",
	Short: "Remove empty folders.",
	Run: func(cmd *cobra.Command, args []string) {
		utli.RmEmptyDir(dir, wet)
	},
}

func init() {
	rootCmd.AddCommand(cmdRmEmptyDir)
	cmdRmEmptyDir.PersistentFlags().BoolVarP(&wet, "wet", "w", false, "wet run, commit changes")
	cmdRmEmptyDir.PersistentFlags().StringVarP(&dir, "dir", "d", ".", "search target dir")
}
