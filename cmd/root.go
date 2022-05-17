// cmd/root
//
// Root command for fetchtotum.

package cmd

import (
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
		Args:    cobra.NoArgs,
		Version: version,
	}
)

// panicIfError - Panic if an error occurred.
func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

// Execute - Executes the root command.
func Execute() {
	err := rootCmd.Execute()
	panicIfError(err)
}
