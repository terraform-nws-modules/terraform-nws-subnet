package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestNwsSubnetPrivateExample(t *testing.T) {

	testCases := []testCaseT{
		{
			[]string{genName()},
			[]string{"10.0.1.100/30"},
			domain,
			false,
			[]string{"80", "31000-31010"},
		},
		{
			[]string{genName(), genName()},
			[]string{"10.0.1.0/30", "10.0.1.10/30"},
			domain,
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
