package converter

import (
	"path/filepath"
	"testing"

	"github.com/Azure/kubelogin/pkg/internal/token"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd"
)

func TestConvertToChainedLogin(t *testing.T) {
	const (
		clusterName   = "aks1"
		otherCluster  = "aks2"
		serverID      = "server-id"
		clientID      = "client-id"
		tenantID      = "tenant-id"
		environment   = "AzurePublicCloud"
		authorityHost = "https://login.microsoftonline.com/"
	)

	tests := []struct {
		name          string
		overrideFlags map[string]string
		expectedArgs  []string
	}{
		{
			name: "does not inherit legacy client tenant or environment",
			overrideFlags: map[string]string{
				flagLoginMethod: token.ChainedLogin,
			},
			expectedArgs: []string{
				getTokenCommand,
				argLoginMethod, token.ChainedLogin,
				argServerID, serverID,
			},
		},
		{
			name: "passes explicit optional cloud flags but not client id",
			overrideFlags: map[string]string{
				flagLoginMethod:   token.ChainedLogin,
				flagClientID:      clientID,
				flagTenantID:      tenantID,
				flagEnvironment:   environment,
				flagAuthorityHost: authorityHost,
			},
			expectedArgs: []string{
				getTokenCommand,
				argLoginMethod, token.ChainedLogin,
				argServerID, serverID,
				argTenantID, tenantID,
				argEnvironment, environment,
				argAuthorityHost, authorityHost,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := createValidTestConfigs(
				clusterName,
				otherCluster,
				"",
				azureAuthProvider,
				map[string]string{
					cfgApiserverID: serverID,
					cfgClientID:    clientID,
					cfgTenantID:    tenantID,
					cfgEnvironment: environment,
					cfgConfigMode:  "0",
				},
				nil,
				"",
			)

			fs := &pflag.FlagSet{}
			o := Options{
				Flags: fs,
				configFlags: genericclioptions.NewTestConfigFlags().
					WithClientConfig(clientcmd.NewNonInteractiveClientConfig(*config, clusterName, &clientcmd.ConfigOverrides{}, nil)),
			}
			o.AddFlags(fs)
			for key, value := range test.overrideFlags {
				if err := o.setFlag(key, value); err != nil {
					t.Fatalf("unable to set flag %s: %v", key, err)
				}
			}

			kubeconfigFile := filepath.Join(t.TempDir(), "config")
			pathOptions := clientcmd.PathOptions{
				ExplicitFileFlag: "kubeconfig",
				LoadingRules: &clientcmd.ClientConfigLoadingRules{
					ExplicitPath: kubeconfigFile,
				},
			}

			if err := Convert(o, &pathOptions); err != nil {
				t.Fatalf("Convert returned error: %v", err)
			}

			validate(t, clusterName, config.AuthInfos[clusterName], test.expectedArgs, "", "", nil)
			validate(t, otherCluster, config.AuthInfos[otherCluster], test.expectedArgs, "", "", nil)
		})
	}
}
