package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

const (
	vpcID  = "fd7a2090-2e34-45b8-9eca-d9d27584475d "
	domain = "my.local"
)

type testCaseT struct {
	name        []string
	cidr        []string
	domain      string
	public      bool
	hasVpc      bool
	aclPortList []string
}

// validates Terraform output versus expected
func validateVpcPrivate(t *testing.T, opts *terraform.Options) {
	subnetID := terraform.Output(t, opts, "subnet_id")
	subnetACLID := terraform.Output(t, opts, "subnet_acl_id")
	subnetACLRuleID := terraform.Output(t, opts, "subnet_acl_rule_id")
	public := terraform.Output(t, opts, "subnet_public")
	hasVpc := terraform.Output(t, opts, "subnet_has_vpc")

	assert.NotEmpty(t, subnetID)
	assert.Equal(t, "[]", subnetACLID)
	assert.Equal(t, "[]", subnetACLRuleID)
	assert.Equal(t, "false", public)
	assert.Equal(t, "true", hasVpc)
}

func validateVpcPublic(t *testing.T, opts *terraform.Options) {
	subnetID := terraform.Output(t, opts, "subnet_id")
	subnetACLID := terraform.Output(t, opts, "subnet_acl_id")
	subnetACLRuleID := terraform.Output(t, opts, "subnet_acl_rule_id")
	public := terraform.Output(t, opts, "subnet_public")
	hasVpc := terraform.Output(t, opts, "subnet_has_vpc")

	assert.NotEmpty(t, subnetID)
	assert.NotEmpty(t, subnetACLID)
	assert.NotEmpty(t, subnetACLRuleID)
	assert.Equal(t, "true", public)
	assert.Equal(t, "true", hasVpc)
}

func genName() string {
	return fmt.Sprintf("test-subnet-%s", random.UniqueId())
}

func validateSimple(t *testing.T, opts *terraform.Options) {
	subnetID := terraform.Output(t, opts, "subnet_id")
	subnetACLID := terraform.Output(t, opts, "subnet_acl_id")
	subnetACLRuleID := terraform.Output(t, opts, "subnet_acl_rule_id")
	hasVpc := terraform.Output(t, opts, "subnet_has_vpc")

	assert.NotEmpty(t, subnetID)
	assert.Equal(t, "[]", subnetACLID)
	assert.Equal(t, "[]", subnetACLRuleID)
	assert.Equal(t, "false", hasVpc)
}
