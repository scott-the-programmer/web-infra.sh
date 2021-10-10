package test

import (
	"context"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

func GetStack(ctx context.Context) (*auto.Stack, error) {
	s, err := auto.SelectStackLocalSource(ctx, "web", ".")
	if err != nil {
		return nil, err
	}
	return &s, nil
}
