variable "P1_notify" {
  type    = list(string)
  default = []
}
variable "P2_notify" {
  type    = list(string)
  default = []
}
variable "additional_dd_tags" {
  type    = list(string)
  default = []
}
variable "dd_team_tag" {
  type    = string
  default = "cad-fri"
}

locals {
  ddTags = concat(["suryaintro", "service:suryaintro", "team:${var.dd_team_tag}"], var.additional_dd_tags)
}

resource "datadog_monitor" "argocd_application_health_status" {
  type                = "query alert"
  name                = "Suryaintro argocd application health status"
  query               = "max(last_15m):default_zero(clamp_max(sum:application_controller.argocd_app_info{name:suryaintro,health_status:healthy} by {cluster_name}.fill(zero, 3), 1)) < 1"
  tags                = local.ddTags
  message             = <<EOF
  ArgoCD Health status has been unhealthy over the window last_15m.
  Note: This monitor will auto-resolve after a Healthy status is reported within the specified evaluation timeframe
  Runbook: "https://outreach-io.atlassian.net/wiki/spaces/DT/pages/2390589626/ArgoCD+Runbooks"
  EOF
  require_full_window = false
}

resource "datadog_monitor" "argocd_application_sync_status" {
  type                = "query alert"
  name                = "Suryaintro argocd application sync status"
  query               = "max(last_15m):default_zero(clamp_max(sum:application_controller.argocd_app_info{name:suryaintro,sync_status:synced} by {cluster_name}.fill(zero, 3), 1)) < 1"
  tags                = local.ddTags
  message             = <<EOF
  ArgoCD Sync status has not been synced over the window last_15m.
  Note: This monitor will auto-resolve after a Synced status is reported within the specified evaluation timeframe
  Runbook: "https://outreach-io.atlassian.net/wiki/spaces/DT/pages/2390589626/ArgoCD+Runbooks"
  EOF
  require_full_window = false
}

# splitting the interval 15 mins to 3 windows (moving rollup by 5mins) and if each of them contains restart -> alert
resource "datadog_monitor" "pod_restarts" {
  type                = "query alert"
  name                = "Suryaintro Pod Restarts > ${var.pod_restart_low_count_threshold} last 15m"
  query               = "min(last_15m):moving_rollup(diff(sum:kubernetes_state.container.restarts{kube_container_name:suryaintro,!env:development} by {kube_namespace}), 300, 'sum') > ${var.pod_restart_low_count_threshold}"
  tags                = local.ddTags
  message             = <<EOF
  If we ever have a pod restart, we want to know.
  Note: This monitor will auto-resolve after 15 minutes of no restarts.
  Runbook: "https://github.com/getoutreach/suryaintro/blob/main/documentation/runbooks/pod-restarts.md"
  Notify: ${join(" ", var.P2_notify)}
  EOF
  require_full_window = false
}

resource "datadog_monitor" "pod_restarts_high" {
  type                = "query alert"
  name                = "Suryaintro Pod Restarts > ${var.pod_restart_high_count_threshold} last 30m"
  query               = "max(last_30m):diff(sum:kubernetes_state.container.restarts{kube_container_name:suryaintro,!env:development} by {kube_namespace}) > ${var.pod_restart_high_count_threshold}"
  tags                = local.ddTags
  message             = <<EOF
  Several pods are being restarted.
  Note: This monitor will auto-resolve after 30 minutes of no restarts.
  Runbook: "https://github.com/getoutreach/suryaintro/blob/main/documentation/runbooks/pod-restarts.md"
  Notify: ${join(" ", var.P1_notify)}
  EOF
  require_full_window = false
}

# default to 0 if the pod was running on high CPU and then stopped/killed
resource "datadog_monitor" "pod_cpu_high" {
  type                = "query alert"
  name                = "Suryaintro Pod CPU > ${var.cpu_high_threshold}% of limit last ${var.cpu_high_window}m"
  query               = "avg(last_${var.cpu_high_window}m):100 * (default_zero(avg:kubernetes.cpu.usage.total{app:suryaintro,!env:development} by {kube_namespace,pod_name}) / 1000000000) / avg:kubernetes.cpu.limits{app:suryaintro,!env:development} by {kube_namespace,pod_name} >= ${var.cpu_high_threshold}"
  tags                = local.ddTags
  message             = <<EOF
  One of the service's pods has been using over ${var.cpu_high_threshold}% of its limit CPU on average for the last ${var.cpu_high_window} minutes.  This almost certainly means that the service needs more CPU to function properly and is being throttled in its current form.
  Runbook: "https://github.com/getoutreach/suryaintro/blob/main/documentation/runbooks/pod-cpu.md"
  Notify: ${join(" ", var.P2_notify)}
  EOF
  require_full_window = false
}

resource "datadog_monitor" "panics" {
  type    = "log alert"
  name    = "Suryaintro Service panics"
  query   = "logs(\"panic status:error kube_namespace:suryaintro--* -env:development\").index(\"*\").rollup(\"count\").by(\"kube_namespace,service\").last(\"5m\") > 0"
  tags    = local.ddTags
  message = <<EOF
  Log based monitor of runtime error panics.
  Note: This P1 alert only includes production
  Runbook: "https://github.com/getoutreach/suryaintro/blob/main/documentation/runbooks/service-panics.md"
  Notify: ${join(" ", var.P1_notify)}
  EOF
}

resource "datadog_downtime" "panics_silence" {
  scope      = ["*"]
  monitor_id = datadog_monitor.panics.id
  count      = var.alert_on_panics ? 0 : 1
}

// <<Stencil::Block(tfCustomDatadog)>>

// <</Stencil::Block>>
