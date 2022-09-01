package mailman3

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceDomain() *schema.Resource {
	return &schema.Resource{
		Description: "This data source can be used to retrieve information about a domain.",
		ReadContext: resourceDomainRead,
		Schema:      schemaDomain(),
	}
}
