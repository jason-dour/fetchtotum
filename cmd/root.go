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

	// debugFlag - debug flag.
	debugFlag bool

	// Cobra definition.
	rootCmd = &cobra.Command{
		Use:     basename + " [-d | --debug] command",
		Short:   "Fetch secrets from cloud secret managers.",
		Long:    "Fetch secrets from cloud secret managers.",
		Args:    cobra.NoArgs,
		Version: version,
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
		// DisableFlagsInUseLine: true,
	}
)

// Initialization.
func init() {
	// Command Flags
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "d", false, "debug mode")
}

// exitIfError - Exit if an error occurred.
func exitIfError(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
}

// debug - Prints a debug message if debug mode is enabled.
func debug(format string, args ...interface{}) {
	if debugFlag {
		fmt.Printf("debug: "+format+"\n", args...)
	}
}

// Execute - Executes the root command.
func Execute() {
	debug("Execute()")
	err := rootCmd.Execute()
	exitIfError(err)
}
