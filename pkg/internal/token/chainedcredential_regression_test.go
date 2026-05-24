package token

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

func TestChainedLoginIsSupported(t *testing.T) {
	o := NewOptions(true)
	o.LoginMethod = ChainedLogin
	o.Timeout = time.Second

	if err := o.Validate(); err != nil {
		t.Fatalf("expected chained login to validate: %v", err)
	}
	if !strings.Contains(GetSupportedLogins(), ChainedLogin) {
		t.Fatalf("expected supported logins to include %q, got %q", ChainedLogin, GetSupportedLogins())
	}
}

func TestNewAzIdentityCredentialForChainedLogin(t *testing.T) {
	clearDefaultAzureCredentialEnv(t)

	provider, err := NewAzIdentityCredential(azidentity.AuthenticationRecord{}, &Options{
		LoginMethod: ChainedLogin,
		ServerID:    "server-id",
		Timeout:     time.Second,
	})
	if err != nil {
		t.Fatalf("expected chained provider creation to succeed: %v", err)
	}
	if provider == nil {
		t.Fatal("expected chained provider, got nil")
	}
	if provider.Name() != "ChainedCredential" {
		t.Fatalf("expected ChainedCredential provider, got %q", provider.Name())
	}
}

func TestChainedCredentialDoesNotRequireAuthenticate(t *testing.T) {
	clearDefaultAzureCredentialEnv(t)

	provider, err := newChainedCredential(&Options{Timeout: time.Second})
	if err != nil {
		t.Fatalf("expected chained credential creation to succeed: %v", err)
	}
	if provider.NeedAuthenticate() {
		t.Fatal("chained credential should not require a separate Authenticate step")
	}

	_, err = provider.Authenticate(context.Background(), &policy.TokenRequestOptions{})
	if err != errAuthenticateNotSupported {
		t.Fatalf("expected Authenticate to return errAuthenticateNotSupported, got %v", err)
	}
}

func clearDefaultAzureCredentialEnv(t *testing.T) {
	t.Helper()

	envVars := []string{
		"AZURE_AUTHORITY_HOST",
		"AZURE_CLIENT_CERTIFICATE_PASSWORD",
		"AZURE_CLIENT_CERTIFICATE_PATH",
		"AZURE_CLIENT_ID",
		"AZURE_CLIENT_SECRET",
		"AZURE_FEDERATED_TOKEN_FILE",
		"AZURE_PASSWORD",
		"AZURE_TENANT_ID",
		"AZURE_USERNAME",
		"IDENTITY_ENDPOINT",
		"IDENTITY_HEADER",
		"IMDS_ENDPOINT",
		"MSI_ENDPOINT",
		"MSI_SECRET",
	}

	originals := make(map[string]string, len(envVars))
	present := make(map[string]bool, len(envVars))
	for _, key := range envVars {
		originals[key], present[key] = os.LookupEnv(key)
		if err := os.Unsetenv(key); err != nil {
			t.Fatalf("failed to unset %s: %v", key, err)
		}
	}

	t.Cleanup(func() {
		for _, key := range envVars {
			var err error
			if present[key] {
				err = os.Setenv(key, originals[key])
			} else {
				err = os.Unsetenv(key)
			}
			if err != nil {
				t.Fatalf("failed to restore %s: %v", key, err)
			}
		}
	})
}
