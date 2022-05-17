// cmd/aws
//
// aws command for fetchtotum.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	awsCmd = &cobra.Command{
		Use:   "aws",
		Short: "Retrieves secrets from Amazon Web Services Secrets Manager",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			runAws(args)
		},
	}
)

// Initialization.
func init() {
	// Add the list command to the root command.
	rootCmd.AddCommand(awsCmd)
}

// runList() - runs the list command.
func runAws(args []string) {
	fmt.Println("aws not yet implemented")
}
