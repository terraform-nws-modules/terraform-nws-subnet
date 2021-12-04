package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestNwsSubnetVpcExample(t *testing.T) {

	const (
		vpc_id = "79eeb029-e396-4d9d-878a-338f8da07cb8"
	)

	testCases := []struct {
		name   []string
		cidr   []string
		domain []string
		public []bool
	}{
		{
			[]string{genName()},
			[]string{"10.0.1.50/30"},
			[]string{"test.local"},
			[]bool{false},
		},
		{
			[]string{genName(), genName()},
			[]string{"10.0.1.0/30", "10.0.1.100/30"},
			[]string{"test.local", "test.local"},
			[]bool{false, false},
		},
	}

	for _, testCase := range testCases {
		cfg := testCase

		terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
			TerraformDir: "../examples/vpc-single-private",
			// Variables to pass to our Terraform code using -var options
			Vars: map[string]interface{}{
				"name":   cfg.name,
				"cidr":   cfg.cidr,
				"domain": cfg.domain[0],
				"vpc_id": vpc_id,
				"public": cfg.public[0],
			},
		})

		defer terraform.Destroy(t, terraformOptions)
		terraform.InitAndApply(t, terraformOptions)

		validate(t, terraformOptions, cfg.domain)
	}
}
