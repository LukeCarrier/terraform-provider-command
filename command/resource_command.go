package command

import (
	"bytes"
	"context"
	"os/exec"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/LukeCarrier/terraform-provider-command/command/validate"
)

func resourceCommand() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCommandCreate,
		ReadContext:   resourceCommandRead,
		DeleteContext: resourceCommandDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.PathIsExecutableFile(),
				ForceNew:         true,
			},
			"arguments": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				ForceNew: true,
			},
			"stderr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"stdout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCommandCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var err error
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	name := d.Get("name").(string)
	rawArguments := d.Get("arguments").([]interface{})
	arguments := make([]string, len(rawArguments))
	for i, rawArgument := range rawArguments {
		arguments[i] = rawArgument.(string)
	}

	cmd := exec.Command(name, arguments...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to run command",
			Detail:   err.Error(),
		})
	}

	d.SetId(cmd.String())
	d.Set("stdout", stdout.String())
	d.Set("stderr", stderr.String())

	return diags
}

func resourceCommandRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceCommandDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
