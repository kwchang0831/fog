package cmd

import (
	"github.com/kwchang0831/fog/utli"
	"github.com/spf13/cobra"
)

var cmdFolderin = &cobra.Command{
	Use:   "folderin",
	Short: "Move files into their own folders.",
	Long:  "Create folders and Move files in the target directory.",
	Run: func(cmd *cobra.Command, args []string) {
		utli.Folderin(dir, to, wet)
	},
}

func init() {
	rootCmd.AddCommand(cmdFolderin)
	cmdFolderin.PersistentFlags().BoolVarP(&wet, "wet", "w", false, "wet run, commit changes")
	cmdFolderin.PersistentFlags().StringVarP(&dir, "dir", "d", ".", "search target dir")
	cmdFolderin.PersistentFlags().StringVarP(&to, "to", "t", "", "move to target dir")
}
