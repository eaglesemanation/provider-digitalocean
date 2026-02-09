package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	dropletCluster "github.com/eaglesemanation/provider-digitalocean/config/cluster/droplet"
	imageCluster "github.com/eaglesemanation/provider-digitalocean/config/cluster/image"
	reservedipCluster "github.com/eaglesemanation/provider-digitalocean/config/cluster/reservedip"
	reservedipv6Cluster "github.com/eaglesemanation/provider-digitalocean/config/cluster/reservedipv6"
	sshkeyCluster "github.com/eaglesemanation/provider-digitalocean/config/cluster/sshkey"
	dropletNamespaced "github.com/eaglesemanation/provider-digitalocean/config/namespaced/droplet"
	imageNamespaced "github.com/eaglesemanation/provider-digitalocean/config/namespaced/image"
	reservedipNamespaced "github.com/eaglesemanation/provider-digitalocean/config/namespaced/reservedip"
	reservedipv6Namespaced "github.com/eaglesemanation/provider-digitalocean/config/namespaced/reservedipv6"
	sshkeyNamespaced "github.com/eaglesemanation/provider-digitalocean/config/namespaced/sshkey"
)

const (
	resourcePrefix = "digitalocean"
	modulePath     = "github.com/eaglesemanation/provider-digitalocean"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("digitalocean.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		imageCluster.Configure,
		sshkeyCluster.Configure,
		dropletCluster.Configure,
		reservedipCluster.Configure,
		reservedipv6Cluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("digitalocean.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		imageNamespaced.Configure,
		sshkeyNamespaced.Configure,
		dropletNamespaced.Configure,
		reservedipNamespaced.Configure,
		reservedipv6Namespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
