package validate

import (
	"os/exec"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/sys/unix"
)

func PathIsExecutableFile() schema.SchemaValidateDiagFunc {
	return func(v interface{}, path cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics

		absolutePath, err := exec.LookPath(v.(string))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Invalid path",
				Detail:   err.Error(),
			})
		}

		if unix.Access(absolutePath, unix.X_OK) != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Invalid path",
				Detail:   "The specified path is not an executable file.",
			})
		}

		return diags
	}
}
