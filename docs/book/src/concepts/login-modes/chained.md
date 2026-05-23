# Chained Credential

This login mode uses Azure SDK's `DefaultAzureCredential` to try multiple Azure authentication mechanisms in order until one returns a token.

The credential chain follows the Azure SDK for Go order:

1. **Environment Credential** - Authenticates using environment variables such as `AZURE_CLIENT_ID`, `AZURE_CLIENT_SECRET`, and `AZURE_TENANT_ID`
2. **Workload Identity Credential** - Authenticates using workload identity when the Azure workload identity webhook configuration is present
3. **Managed Identity Credential** - Authenticates using the managed identity assigned to the Azure resource
4. **Azure CLI Credential** - Authenticates using the signed-in Azure CLI user from `az login`
5. **Azure Developer CLI Credential** - Authenticates using the signed-in Azure Developer CLI user from `azd auth login`

The first successful credential is used. Subsequent credentials in the chain are not attempted for that token request.

> ### NOTE
>
> Chained credential mode does not require additional kubelogin-specific flags. Configure one of the Azure SDK supported credential sources in the environment where kubelogin runs.

## Usage Examples

### Basic Usage

```sh
export KUBECONFIG=/path/to/kubeconfig

kubelogin convert-kubeconfig -l chained

kubectl get nodes
```

### Using Environment Variables for Service Principal

```sh
export KUBECONFIG=/path/to/kubeconfig
export AZURE_CLIENT_ID=<service-principal-client-id>
export AZURE_CLIENT_SECRET=<service-principal-client-secret>
export AZURE_TENANT_ID=<tenant-id>

kubelogin convert-kubeconfig -l chained

kubectl get nodes
```

### Using Azure CLI Authentication

```sh
az login

export KUBECONFIG=/path/to/kubeconfig

kubelogin convert-kubeconfig -l chained

kubectl get nodes
```

### Direct Token Retrieval

```sh
kubelogin get-token \
  --login chained \
  --server-id <cluster-server-id>
```

## How It Works

`DefaultAzureCredential` automatically detects the first available credential source:

1. Checks environment variables for service principal configuration
2. Checks workload identity configuration
3. Checks managed identity from Azure resource metadata
4. Checks Azure CLI cached credentials
5. Checks Azure Developer CLI cached credentials

## References

- https://learn.microsoft.com/en-us/azure/developer/go/sdk/authentication/credential-chains
- https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azidentity#DefaultAzureCredential
