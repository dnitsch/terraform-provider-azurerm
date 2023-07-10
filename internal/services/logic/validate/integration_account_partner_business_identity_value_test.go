// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validate

import (
	"strings"
	"testing"
)

func TestIntegrationAccountPartnerBusinessIdentityValue(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
	}{
		{
			input: "",
			valid: false,
		},
		{
			input: "test1",
			valid: true,
		},
		{
			input: "a2-.()b",
			valid: true,
		},
		{
			input: "a2&b",
			valid: false,
		},
		{
			input: "a b",
			valid: true,
		},
		{
			input: strings.Repeat("s", 127),
			valid: true,
		},
		{
			input: strings.Repeat("s", 128),
			valid: true,
		},
		{
			input: strings.Repeat("s", 129),
			valid: false,
		},
	}

	validationFunction := IntegrationAccountPartnerBusinessIdentityValue()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := validationFunction(tt.input, "")
			valid := err == nil
			if valid != tt.valid {
				t.Errorf("Expected valid status %t but got %t for input %s", tt.valid, valid, tt.input)
			}
		})
	}
}
