variable "grpc_request_source" {
  type    = string
  default = "grpc_request_handled"
}

variable "grpc_tags" {
  type    = list(string)
  default = ["*", "!env:development"]
}

variable "grpc_evaluation_window" {
  type    = string
  default = "last_15m"
}

variable "grpc_low_count_threshold" {
  type    = number
  default = 1000
}

variable "grpc_qos_low_traffic_threshold" {
  type    = number
  default = 50
}

variable "grpc_qos_high_traffic_threshold" {
  type    = number
  default = 99
}

locals {
  grpc_request_source = local.prefix ? "suryaintro.${var.grpc_request_source}" : "${var.grpc_request_source}"
}

module "grpc_success_rate_low" {
  source              = "git@github.com:getoutreach/monitoring-terraform.git//modules/alerts/datadog/low-traffic-composite-monitor"
  name                = "Suryaintro GRPC Success Rate Low"
  tags                = local.ddTags
  message             = <<EOF
  Composite monitor of GRPC QoS based on traffic
  Calculating the success rate (!statuscategory:categoryservererror) of GRPC requests as a 0-100% monitor.
  High traffic -> ${var.grpc_qos_high_traffic_threshold}%
  Low traffic -> ${var.grpc_qos_low_traffic_threshold}%
  Runbook: "https://github.com/getoutreach/suryaintro/blob/main/documentation/runbooks/grpc-success-rate-low.md"
  Notify: ${join(" ", var.P2_notify)}
  EOF
  require_full_window = false
  low_count_query     = "sum(${var.grpc_evaluation_window}):clamp_min(default_zero(count:${local.grpc_request_source}{${join(", ", var.grpc_tags)},app:suryaintro} by {kube_namespace}.as_count()), 1) < ${var.grpc_low_count_threshold}"
  low_traffic_query   = "sum(${var.grpc_evaluation_window}):100 * clamp_min(default_zero(count:${local.grpc_request_source}{${join(", ", var.grpc_tags)},app:suryaintro,statuscategory:categoryservererror} by {kube_namespace}.as_count()), 1) / clamp_min(default_zero(count:${local.grpc_request_source}{${join(", ", var.grpc_tags)},app:suryaintro} by {kube_namespace}.as_count()), 1) >= ${var.grpc_qos_low_traffic_threshold}"
  high_traffic_query  = "sum(${var.grpc_evaluation_window}):100 * clamp_min(default_zero(count:${local.grpc_request_source}{${join(", ", var.grpc_tags)},app:suryaintro, !statuscategory:categoryservererror} by {kube_namespace}.as_count()), 1) / clamp_min(default_zero(count:${local.grpc_request_source}{${join(", ", var.grpc_tags)},app:suryaintro} by {kube_namespace}.as_count()), 1) < ${var.grpc_qos_high_traffic_threshold}"
}

variable "grpc_latency_low_traffic_percentile" {
  type    = number
  default = 90
}

variable "grpc_latency_high_traffic_percentile" {
  type    = number
  default = 99
}

variable "grpc_latency_low_traffic_threshold" {
  type    = number
  default = 2
}

variable "grpc_latency_high_traffic_threshold" {
  type    = number
  default = 2
}

module "grpc_latency_high" {
  source              = "git@github.com:getoutreach/monitoring-terraform.git//modules/alerts/datadog/low-traffic-composite-monitor"
  name                = "Suryaintro GRPC Latency High"
  tags                = local.ddTags
  message             = <<EOF
  Composite monitor of GRPC request latency based on traffic
  High traffic -> P${var.grpc_latency_high_traffic_percentile}
  Low traffic -> P${var.grpc_latency_low_traffic_percentile}
  Runbook: "https://github.com/getoutreach/suryaintro/blob/main/documentation/runbooks/grpc-latency-high.md"
  Notify: ${join(" ", var.P2_notify)}
  EOF
  require_full_window = false
  low_count_query     = "sum(${var.grpc_evaluation_window}):clamp_min(default_zero(count:${local.grpc_request_source}{${join(", ", var.grpc_tags)},app:suryaintro} by {kube_namespace}.as_count()), 1) < ${var.grpc_low_count_threshold}"
  low_traffic_query   = "avg(${var.grpc_evaluation_window}):default_zero(p${var.grpc_latency_low_traffic_percentile}:${local.grpc_request_source}{${join(", ", var.grpc_tags)},app:suryaintro} by {kube_namespace}) > ${var.grpc_latency_low_traffic_threshold}"
  high_traffic_query  = "avg(${var.grpc_evaluation_window}):default_zero(p${var.grpc_latency_high_traffic_percentile}:${local.grpc_request_source}{${join(", ", var.grpc_tags)},app:suryaintro} by {kube_namespace}) > ${var.grpc_latency_high_traffic_threshold}"
}

// <<Stencil::Block(tfCustomGRPCDatadog)>>

// <</Stencil::Block>>
