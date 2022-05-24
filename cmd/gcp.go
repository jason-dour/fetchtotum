// cmd/gcp
//
// gcp command for fetchtotum.

package cmd

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var (
	// projectFlag - GCP project for the secret; detect if not provided.
	projectFlag string

	// versionFlag - the version of the secret to retrieve.
	versionFlag string

	// gcpCmd - Cobra definition for GCP command.
	gcpCmd = &cobra.Command{
		Use:   "gcp [-p project_id] secret_name",
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
	gcpCmd.Flags().StringVarP(&projectFlag, "project", "p", "", "GCP project ID for the secret; detect if not provided")
	gcpCmd.Flags().StringVarP(&versionFlag, "version", "v", "latest", "version for the secret; default: latest")
}

// versionSelected - Returns the version requested.
func versionSelected() string {
	debug("versionSelected(): begin")
	version := ""
	// Check if the version flag was provided...
	if len(versionFlag) == 0 {
		// ...and if not, use the default...
		version = "latest"
		debug("versionSelected(): versionFlag is empty; using default: %s", version)
	} else {
		// ...otherwise, use the version flag.
		version = versionFlag
		debug("versionSelected(): versionFlag is not empty; using: %s", version)
	}
	debug("versionSelected(): end: version=%s", version)
	return version
}

// determineProject - Returns the project name.
func determineProject(credentials *google.Credentials) (string, error) {
	debug("determineProject(): begin")
	// Check if the project flag was provided...
	if len(projectFlag) != 0 {
		// ...and if so, use it; it overrides defaults...
		debug("determineProject(): projectFlag is set; using projectFlag: %s", projectFlag)
		debug("determineProject(): end: project=%s", projectFlag)
		return projectFlag, nil
	} else if len(credentials.ProjectID) != 0 {
		// ...otherwise, if the credentials have a project ID, and use it...
		debug("determineProject(): projectFlag is not set; project from credentials: %s", credentials.ProjectID)
		debug("determineProject(): end: project=%s", credentials.ProjectID)
		return credentials.ProjectID, nil
	} else {
		// ...otherwise, fail.
		debug("determineProject(): projectFlag is not set; credentials do not provide project")
		return "", fmt.Errorf("project not specified by flag or credentials")
	}
}

// runGcp - runs the GCP command.
func runGcp(args []string) {
	debug("runGcp(): begin: []args=%s", args)

	// Create context.
	ctx := context.Background()

	// Find credentials.
	debug("runGcp(): looking up credentials")
	credentials, err := google.FindDefaultCredentials(ctx, compute.ComputeScope)
	exitIfError(err)
	debug("runGcp(): credentials: %+v", credentials)

	// Determine project.
	debug("runGcp(): determine project")
	project, err := determineProject(credentials)
	exitIfError(err)
	debug("runGcp(): project: %s", project)

	// Select version.
	debug("runGcp(): select version")
	version := versionSelected()
	debug("runGcp(): version: %s", version)

	// Create client.
	debug("runGcp(): creating client")
	client, err := secretmanager.NewClient(ctx)
	exitIfError(err)
	defer client.Close()
	debug("runGcp(): client: %+v", client)

	// Create secret request.
	debug("runGcp(): creating request")
	request := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/%s", project, args[0], version),
	}
	debug("runGcp(): request: %#v", request)

	// Call secret manager with the request.
	debug("runGcp(): issuing request")
	response, err := client.AccessSecretVersion(ctx, request)
	exitIfError(err)
	debug("runGcp(): response: %#v", response)

	// Print the secret.
	debug("runGcp(): output secret.")
	fmt.Printf("%s\n", response.Payload.Data)

	// All done.
	debug("runGcp(): end")
}
