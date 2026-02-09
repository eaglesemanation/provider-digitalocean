package image

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("digitalocean_custom_image", func(r *config.Resource) {
		r.ShortGroup = "image"
	})
}
