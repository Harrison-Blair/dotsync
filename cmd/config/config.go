package config

import (
	"github.com/Harrison-Blair/dot/internal/config"
	"github.com/spf13/cobra"
)

func ConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Show or set the config of dot",
		Long:  `Show or set the configuration of dot`,
		RunE: func(cmd *cobra.Command, args []string) error {
			verbose, err := cmd.Flags().GetBool("verbose")
			if err != nil {
				return err
			}

			err = config.Get(verbose)
			return err
		},
	}

	return cmd
}
