// cmd/azure
//
// azure command for fetchtotum.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	azureCmd = &cobra.Command{
		Use:   "azure",
		Short: "Retrieves secrets from Azure Key Vault",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			runazure(args)
		},
	}
)

// Initialization.
func init() {
	// Add the list command to the root command.
	rootCmd.AddCommand(azureCmd)
}

// runList() - runs the list command.
func runazure(args []string) {
	fmt.Println("azure not yet implemented")
}
