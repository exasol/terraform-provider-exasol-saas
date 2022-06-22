package provider

import (
	openapi "github.com/exasol/terraform-provider-exasol-saas/internal/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_contains(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		expected string
		want     bool
	}{
		{name: "Contains value", values: []string{"A", "AB"}, expected: "A", want: true},
		{name: "Not contains value", values: []string{"A", "AB"}, expected: "B", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, contains(tt.values, tt.expected), tt.want)
		})
	}
}

func Test_containsStatus(t *testing.T) {
	tests := []struct {
		name     string
		values   []openapi.Status
		expected openapi.Status
		want     bool
	}{
		{name: "Contains value", values: []openapi.Status{openapi.RUNNING, openapi.STOPPED}, expected: openapi.RUNNING, want: true},
		{name: "Not contains value", values: []openapi.Status{openapi.RUNNING, openapi.STOPPED}, expected: openapi.STOPPING, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, containsStatus(tt.values, tt.expected), tt.want)
		})
	}
}

func Test_statusesToString(t *testing.T) {
	tests := []struct {
		name     string
		statuses []openapi.Status
		want     string
	}{
		{name: "Contains value", statuses: []openapi.Status{openapi.RUNNING, openapi.STOPPED}, want: "running, stopped"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, statusesToString(tt.statuses), tt.want)
		})
	}
}
