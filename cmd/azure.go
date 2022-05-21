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
		Long:  "Retrieves secrets from Azure Key Vault",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			runAzure(args)
		},
		DisableFlagsInUseLine: true,
	}
)

// Initialization.
func init() {
	// Add the list command to the root command.
	rootCmd.AddCommand(azureCmd)
}

// runAzure - runs the Azure command.
func runAzure(args []string) {
	exitIfError(fmt.Errorf("azure not yet implemented"))
}
