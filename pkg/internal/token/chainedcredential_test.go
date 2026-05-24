package token

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/stretchr/testify/assert"
)

func TestChainedCredential(t *testing.T) {
	isolateAzureSDKEnv(t)

	provider, err := newChainedCredential(&Options{
		TenantID:                 "tenant-id",
		Environment:              "AzurePublicCloud",
		DisableInstanceDiscovery: true,
	})
	assert.NoError(t, err)
	assert.NotNil(t, provider)
	assert.Equal(t, "ChainedCredential", provider.Name())
	assert.False(t, provider.NeedAuthenticate())

	record, err := provider.Authenticate(context.Background(), &policy.TokenRequestOptions{})
	assert.ErrorIs(t, err, errAuthenticateNotSupported)
	assert.Empty(t, record)
}
