package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestNwsSubnetPrivateExample(t *testing.T) {

	const (
		vpcID = "79eeb029-e396-4d9d-878a-338f8da07cb8"
	)

	testCases := []testCaseT{
		{
			[]string{genName()},
			[]string{"10.0.1.100/30"},
			"test.local",
			false,
			[]string{"80", "31000-31010"},
		},
		{
			[]string{genName(), genName()},
			[]string{"10.0.1.0/30", "10.0.1.10/30"},
			"test.local",
			false,
			[]string{"80", "31000-31010", "32000-32010"},
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
				"domain": cfg.domain,
				"vpc_id": vpcID,
				"public": cfg.public,
			},
		})

		defer terraform.Destroy(t, terraformOptions)
		terraform.InitAndApply(t, terraformOptions)

		validatePrivate(t, terraformOptions)
	}
}
