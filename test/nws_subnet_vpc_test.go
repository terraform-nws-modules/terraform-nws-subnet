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

	cfg := make(map[string]string)

	cfg["name"] = fmt.Sprintf("terratest-subnet-%s", random.UniqueId())
	cfg["cidr"] = "10.0.1.0/24"
	cfg["domain"] = "alfa.local"
	cfg["vpc_id"] = "aba8788b-ec7c-4e6f-a873-b38840e3c834"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/vpc",
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name":   cfg["name"],
			"cidr":   cfg["cidr"],
			"domain": cfg["domain"],
			"vpc_id": cfg["vpc_id"],
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	id := terraform.Output(t, terraformOptions, "id")
	domain := terraform.Output(t, terraformOptions, "domain")

	assert.Equal(t, cfg["domain"], domain)
	assert.NotEmpty(t, id)
	assert.NotNil(t, id)
}
