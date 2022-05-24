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
	gcpCmd.Flags().StringVarP(&versionFlag, "version", "v", "latest", "version for the secret; default: latest")
}

// versionSelected - Returns the version requested.
func versionSelected() string {
	debug("versionSelected(): begin")
	version := ""
	if len(versionFlag) == 0 {
		version = "latest"
		debug("versionSelected(): versionFlag is empty; using default: %s", version)
	} else {
		version = versionFlag
		debug("versionSelected(): versionFlag is not empty; using: %s", version)
	}
	debug("versionSelected(): end: version=%s", version)
	return version
}

// determineProject - Returns the project name.
func determineProject(credentials *google.Credentials) (string, error) {
	debug("determineProject(): begin")
	project := ""
	// Project flag overrides the project in the credentials.
	if len(projectFlag) != 0 {
		debug("determineProject(): projectFlag is set; using projectFlag: %s", projectFlag)
		return credentials.ProjectID, nil
	} else if len(credentials.ProjectID) != 0 {
		debug("determineProject(): projectFlag is not set; project from credentials: %s", project)
		return credentials.ProjectID, nil
	} else {
		debug("determineProject(): projectFlag is not set; credentials do not provide project")
		return "", fmt.Errorf("project not specified by flag or credentials")
	}
}

// runGcp - runs the GCP command.
func runGcp(args []string) {
	debug("runGcp(): begin: []args=%s", args)

	ctx := context.Background()

	debug("runGcp(): looking up credentials")
	credentials, err := google.FindDefaultCredentials(ctx, compute.ComputeScope)
	exitIfError(err)
	debug("runGcp(): credentials: %+v", credentials)

	debug("runGcp(): determine project")
	project, err := determineProject(credentials)
	exitIfError(err)
	debug("runGcp(): project: %s", project)

	debug("runGcp(): creating client")
	client, err := secretmanager.NewClient(ctx)
	exitIfError(err)
	defer client.Close()
	debug("runGcp(): client: %+v", client)

	debug("runGcp(): creating request")
	request := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/%s", project, args[0], versionSelected()),
	}
	debug("runGcp(): request: %#v", request)

	debug("runGcp(): issuing request")
	response, err := client.AccessSecretVersion(ctx, request)
	exitIfError(err)
	debug("runGcp(): response: %#v", response)

	debug("runGcp(): output secret.")
	fmt.Printf("%s\n", response.Payload.Data)

	debug("runGcp(): end")
}
