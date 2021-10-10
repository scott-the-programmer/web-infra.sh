package linode

import (
	"context"
	"net"
	"testing"
	"time"

	"web-infra.sh/test"
)

func TestLinodeMachinePort(t *testing.T) {
	ctx := context.Background()
	s, err := test.GetStack(ctx)
	if err != nil {
		t.Error(err)
	}

	outputs, err := s.Outputs(ctx)
	if err != nil {
		t.Error(err)
	}

	host := outputs["vmIpAddress"].Value.(string)

	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, "443"), timeout)
	if err != nil {
		t.Error(err)
	}

	conn.Close()
}
