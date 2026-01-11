package generate

import (
	"github.com/jemiezler/gofiber-cli/optics/internal/generator"
	"github.com/spf13/cobra"
)

func init() {
	GenerateCmd.AddCommand(moduleCmd)
}

var moduleCmd = &cobra.Command{
	Use:   "module [name]",
	Short: "Generate module",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		generator.GenerateModule(args[0])
	},
}
