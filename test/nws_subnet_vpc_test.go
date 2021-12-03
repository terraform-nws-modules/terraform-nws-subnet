package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestNwsSubnetVpcExample(t *testing.T) {

	const (
		vpc_id = "db0b0f50-19b8-4d62-8f7d-c196a7353f4d"
	)

	testCases := []struct {
		name   []string
		cidr   []string
		domain []string
		public []bool
	}{
		// {
		// 	[]string{genName()},
		// 	[]string{"10.0.1.0/24"},
		// 	[]string{"my.local"},
		// 	[]bool{false},
		// },
		{
			[]string{genName(), genName()},
			[]string{"10.0.1.0/30", "10.0.1.100/30"},
			[]string{"my.local", "my.local"},
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
