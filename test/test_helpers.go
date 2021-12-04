package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// validates Terraform output versus expected
func validate(t *testing.T, opts *terraform.Options, domain []string) {
	id := terraform.Output(t, opts, "id")

	fmt.Println(">>>>> Output IDs: ", id)

	assert.NotEmpty(t, id)
}

func genName() string {
	return fmt.Sprintf("test-subnet-%s", random.UniqueId())
}
