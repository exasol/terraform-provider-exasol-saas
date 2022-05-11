package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceIp(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceIp,
				ExpectError: regexp.MustCompile("Invalid credentials"),
			},
		},
	})
}

const testAccResourceIp = `
resource "exasol_network_allow_list" "ip1" {
  name       = "companyNetwork"
  cidr_block = "127.0.0.1/32"
}
`
