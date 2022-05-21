// cmd/gcp
//
// gcp command for fetchtotum.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	projectFlag string

	gcpCmd = &cobra.Command{
		Use:   "gcp [-p project] secret_name",
		Short: "Retrieves secrets from Google Cloud Platform Secret Manager",
		Long:  `Retrieves secrets from Google Cloud Platform Secret Manager`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runGcp(args)
		},
		DisableFlagsInUseLine: true,
	}
)

// Initialization.
func init() {
	// Add the list command to the root command.
	rootCmd.AddCommand(gcpCmd)

	// Command Flags
	gcpCmd.Flags().StringVarP(&projectFlag, "project", "p", "", "GCP project for the secret; detect if not provided")
}

//
// Some quick hackery found on GH for consideration.
//
// secret, ok := os.LookupEnv("SECRET")
// if !ok {
// 	log.Fatalf("Environment variable SECRET is required")
// }
// ctx := context.Background()
// credentials, err := google.FindDefaultCredentials(ctx, compute.ComputeScope)
// if err != nil {
// 	fmt.Println(err)
// 	os.Exit(1)
// }
// c, err := secretmanager.NewClient(ctx)
// if err != nil {
// 	fmt.Println(err)
// 	os.Exit(2)
// }
// defer c.Close()

// req := &secretmanagerpb.AccessSecretVersionRequest{
// 	Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", credentials.ProjectID, secret),
// }
// res, err := c.AccessSecretVersion(ctx, req)
// if err != nil {
// 	log.Fatalf("unable to access secret-version: %v", err)
// }
// fmt.Println(string(res.Payload.Data))

// runGcp - runs the GCP command.
func runGcp(args []string) {
	exitIfError(fmt.Errorf("gcp not yet implemented (project:%s; args: %#v)", projectFlag, args))
}
