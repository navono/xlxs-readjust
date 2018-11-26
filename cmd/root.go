package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xlxs-helper",
	Short: "xlxs 助手",
	Long:  `辅助 xlxs 完成一些重复性的工作`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello world :)")
	},
}

// Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
