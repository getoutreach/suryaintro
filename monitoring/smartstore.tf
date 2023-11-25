module "smartstore_db_dashboard" {
  source               = "git@github.com:getoutreach/database-monitoring.git//modules/smartstore/database?ref=5a2ca62"
  outreach_application = "suryaintro"
  P1_notify            = var.P1_notify
  P2_notify            = var.P2_notify
  additional_dd_tags   = local.ddTags
}


module "smartstore_dirtyschema_monitor" {
  source               = "git@github.com:getoutreach/database-monitoring.git//modules/smartstore/migrations?ref=5a2ca62"
  outreach_application = "suryaintro"
  P1_notify            = var.P1_notify
  P2_notify            = var.P2_notify
  additional_dd_tags   = local.ddTags
}
