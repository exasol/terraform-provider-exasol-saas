package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceDatabase(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceDatabase,
				ExpectError: regexp.MustCompile("Invalid credentials"),
			},
		},
	})
}

const testAccResourceDatabase = `
resource "exasol_database" "mydb" {
  name   = "dbName"
  region = "us-east-2"
  size   = "XS"
}
`
