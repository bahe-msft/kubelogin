package token

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewChainedCredential(t *testing.T) {
	tests := []struct {
		name string
		opts *Options
	}{
		{
			name: "valid credential creation",
			opts: &Options{},
		},
		{
			name: "with tenant ID",
			opts: &Options{
				TenantID: "test-tenant-id",
			},
		},
		{
			name: "with custom cloud",
			opts: &Options{
				Environment: "AzureUSGovernmentCloud",
			},
		},
		{
			name: "with instance discovery disabled",
			opts: &Options{
				DisableInstanceDiscovery: true,
			},
		},
		{
			name: "with custom http client",
			opts: &Options{
				httpClient: http.DefaultClient,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cred, err := newChainedCredential(test.opts)

			require.NoError(t, err)
			require.NotNil(t, cred)
			assert.Equal(t, "ChainedCredential", cred.Name())
			assert.False(t, cred.NeedAuthenticate())

			chainedCredential, ok := cred.(*ChainedCredential)
			require.True(t, ok)
			assert.NotNil(t, chainedCredential.cred)

			record, err := cred.Authenticate(context.Background(), &policy.TokenRequestOptions{})
			assert.ErrorIs(t, err, errAuthenticateNotSupported)
			assert.Empty(t, record)
		})
	}
}

func TestChainedLoginIsSupported(t *testing.T) {
	opts := NewOptions(false)
	opts.LoginMethod = ChainedLogin
	opts.Timeout = time.Second

	require.NoError(t, opts.Validate())
	assert.Contains(t, strings.Split(GetSupportedLogins(), ", "), ChainedLogin)
}
