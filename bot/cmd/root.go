package cmd

import (
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
}
