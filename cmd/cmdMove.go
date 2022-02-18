package cmd

import (
	"github.com/kwchang0831/fog/utli"
	"github.com/spf13/cobra"
)

var cmdMove = &cobra.Command{
	Use:   "move find toDirectory",
	Short: "Move matching files/folders into destination directory.",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		utli.Movematches(dir, args[1], args[0], exclude, mode, wet)
	},
}

func init() {
	rootCmd.AddCommand(cmdMove)
	cmdMove.PersistentFlags().BoolVarP(&wet, "wet", "w", false, "wet run, commit changes")
	cmdMove.PersistentFlags().StringVarP(&dir, "dir", "d", ".", "search target dir")
	cmdMove.PersistentFlags().StringVarP(&exclude, "exclude", "e", "", "exclude the pattern")
	cmdMove.PersistentFlags().IntVarP(&mode, "mode", "m", 0, "0: File Only 1: Folder Only 2: Both")
}
