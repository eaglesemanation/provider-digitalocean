// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	droplet "github.com/eaglesemanation/provider-digitalocean/internal/controller/namespaced/droplet/droplet"
	image "github.com/eaglesemanation/provider-digitalocean/internal/controller/namespaced/image/image"
	providerconfig "github.com/eaglesemanation/provider-digitalocean/internal/controller/namespaced/providerconfig"
	ip "github.com/eaglesemanation/provider-digitalocean/internal/controller/namespaced/reservedip/ip"
	ipassignment "github.com/eaglesemanation/provider-digitalocean/internal/controller/namespaced/reservedip/ipassignment"
	ipv6 "github.com/eaglesemanation/provider-digitalocean/internal/controller/namespaced/reservedipv6/ipv6"
	ipv6assignment "github.com/eaglesemanation/provider-digitalocean/internal/controller/namespaced/reservedipv6/ipv6assignment"
	key "github.com/eaglesemanation/provider-digitalocean/internal/controller/namespaced/sshkey/key"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		droplet.Setup,
		image.Setup,
		providerconfig.Setup,
		ip.Setup,
		ipassignment.Setup,
		ipv6.Setup,
		ipv6assignment.Setup,
		key.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		droplet.SetupGated,
		image.SetupGated,
		providerconfig.SetupGated,
		ip.SetupGated,
		ipassignment.SetupGated,
		ipv6.SetupGated,
		ipv6assignment.SetupGated,
		key.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
