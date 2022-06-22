package provider

import (
	"fmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"regexp"
	"strings"
)

var validateName = validation.ToDiagFunc(validation.StringMatch(regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]+$"), "Allowed characters: a-z (lower and upper case), numbers and underscore. First character must be a-z (lower and upper case)"))

func validateSize(i interface{}, path cty.Path) diag.Diagnostics {
	sizes := []string{"XS", "S", "M", "L", "XL", "2XL", "3XL"}
	size := i.(string)
	if contains(sizes, size) {
		return nil
	}

	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity:      diag.Error,
		Summary:       fmt.Sprintf("Invalid size: %s", size),
		Detail:        fmt.Sprintf("Supported sizes are %s", strings.Join(sizes, ",")),
		AttributePath: path,
	})
	return diags
}

func validateRegions(i interface{}, path cty.Path) diag.Diagnostics {
	regions := []string{"eu-central-1", "eu-west-2", "us-west-1", "us-west-2", "us-east-1", "us-east-2"}
	region := i.(string)
	if contains(regions, region) {
		return nil
	}

	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity:      diag.Error,
		Summary:       fmt.Sprintf("Invalid region: %s", region),
		Detail:        fmt.Sprintf("Supported regions are %s", strings.Join(regions, ",")),
		AttributePath: path,
	})
	return diags
}
