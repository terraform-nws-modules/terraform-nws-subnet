package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestNwsSubnetVpcExample(t *testing.T) {
	t.Parallel()

	const (
		cidr       = "10.0.1.0/24"
		exp_domain = "alfa.local"
		vpc_id     = "aba8788b-ec7c-4e6f-a873-b38840e3c834"
	)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/vpc",
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name":   fmt.Sprintf("terratest-subnet-%s", random.UniqueId()),
			"cidr":   cidr,
			"domain": exp_domain,
			"vpc_id": vpc_id,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	id := terraform.Output(t, terraformOptions, "id")
	act_domain := terraform.Output(t, terraformOptions, "domain")

	assert.Equal(t, exp_domain, act_domain)
	assert.NotEmpty(t, id)
	assert.NotNil(t, id)
}
