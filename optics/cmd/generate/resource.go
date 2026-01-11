package generate

import (
	"github.com/jemiezler/gofiber-cli/optics/internal/generator"
	"github.com/spf13/cobra"
)

var resourceCmd = &cobra.Command{
	Use:   "resource [name]",
	Short: "Generate full resource (controller, service, repository, model, routes)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		generator.GenerateResource(args[0])
	},
}

func init() {
	GenerateCmd.AddCommand(resourceCmd)
}
