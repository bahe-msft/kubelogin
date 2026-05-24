package token

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewChainedCredential(t *testing.T) {
	tests := []struct {
		name        string
		opts        *Options
		expectError bool
	}{
		{
			name:        "valid credential creation",
			opts:        &Options{},
			expectError: false,
		},
		{
			name: "with persistent cache enabled",
			opts: &Options{
				UsePersistentCache: true,
			},
			expectError: false,
		},
		{
			name: "with instance discovery disabled",
			opts: &Options{
				DisableInstanceDiscovery: true,
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cred, err := newChainedCredential(test.opts)

			if test.expectError {
				require.Error(t, err)
				assert.Nil(t, cred)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, cred)
				assert.Equal(t, "ChainedCredential", cred.Name())
				assert.False(t, cred.NeedAuthenticate())
			}
		})
	}
}

func TestNewChainedCredentialPassesExplicitTenantID(t *testing.T) {
	const tenantID = "00000000-0000-0000-0000-000000000000"

	originalNewDefaultAzureCredential := newDefaultAzureCredential
	t.Cleanup(func() {
		newDefaultAzureCredential = originalNewDefaultAzureCredential
	})

	var capturedOptions *azidentity.DefaultAzureCredentialOptions
	newDefaultAzureCredential = func(opts *azidentity.DefaultAzureCredentialOptions) (*azidentity.DefaultAzureCredential, error) {
		capturedOptions = opts
		return originalNewDefaultAzureCredential(opts)
	}

	cred, err := newChainedCredential(&Options{TenantID: tenantID})

	require.NoError(t, err)
	require.NotNil(t, cred)
	require.NotNil(t, capturedOptions)
	assert.Equal(t, tenantID, capturedOptions.TenantID)
}
