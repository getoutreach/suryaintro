
resource "datadog_service_level_objective" "grpc_success" {
  name        = "Suryaintro GRPC Success Response"
  type        = "metric"
  description = "Comparing (status:ok) responses to all requests as a ratio, broken out by bento."
  tags        = local.ddTags
  query {
    numerator   = "clamp_min(default_zero(count:${local.grpc_request_source}{${join(", ", var.grpc_tags)},app:suryaintro, !statuscategory:categoryservererror} by {kube_namespace}.as_count()), 1)"
    denominator = "clamp_min(default_zero(count:${local.grpc_request_source}{${join(", ", var.grpc_tags)},app:suryaintro} by {kube_namespace}.as_count()), 1)"
  }
  thresholds {
    timeframe = "7d"
    target    = 99.9
    warning   = 99.95
  }
}

// <<Stencil::Block(tfCustomSLODatadog)>>

// <</Stencil::Block>>
