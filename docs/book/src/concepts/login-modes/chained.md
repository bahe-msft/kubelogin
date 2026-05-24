# Chained

This login mode uses the Azure SDK for Go
[`DefaultAzureCredential`](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azidentity#DefaultAzureCredential)
chain to get a token for the Kubernetes cluster.

With the current `azidentity` version used by kubelogin (`v1.8.0`), the chain tries these credentials in order and stops after one returns a token:

1. Environment
2. Workload Identity
3. Managed Identity
4. Azure CLI
5. Azure Developer CLI

Most credential inputs for this mode come from Azure SDK environment variables or from the signed-in Azure tools, not from kubelogin flags. For example, environment and workload identity credentials use variables such as `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, `AZURE_CLIENT_SECRET`, and `AZURE_FEDERATED_TOKEN_FILE`; Azure CLI and Azure Developer CLI use their own login state from `az login` or `azd auth login`.

`kubelogin` still needs the cluster server application ID through `--server-id`, which is usually already present in kubeconfigs downloaded from AKS. `kubelogin` does not cache tokens for chained login mode; caching, when any, is managed by the credential selected by the Azure SDK.

## Usage Examples

### Using Azure CLI login state

```sh
az login

export KUBECONFIG=/path/to/kubeconfig

kubelogin convert-kubeconfig -l chained

kubectl get nodes
```

### Using Azure SDK environment variables

Configure the Azure SDK environment variables from a secret manager or CI/CD
environment injection before running `kubelogin`. Avoid typing client secrets or
certificate paths directly in shell commands so they are not captured in shell
history, terminal logs, or copied examples.

```sh
export KUBECONFIG=/path/to/kubeconfig

kubelogin convert-kubeconfig -l chained

kubectl get nodes
```

## References

- https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azidentity#DefaultAzureCredential
- https://learn.microsoft.com/azure/developer/go/sdk/authentication/credential-chains
