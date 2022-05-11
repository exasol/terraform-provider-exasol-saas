package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	openapi "github.com/exasol/terraform-provider-exasol-saas/internal/client"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func handleApiError(ctx context.Context, err error, resp *http.Response) diag.Diagnostics {
	if err != nil {
		var apiError openapi.APIError
		if resp != nil {
			err := json.NewDecoder(resp.Body).Decode(&apiError)
			if err == nil {
				tflog.Error(ctx, apiError.Message)
				return diag.Errorf(apiError.Message)
			}

		} else {
			tflog.Error(ctx, fmt.Sprintf("Api error: %+v", err))
		}

		return diag.FromErr(err)
	}
	return nil
}

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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
