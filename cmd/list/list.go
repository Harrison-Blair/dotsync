package list

import (
	"fmt"

	"github.com/Harrison-Blair/dot/internal/list"
	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List config entries eligable to be synced",
		Long: `List config entries eligable to be synced
Entries are filtered out based on the sync repo's .gitignore file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			verbose, err := cmd.Flags().GetBool("verbose")
			if err != nil {
				return err
			}

			entries, err := list.List(verbose)
			for _, entry := range entries {
				fmt.Printf("%q\n", entry.Name())
			}

			return nil
		},
	}

	// cmd.addCommand()

	return cmd
}
