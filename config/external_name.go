package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"digitalocean_ssh_key": config.IdentifierFromProvider,

	"digitalocean_custom_image": config.IdentifierFromProvider,

	"digitalocean_droplet": config.IdentifierFromProvider,

	"digitalocean_reserved_ip":              config.ParameterAsIdentifier("ip_address"),
	"digitalocean_reserved_ip_assignment":   config.IdentifierFromProvider,
	"digitalocean_reserved_ipv6":            config.ParameterAsIdentifier("ip"),
	"digitalocean_reserved_ipv6_assignment": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
