package cmd

import (
	"xgit/internal/repo"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds all the files to the xgit",
	Long:  "adds all untracked files from the working directory to xgit",

	Args: cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}
		return repo.AddRepo(path)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
