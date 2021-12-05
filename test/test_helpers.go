package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

type testCaseT struct {
	name        []string
	cidr        []string
	domain      string
	public      bool
	aclPortList []string
}

// validates Terraform output versus expected
func validatePrivate(t *testing.T, opts *terraform.Options) {
	subnetID := terraform.Output(t, opts, "subnet_id")
	subnetACLID := terraform.Output(t, opts, "subnet_acl_id")
	subnetACLRuleID := terraform.Output(t, opts, "subnet_acl_rule_id")
	public := terraform.Output(t, opts, "subnet_public")

	assert.NotEmpty(t, subnetID)
	assert.Equal(t, "[]", subnetACLID)
	assert.Equal(t, "[]", subnetACLRuleID)
	assert.Equal(t, "false", public)
}

func validatePublic(t *testing.T, opts *terraform.Options) {
	subnetID := terraform.Output(t, opts, "subnet_id")
	subnetACLID := terraform.Output(t, opts, "subnet_acl_id")
	subnetACLRuleID := terraform.Output(t, opts, "subnet_acl_rule_id")
	public := terraform.Output(t, opts, "subnet_public")

	assert.NotEmpty(t, subnetID)
	assert.NotEmpty(t, subnetACLID)
	assert.NotEmpty(t, subnetACLRuleID)
	assert.Equal(t, "true", public)
}

func genName() string {
	return fmt.Sprintf("test-subnet-%s", random.UniqueId())
}
