package cmd

import (
	"os"

	"github.com/Harrison-Blair/dot/cmd/config"
	"github.com/Harrison-Blair/dot/cmd/list"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dot",
	Short: "A CLI to sync your dotfiles",
	Long:  `A CLI to sync your dotfiles with git to the configured remote`,

	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(
		list.ListCmd(),
		config.ConfigCmd(),
	)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dotsync.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "show verbose output")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
