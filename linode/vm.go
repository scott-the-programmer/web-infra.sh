package linode

import (
	"github.com/pulumi/pulumi-linode/sdk/v3/go/linode"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func NewVm(ctx *pulumi.Context, pub string) (*linode.Instance, error) {
	instance, err := linode.NewInstance(ctx, "vm-web-sh", &linode.InstanceArgs{
		Type:           pulumi.String("g6-nanode-1"),
		Region:         pulumi.String("us-east"),
		Image:          pulumi.String("linode/ubuntu20.04"),
		AuthorizedKeys: pulumi.StringArray{pulumi.String(pub)},
	})

	if err != nil {
		return nil, err
	}

	return instance, err
}
