package reservedipv6

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("digitalocean_reserved_ipv6", func(r *config.Resource) {
		r.ShortGroup = ""
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "ip")
	})
	p.AddResourceConfigurator("digitalocean_reserved_ipv6_assignment", func(r *config.Resource) {
		r.ShortGroup = ""
		r.References["ip"] = config.Reference{
			TerraformName: "digitalocean_reserved_ipv6",
		}
		r.References["droplet_id"] = config.Reference{
			TerraformName: "digitalocean_droplet",
		}
	})
}
