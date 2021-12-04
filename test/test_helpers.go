package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

type TestCaseT struct {
	name          []string
	cidr          []string
	domain        string
	public        bool
	acl_port_list []string
}

// validates Terraform output versus expected
func validatePrivate(t *testing.T, opts *terraform.Options) {
	subnet_id := terraform.Output(t, opts, "subnet_id")
	subnet_acl_id := terraform.Output(t, opts, "subnet_acl_id")
	subnet_acl_rule_id := terraform.Output(t, opts, "subnet_acl_rule_id")
	public := terraform.Output(t, opts, "subnet_public")

	assert.NotEmpty(t, subnet_id)
	assert.Equal(t, "[]", subnet_acl_id)
	assert.Equal(t, "[]", subnet_acl_rule_id)
	assert.Equal(t, "false", public)
}

func validatePublic(t *testing.T, opts *terraform.Options) {
	subnet_id := terraform.Output(t, opts, "subnet_id")
	subnet_acl_id := terraform.Output(t, opts, "subnet_acl_id")
	subnet_acl_rule_id := terraform.Output(t, opts, "subnet_acl_rule_id")
	public := terraform.Output(t, opts, "subnet_public")

	assert.NotEmpty(t, subnet_id)
	assert.NotEmpty(t, subnet_acl_id)
	assert.NotEmpty(t, subnet_acl_rule_id)
	assert.Equal(t, "true", public)
}

func genName() string {
	return fmt.Sprintf("test-subnet-%s", random.UniqueId())
}
