# Chained Credential

This login mode uses Azure SDK's `DefaultAzureCredential` to try multiple credential types until one succeeds.

The default chain includes environment credentials, workload identity, managed identity, and Azure CLI credentials. This is useful when the same kubeconfig needs to work across local development, CI, and Azure-hosted environments without choosing a specific login mode for each environment.

## Usage

```sh
export KUBECONFIG=/path/to/kubeconfig

kubelogin convert-kubeconfig --login chained

kubectl get nodes
```

To get a token directly:

```sh
kubelogin get-token --login chained --server-id <server-id>
```

## Environment Credential

Set Azure SDK environment variables to authenticate as a service principal:

```sh
export AZURE_CLIENT_ID=<service-principal-client-id>
export AZURE_CLIENT_SECRET=<service-principal-client-secret>
export AZURE_TENANT_ID=<tenant-id>

kubelogin convert-kubeconfig --login chained
```

## Azure CLI Credential

For local development, sign in with Azure CLI before using chained login:

```sh
az login

kubelogin convert-kubeconfig --login chained
```

## Managed Identity And Workload Identity

When running in Azure environments with managed identity, or in AKS pods configured for workload identity, the chained credential can use the identity exposed by the platform.

## References

- https://learn.microsoft.com/azure/developer/go/sdk/authentication/credential-chains
- https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azidentity#DefaultAzureCredential
