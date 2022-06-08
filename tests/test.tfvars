resource_group = {
    name = "deb-test-devops"
    location = "eastus"
}

log_analytics = {
    sku                                = "PerGB2018"
    retention_in_days                  = 30
    daily_quota_gb                     = 0.5
    custom_tags                     = {}
    internet_ingestion_enabled         = true
    internet_query_enabled             = true
    reservation_capacity_in_gb_per_day = 100
}

app_insights = {
    application_type = "web"
    custom_tags   = {}
}

app_insights_name = "demo-eastus-dev-000-appins-000"

log_analytics_workspace_name = "demo-eastus-dev-000-logs-000"
