package cmd

import (
	"github.com/spf13/cobra"
)

var generatorCmd = &cobra.Command{
	Use:   "generator",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
func init() {
	rootCmd.AddCommand(generatorCmd)

}
