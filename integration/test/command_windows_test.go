package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	terraformOptions := makeTerraformOptions(t, "../fixtures/command_hello_windows")
	terraformOptions.Vars = map[string]interface{}{
		"name": "Alice",
	}
	defer terraform.Destroy(t, terraformOptions)

	terraform.Apply(t, terraformOptions)
	output := terraform.Output(t, terraformOptions, "greeting")
	assert.Contains(t, output, "Hello, Alice")
}
