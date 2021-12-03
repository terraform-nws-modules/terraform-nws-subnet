package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestNwsSubnetVpcExample(t *testing.T) {
	t.Parallel()

	const (
		domain = "my.local"
		vpc_id = "db0b0f50-19b8-4d62-8f7d-c196a7353f4d"
	)

	cfg := struct {
		name []string
		cidr []string
	}{
		[]string{genName()},
		[]string{"10.0.1.0/24"},
	}

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/vpc-single-net",
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name":   cfg.name,
			"cidr":   cfg.cidr,
			"domain": domain,
			"vpc_id": vpc_id,
		},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	validate(t, terraformOptions, domain)
}
