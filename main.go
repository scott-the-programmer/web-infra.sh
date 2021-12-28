package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"web-infra.sh/cloudflare"
	"web-infra.sh/linode"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		pub, ok := os.LookupEnv("WEB_INFRA_PUB_KEY")
		if !ok {
			return errors.New("Missing WEB_INFRA_PUB_KEY")
		}

		dns, ok := os.LookupEnv("WEB_INFRA_DNS")
		if !ok {
			return errors.New("Missing WEB_INFRA_DNS")
		}

		zoneId, ok := os.LookupEnv("CLOUDFLARE_ZONE_ID")
		if !ok {
			return errors.New("Missing CLOUDFLARE_ZONE_ID")
		}

		instance, err := linode.NewVm(ctx, pub)

		if err != nil {
			return err
		}

		_, err = linode.NewFirewall(ctx, instance)
		if err != nil {
			return err
		}

		_, err = cloudflare.NewARecord(ctx, cloudflare.ARecordArgs{
			Source: dns,
			Target: instance.IpAddress,
			ZoneId: zoneId,
		})

		if err != nil {
			return err
		}

		_, err = cloudflare.NewARecord(ctx, cloudflare.ARecordArgs{
			Source: fmt.Sprintf("%s.%s", "www", dns),
			Target: instance.IpAddress,
			ZoneId: zoneId,
		})

		if err != nil {
			return err
		}

		ctx.Export("aRecordSource", pulumi.String(dns))
		ctx.Export("aRecordTarget", instance.IpAddress)
		ctx.Export("vmIpAddress", instance.IpAddress)
		return nil
	})
}
