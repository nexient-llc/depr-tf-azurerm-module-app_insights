package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAzureAI_nameIsCorrect(t *testing.T) {
	t.Parallel()
	// https://github.com/gruntwork-io/terratest/issues/337
	defer os.Remove("../test-provider.tf")
	err := files.CopyFile("../../provider.tf", "../test-provider.tf")
	if err != nil {
		fmt.Printf("There was an error copying the file.")
	}

	planFilePath := filepath.Join("./tests/pre_deploy_functional", "terraform.tfplan")
	//Set variable values to be supplied to terraform module
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../..",
		Vars: map[string]interface{}{
			"resource_group": map[string]interface{}{
				"location": "westus",
				"name":     "iac-dev-000-rg-000",
			},
			"app_insights": map[string]interface{}{
				"application_type": "web",
				"custom_tags": map[string]interface{}{
					"environment": "Development",
					"creator":     "Nexient",
				},
			},
			"app_insights_name": "demo-westus-dev-001-ai-001",
			"log_analytics": map[string]interface{}{
				"sku":                                "PerGB2018",
				"retention_in_days":                  30,
				"daily_quota_gb":                     0.5,
				"custom_tags":                        map[string]interface{}{},
				"internet_ingestion_enabled":         true,
				"internet_query_enabled":             true,
				"reservation_capacity_in_gb_per_day": 100,
			},
			"log_analytics_workspace_name": "demo-westus-dev-001-log-001",
		},
		PlanFilePath: planFilePath,
	})

	plan := terraform.InitAndPlanAndShowWithStruct(t, terraformOptions)
	terraform.RequirePlannedValuesMapKeyExists(t, plan, "azurerm_application_insights.application_insights")
	aiResource := plan.ResourcePlannedValuesMap["azurerm_application_insights.application_insights"]
	aiName := aiResource.AttributeValues["name"].(string)
	assert.Equal(t, "demo-westus-dev-001-ai-001", aiName)
}
