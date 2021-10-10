package cloudflare

import (
	"context"
	"fmt"
	"testing"

	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
	"web-infra.sh/test"
)

func TestARecords(t *testing.T) {

	ctx := context.Background()
	s, err := test.GetStack(ctx)
	if err != nil {
		t.Error(err)
	}

	outputs, err := s.Outputs(ctx)
	if err != nil {
		t.Fatal(err)
	}

	recordSource := outputs["aRecordSource"]

	sources := []struct {
		url string
	}{
		{fmt.Sprintf("%s.", recordSource.Value)},
		{fmt.Sprintf("%s.%s.", "www", recordSource.Value)},
	}

	config, err := dns.ClientConfigFromFile("/etc/resolv.conf")
	if err != nil {
		assert.Error(t, err)
	}

	for _, s := range sources {
		dnsClient := new(dns.Client)
		msg := new(dns.Msg)
		msg.SetQuestion(s.url, dns.TypeA)
		msg.RecursionDesired = true
		result, _, err := dnsClient.Exchange(msg, config.Servers[0]+":"+config.Port)
		if err != nil {
			assert.Error(t, err)
		}

		if len(result.Answer) == 0 {
			assert.Fail(t, fmt.Sprintf("no record found for %s", s))
		}

		assert.NotNil(t, result.Answer[0].(*dns.A).A.String())
	}
}
