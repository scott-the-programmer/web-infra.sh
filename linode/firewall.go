package linode

import (
	"strconv"

	"github.com/pulumi/pulumi-linode/sdk/v3/go/linode"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func NewFirewall(ctx *pulumi.Context, virtualMachine *linode.Instance) (*linode.Firewall, error) {

	id := virtualMachine.ID().ToStringOutput().ApplyT(func(id string) int {
		val, err := strconv.ParseInt(id, 0, 32)
		if err != nil {
			return 1
		}
		return int(val)
	})

	var resolvedId pulumi.IntOutput

	switch id.(type) {
	default:
		resolvedId = pulumi.Int(1).ToIntOutput()
	case pulumi.IntOutput:
		resolvedId = id.(pulumi.IntOutput)
	}

	fw, err := linode.NewFirewall(ctx, "fw-web-sh", &linode.FirewallArgs{
		Label: pulumi.String("fw-web-sh"),
		Inbounds: linode.FirewallInboundArray{
			&linode.FirewallInboundArgs{
				Label:    pulumi.String("allow-http"),
				Action:   pulumi.String("ACCEPT"),
				Protocol: pulumi.String("TCP"),
				Ports:    pulumi.String("80"),
				Ipv4s:    pulumi.ToStringArray(getCloudflareIPv4()),
				Ipv6s:    pulumi.ToStringArray(getCloudflareIPv6()),
			},
			&linode.FirewallInboundArgs{
				Label:    pulumi.String("allow-https"),
				Action:   pulumi.String("ACCEPT"),
				Protocol: pulumi.String("TCP"),
				Ports:    pulumi.String("443"),
				Ipv4s:    pulumi.ToStringArray(getCloudflareIPv4()),
				Ipv6s:    pulumi.ToStringArray(getCloudflareIPv6()),
			},
		},
		InboundPolicy:  pulumi.String("DROP"),
		OutboundPolicy: pulumi.String("ACCEPT"),
		Linodes: pulumi.IntArray{
			resolvedId,
		},
	})

	return fw, err
}

func getCloudflareIPv4() []string {
	return []string{"103.21.244.0/22",
		"103.22.200.0/22",
		"103.31.4.0/22",
		"104.16.0.0/13",
		"104.24.0.0/14",
		"108.162.192.0/18",
		"131.0.72.0/22",
		"141.101.64.0/18",
		"162.158.0.0/15",
		"172.64.0.0/13",
		"173.245.48.0/20",
		"188.114.96.0/20",
		"190.93.240.0/20",
		"197.234.240.0/22",
		"198.41.128.0/17"}
}

func getCloudflareIPv6() []string {
	return []string{"2400:cb00::/32",
		"2606:4700::/32",
		"2803:f800::/32",
		"2405:b500::/32",
		"2405:8100::/32",
		"2c0f:f248::/32",
		"2a06:98c0::/29"}
}
