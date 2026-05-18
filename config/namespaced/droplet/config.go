package droplet

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("digitalocean_droplet", func(r *config.Resource) {
		r.ShortGroup = ""
		r.References["image"] = config.Reference{
			TerraformName: "digitalocean_custom_image",
		}
		r.References["ssh_keys"] = config.Reference{
			TerraformName: "digitalocean_ssh_key",
		}
	})
}
