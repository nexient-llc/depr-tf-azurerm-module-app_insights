locals {
    default_tags = {
        provisioner = "terraform"
        
    }

    appins_tags = merge({
        name        = var.app_insights_name
    }, local.default_tags, var.app_insights.custom_tags)

    log_tags = merge({
        name        = var.log_analytics_workspace_name
    }, local.default_tags, var.log_analytics.custom_tags)
}