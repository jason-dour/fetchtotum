// cmd/root
//
// Root command for fetchtotum.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// version - the version of the program; injected at compile time.
	version string

	// basename - the basename of the program; injected at compile time.
	basename string

	// Cobra definition.
	rootCmd = &cobra.Command{
		Use:     basename,
		Short:   "Fetch secrets from cloud secret managers.",
		Long:    "Fetch secrets from cloud secret managers.",
		Args:    cobra.NoArgs,
		Version: version,
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
		DisableFlagsInUseLine: true,
	}
)

// exitIfError - Exit if an error occurred.
func exitIfError(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
}

// Execute - Executes the root command.
func Execute() {
	err := rootCmd.Execute()
	exitIfError(err)
}
