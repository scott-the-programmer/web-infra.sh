package linode

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
	"web-infra.sh/test"
)

// Ports should not be accessible outside of cloudflare
func TestLinodeFirewall(t *testing.T) {

	ports := []string{"22", "80", "443"}
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

	for _, element := range ports {
		_, err := net.DialTimeout("tcp", net.JoinHostPort(host, element), timeout)
		assert.Error(t, err)
	}
}

//Ports should be available via DNS (cloudflare)
func TestLinodeFirewallWithCloudflare(t *testing.T) {

	ports := []string{"80", "443"}
	ctx := context.Background()
	s, err := test.GetStack(ctx)
	if err != nil {
		t.Error(err)
	}

	outputs, err := s.Outputs(ctx)
	if err != nil {
		t.Error(err)
	}

	host := outputs["aRecordSource"].Value.(string)

	timeout := time.Second

	for _, element := range ports {
		conn, err := dns.DialTimeout("tcp", net.JoinHostPort(host, element), timeout)
		if err != nil {
			t.Error(err)
		}
		assert.NotNil(t, conn)
		conn.Close()
	}
}
