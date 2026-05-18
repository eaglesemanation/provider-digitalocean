package sshkey

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("digitalocean_ssh_key", func(r *config.Resource) {
		r.ShortGroup = ""
	})
}
