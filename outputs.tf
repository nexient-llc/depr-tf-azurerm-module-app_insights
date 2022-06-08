output "log_daily_quota" {
  value       = azurerm_log_analytics_workspace.log_analytics.daily_quota_gb
  description = "Logs Analytics Daily Quota (in GB)"
}

output "log_sku" {
  value       = azurerm_log_analytics_workspace.log_analytics.sku
  description = "Logs Analytics SKU"
}

output "log_analytics_name" {
  value       = azurerm_log_analytics_workspace.log_analytics.name
  description = "Logs Analytics SKU"
}

output "appins_name" {
  value       = azurerm_application_insights.application_insights.name
  description = "Appliction Insights Name"
}

output "app_insights_id" {
  value       = azurerm_application_insights.application_insights.id
  description = "ID of the Application Insights"
}

output "connection_string" {
  value       = azurerm_application_insights.application_insights.connection_string
  description = "Connection String of Application Insights"
  sensitive   = true
}

output "instrumentation_key" {
  value       = azurerm_application_insights.application_insights.instrumentation_key
  description = "Application Insights Instrumentation key"
  sensitive   = true
}

output "rg_name" {
  value       = azurerm_application_insights.application_insights.resource_group_name
  description = "Name of the Resource Group"
}

output "workspace_id" {
  value       = azurerm_application_insights.application_insights.workspace_id
  description = "Log Analytics Workspace ID"
}