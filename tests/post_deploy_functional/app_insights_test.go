package tests

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/gruntwork-io/go-commons/files"
	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type TerraTestSuite struct {
	suite.Suite
	TerraformOptions *terraform.Options
}

// setup to do before any test runs
func (suite *TerraTestSuite) SetupSuite() {
	tmpDir := test_structure.CopyTerraformFolderToTemp(suite.T(), "../..", ".")
	_ = files.CopyFile(path.Join("..", "..", ".tool-versions"), path.Join(tmpDir, ".tool-versions"))
	cwd, err := os.Getwd()
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.TerraformOptions = terraform.WithDefaultRetryableErrors(suite.T(), &terraform.Options{
		TerraformDir: tmpDir,
		VarFiles:     []string{path.Join(cwd, "..", "test.tfvars")},
	})
	terraform.InitAndApplyAndIdempotent(suite.T(), suite.TerraformOptions)
}

// TearDownAllSuite has a TearDownSuite method, which will run after all the tests in the suite have been run.
func (suite *TerraTestSuite) TearDownSuite() {
	terraform.Destroy(suite.T(), suite.TerraformOptions)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TerraTestSuite))
}

func (suite *TerraTestSuite) TestAppInsights() {
	// NOTE: "subscriptionID" is overridden by the environment variable "ARM_SUBSCRIPTION_ID". <>
	subscriptionID := ""

	app_insights_name := terraform.Output(suite.T(), suite.TerraformOptions, "appins_name")
	log_analytics_name := terraform.Output(suite.T(), suite.TerraformOptions, "log_analytics_name")
	rg_name := "deb-test-devops"
	expected_app_insights_name := "demo-eastus-dev-000-appins-000"
	expected_log_analytics_name := "demo-eastus-dev-000-logs-000"

	workspace := azure.GetLogAnalyticsWorkspace(suite.T(), log_analytics_name, rg_name, subscriptionID)
	fmt.Println(*workspace.Name)
	suite.NotEmpty(*workspace, "Workspace cannot be empty")
	suite.Equal(expected_app_insights_name, app_insights_name, "Name should match")
	suite.Equal(expected_log_analytics_name, log_analytics_name, "Name should match")

}
