package generate

import "github.com/spf13/cobra"

var GenerateCmd = &cobra.Command{
	Use:   "g",
	Short: "Generate resources",
}

func Init(root *cobra.Command) {
	root.AddCommand(GenerateCmd)
}
