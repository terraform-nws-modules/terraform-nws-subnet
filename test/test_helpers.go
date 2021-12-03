package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// validates Terraform output versus expected
func validate(t *testing.T, opts *terraform.Options, domain string) {
	id := terraform.Output(t, opts, "id")
	out_domain := terraform.Output(t, opts, "domain")

	act_domain := trimBrackets(out_domain)

	assert.NotEmpty(t, id)
	assert.Equal(t, domain, act_domain)
}

// Validation helpers
func trimBrackets(s string) string {
	str0 := strings.TrimLeft(s, "[")
	str1 := strings.TrimRight(str0, "]")
	return str1
}

func genName() string {
	return fmt.Sprintf("test-subnet-%s", random.UniqueId())
}
