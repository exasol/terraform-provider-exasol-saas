package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceCluster(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceCluster,
				ExpectError: regexp.MustCompile("Invalid credentials"),
			},
		},
	})
}

const testAccResourceCluster = `
resource "exasol_cluster" "workercluster01" {
  database = "dbid"
  size     = "XS"
  name     = "WorkerClusterName"
}

`
