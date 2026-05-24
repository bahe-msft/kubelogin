package token

import (
	"os"
	"testing"
)

func isolateAzureSDKEnv(t *testing.T) {
	t.Helper()

	for _, key := range []string{
		"AZURE_AUTHORITY_HOST",
		"AZURE_CLIENT_CERTIFICATE_PASSWORD",
		"AZURE_CLIENT_CERTIFICATE_PATH",
		"AZURE_CLIENT_ID",
		"AZURE_CLIENT_SECRET",
		"AZURE_FEDERATED_TOKEN_FILE",
		"AZURE_PASSWORD",
		"AZURE_TENANT_ID",
		"AZURE_TOKEN_CREDENTIALS",
		"AZURE_USERNAME",
		"IDENTITY_ENDPOINT",
		"IDENTITY_HEADER",
		"IDENTITY_SERVER_THUMBPRINT",
		"IMDS_ENDPOINT",
		"MSI_ENDPOINT",
		"MSI_SECRET",
	} {
		key := key
		value, ok := os.LookupEnv(key)
		os.Unsetenv(key)
		t.Cleanup(func() {
			if ok {
				os.Setenv(key, value)
			} else {
				os.Unsetenv(key)
			}
		})
	}
}
