package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareAccessPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The ID of the application the policy is associated with.",
		},
		consts.AccountIDSchemaKey: {
			Description:   "The account identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{consts.ZoneIDSchemaKey},
		},
		consts.ZoneIDSchemaKey: {
			Description:   "The zone identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{consts.AccountIDSchemaKey},
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Friendly name of the Access Policy.",
		},
		"precedence": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "The unique precedence for policies on a single application.",
		},
		"decision": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"allow", "deny", "non_identity", "bypass"}, false),
			Description:  fmt.Sprintf("Defines the action Access will take if the policy matches the user. %s", renderAvailableDocumentationValuesStringSlice([]string{"allow", "deny", "non_identity", "bypass"})),
		},
		"require": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        AccessGroupOptionSchemaElement,
			Description: "A series of access conditions, see [Access Groups](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/resources/access_group#conditions).",
		},
		"exclude": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        AccessGroupOptionSchemaElement,
			Description: "A series of access conditions, see [Access Groups](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/resources/access_group#conditions).",
		},
		"include": {
			Type:        schema.TypeList,
			Required:    true,
			Elem:        AccessGroupOptionSchemaElement,
			Description: "A series of access conditions, see [Access Groups](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/resources/access_group#conditions).",
		},
		"purpose_justification_required": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to prompt the user for a justification for accessing the resource.",
		},
		"purpose_justification_prompt": {
			Type:         schema.TypeString,
			Optional:     true,
			RequiredWith: []string{"purpose_justification_required"},
			Description:  "The prompt to display to the user for a justification for accessing the resource.",
		},
		"approval_required": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"approval_group": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     AccessPolicyApprovalGroupElement,
		},
	}
}

var AccessPolicyApprovalGroupElement = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"email_list_uuid": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"email_addresses": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "List of emails to request approval from.",
		},
		"approvals_needed": {
			Type:         schema.TypeInt,
			Required:     true,
			ValidateFunc: validation.IntAtLeast(0),
			Description:  "Number of approvals needed.",
		},
	},
}
