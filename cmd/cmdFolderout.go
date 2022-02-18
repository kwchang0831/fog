package cmd

import (
	"github.com/kwchang0831/fog/utli"
	"github.com/spf13/cobra"
)

var cmdFolderout = &cobra.Command{
	Use:   "folderout",
	Short: "Move files out of folders",
	Long:  "Move all files in the target directory's folders to target directory and remove empty out folders.",
	Run: func(cmd *cobra.Command, args []string) {
		utli.Folderout(dir, to, wet)
	},
}

func init() {
	rootCmd.AddCommand(cmdFolderout)
	cmdFolderout.PersistentFlags().BoolVarP(&wet, "wet", "w", false, "wet run, commit changes")
	cmdFolderout.PersistentFlags().StringVarP(&dir, "dir", "d", ".", "search target dir")
	cmdFolderout.PersistentFlags().StringVarP(&to, "to", "t", "", "move to target dir")
}
