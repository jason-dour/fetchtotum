# fetchtotum

Fetch secrets from cloud secrets managers.

## Purpose

`fetchtotum` is a CLI for retrieving secrets from cloud services' secrets managers.

## Development Plan

Proceeding with the following implementation plan for providers:

1. GCP
2. AWS
3. Azure

For each provider:

1. Use discovered credentials from the environment.
   1. This is our most common design pattern.  This is intended to be used in automation
        where the secret cannot be directly pulled by your code.
   2. For example: use this in a startup script for a VM or container to retrieve a
        secret for use in the script.
2. Evaluate possible future use cases or designs and place into development parking lot.
