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
		Long:  "Retrieves secrets from Amazon Web Services Secrets Manager",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			runAws(args)
		},
		DisableFlagsInUseLine: true,
	}
)

// Initialization.
func init() {
	// Add the list command to the root command.
	rootCmd.AddCommand(awsCmd)
}

// runAws - runs the AWS command.
func runAws(args []string) {
	exitIfError(fmt.Errorf("aws not yet implemented"))
}
