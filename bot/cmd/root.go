package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "verity",
		Short: "Elite Dangerous companion bot for Discord",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if viper.GetBool("version") {
				versionCmd.Run(cmd, args)
				os.Exit(0)
			}

			if viper.GetString("token") == "" {
				fmt.Fprintln(os.Stderr, "Missing Discord Token. Must be specified with --token or the `TOKEN` environment variable.")
				os.Exit(1)
			}
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "print version")
	viper.BindPFlag("version", rootCmd.Flags().Lookup("version"))

	rootCmd.Flags().StringP("token", "T", "", "discord bot token")
	viper.BindPFlag("token", rootCmd.Flags().Lookup("token"))
	viper.BindEnv("token", "TOKEN")
}
