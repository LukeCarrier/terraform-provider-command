package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func makeTerraformOptions(t *testing.T, terraformDir string) *terraform.Options {
	return terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: terraformDir,
		EnvVars: map[string]string{
			"TF_CLI_CONFIG_FILE": "../.terraformrc",
			"TF_LOG":             "TRACE",
		},
	})
}
