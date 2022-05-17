// cmd/gcp
//
// gcp command for fetchtotum.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	gcpCmd = &cobra.Command{
		Use:   "gcp",
		Short: "Retrieves secrets from Google Cloud Platform Secret Manager",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			runGcp(args)
		},
	}
)

// Initialization.
func init() {
	// Add the list command to the root command.
	rootCmd.AddCommand(gcpCmd)
}

// runList() - runs the list command.
func runGcp(args []string) {
	fmt.Println("gcp not yet implemented")
}
