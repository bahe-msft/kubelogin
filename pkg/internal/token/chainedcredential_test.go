package token

import (
	"testing"

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
				TenantID: "tenant-id",
			},
		},
		{
			name: "with instance discovery disabled",
			opts: &Options{
				DisableInstanceDiscovery: true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cred, err := newChainedCredential(test.opts)

			require.NoError(t, err)
			assert.NotNil(t, cred)
			assert.Equal(t, "ChainedCredential", cred.Name())
			assert.False(t, cred.NeedAuthenticate())
		})
	}
}
