package huaweicloud

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnshuaweicloud "github.com/libdns/huaweicloud"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *libdnshuaweicloud.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.huaweicloud",
		New: func() caddy.Module { return &Provider{new(libdnshuaweicloud.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.AccessKeyId = caddy.NewReplacer().ReplaceAll(p.Provider.AccessKeyId, "")
	p.Provider.SecretAccessKey = caddy.NewReplacer().ReplaceAll(p.Provider.SecretAccessKey, "")
	p.Provider.RegionId = caddy.NewReplacer().ReplaceAll(p.Provider.RegionId, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	huaweicloud {
//	    access_key_id <access_key_id>
//	    secret_access_key <secret_access_key>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "access_key_id":
				if d.NextArg() {
					p.Provider.AccessKeyId = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "secret_access_key":
				if d.NextArg() {
					p.Provider.SecretAccessKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "region_id":
				if d.NextArg() {
					p.Provider.RegionId = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.AccessKeyId == "" || p.Provider.SecretAccessKey == "" {
		return d.Err("missing access_key_id or secret_access_key")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
