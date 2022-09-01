package mailman3

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	gomailman "github.com/mac21/go-mailman"
)

func schemaDomain() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Properties
		"mail_host": &schema.Schema{
			Description: "The mail host for this domain.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"alias_domain": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
	}
}

func resourceDomain() *schema.Resource {
	return &schema.Resource{
		Description:   "This resource can be used to manage mailman domains.",
		CreateContext: resourceDomainCreate,
		ReadContext:   resourceDomainRead,
		DeleteContext: resourceDomainDelete,
		Schema:        schemaDomain(),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func mailmanDomainFromSchemaDomain(d *schema.ResourceData) gomailman.Domain {
	domain := gomailman.Domain{
		MailHost: d.Get("mail_host").(string),
	}

	if description, ok := d.GetOk("description"); ok {
		domain.Description = description.(string)
	}

	if aliasDomain, ok := d.GetOk("alias_domain"); ok {
		domain.AliasDomain = aliasDomain.(string)
	}

	return domain
}

func flattenMailmanDomain(domain *gomailman.Domain) interface{} {
	return map[string]interface{}{
		"mail_host":    domain.MailHost,
		"description":  domain.Description,
		"alias_domain": domain.AliasDomain,
	}
}

func resourceDomainCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cli := m.(*gomailman.Client)
	var diags diag.Diagnostics

	domain := mailmanDomainFromSchemaDomain(d)

	err := cli.AddDomain(&domain)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(domain.MailHost)

	diags = resourceDomainRead(ctx, d, m)

	return diags
}

func resourceDomainRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cli := m.(*gomailman.Client)
	var diags diag.Diagnostics

	mailHost, ok := d.GetOk("mail_host")
	if ok {
		d.SetId(mailHost.(string))
	}

	domain, err := cli.GetDomain(mailHost.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	for key, value := range flattenMailmanDomain(domain).(map[string]interface{}) {
		err := d.Set(key, value)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceDomainDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cli := m.(*gomailman.Client)
	var diags diag.Diagnostics

	mailHost := d.Id()

	err := cli.DeleteDomain(mailHost)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
