package cmd

import (
	"xgit/internal/repo"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds all the files to the xgit",
	Long:  "adds all untracked files from the working directory to xgit",

	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		for _, a := range args {
			if err := repo.AddRepo(a); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
