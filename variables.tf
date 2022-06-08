#########################################
#Common variables
#########################################

variable "resource_group" {
  description = "target resource group resource mask"
  type = object({
    name     = string
    location = string
  })
}

#########################################
#Variables associated with app insights
#########################################

variable "app_insights" {
  description = "App insights primitive."
  type = object({
    application_type = string
    custom_tags   = map(string)
  })
  default = {
    application_type = "web"
    custom_tags   = {}
  }
}

variable "app_insights_name" {
  type        = string
  description = "App insights name."
}

#########################################
#Variables associated with log analytics
#########################################

variable "log_analytics" {
  description = "Log analytics primitive."
  type = object({
    sku                                = string
    retention_in_days                  = number
    daily_quota_gb                     = number
    custom_tags                    = map(string)
    internet_ingestion_enabled         = bool
    internet_query_enabled             = bool
    reservation_capacity_in_gb_per_day = number
  })
  default = {
    sku                                = "PerGB2018"
    retention_in_days                  = 30
    daily_quota_gb                     = 0.5
    custom_tags                     = {}
    internet_ingestion_enabled         = true
    internet_query_enabled             = true
    reservation_capacity_in_gb_per_day = 100
  }
}

variable "log_analytics_workspace_name" {
  type        = string
  description = "Log analytics workspace name."
}