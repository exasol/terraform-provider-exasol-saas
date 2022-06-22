package provider

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_validateRegions(t *testing.T) {
	tests := []struct {
		name   string
		region string
		path   cty.Path
		want   bool
	}{
		{name: "Valid region", region: "us-west-1", want: true},
		{name: "Invalid region", region: "us-west-100", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateRegions(tt.region, tt.path)
			assert.Equal(t, !got.HasError(), tt.want)
		})
	}
}

func Test_validateSize(t *testing.T) {
	tests := []struct {
		name string
		size string
		path cty.Path
		want bool
	}{
		{name: "Valid size", size: "XS", want: true},
		{name: "Invalid size", size: "CS", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateRegions(tt.size, tt.path)
			assert.Equal(t, !got.HasError(), tt.want)
		})
	}
}
