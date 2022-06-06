resource "azurerm_application_insights" "application_insights" {
  name                = var.app_insights_name
  location            = var.resource_group.location
  resource_group_name = var.resource_group.name
  application_type    = var.app_insights.application_type
  tags                = local.appins_tags
  workspace_id        = azurerm_log_analytics_workspace.log_analytics.id
}

resource "azurerm_log_analytics_workspace" "log_analytics" {
  name                = var.log_analytics_workspace_name
  location            = var.resource_group.location
  resource_group_name = var.resource_group.name
  sku                 = var.log_analytics.sku
  retention_in_days   = var.log_analytics.retention_in_days
  tags                = local.log_tags
}