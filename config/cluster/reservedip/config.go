package reservedip

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("digitalocean_reserved_ip", func(r *config.Resource) {
		r.ShortGroup = "reservedip"
		r.References["droplet_id"] = config.Reference{
			TerraformName: "digitalocean_droplet",
		}
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "ip_address")
	})
	p.AddResourceConfigurator("digitalocean_reserved_ip_assignment", func(r *config.Resource) {
		r.ShortGroup = "reservedip"
		r.References["ip_address"] = config.Reference{
			TerraformName: "digitalocean_reserved_ip",
		}
		r.References["droplet_id"] = config.Reference{
			TerraformName: "digitalocean_droplet",
		}
	})
}
