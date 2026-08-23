package cmd

import (
	"github.com/spf13/cobra"

	"xgit/internal/repo"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes an empty xgit repo",
	Long:  "This command initializes an empty xgit repo",

	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}
		return repo.InitRepo(path)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
