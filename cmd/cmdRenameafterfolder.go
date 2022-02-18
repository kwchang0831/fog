package cmd

import (
	"github.com/kwchang0831/fog/utli"
	"github.com/spf13/cobra"
)

var cmdRenameafterfolder = &cobra.Command{
	Use:   "renameafterfolder find",
	Short: "Rename files inside matching folders with the folder name.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utli.RenameAfterFolder(dir, args[0], exclude, wet)
	},
}

func init() {
	rootCmd.AddCommand(cmdRenameafterfolder)
	cmdRenameafterfolder.PersistentFlags().BoolVarP(&wet, "wet", "w", false, "wet run, commit changes")
	cmdRenameafterfolder.PersistentFlags().StringVarP(&dir, "dir", "d", ".", "search target dir")
	cmdRenameafterfolder.PersistentFlags().StringVarP(&exclude, "exclude", "e", "", "exclude the pattern")
}
