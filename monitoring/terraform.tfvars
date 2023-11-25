# Fill these in with DataDog users/integrations to notify
// <<Stencil::Block(tfNotificationPriorities)>>
P1_notify = []
P2_notify = []
// <</Stencil::Block>>

# Fill these in with tags for your datadog dashboards/monitors
# Team and service names will be added automatically elsewhere, add anything additional to those two in here
// <<Stencil::Block(tfAdditionalDdTags)>>
additional_dd_tags = []
// <</Stencil::Block>>

# Replace the following values with adequate yellow/red
# thresholds for your service call latencies
#
# Note that threshold affects presentation of Performance charts
# and not used for monitors/alerts.
// <<Stencil::Block(tfLatencyThresholdsMs)>>
Latency_red_line_ms    = 500
Latency_yellow_line_ms = 200
// <</Stencil::Block>>

# Replace the following values with adequate yellow/red
# thresholds (in percentage) for your service call latencies
#
# Note that threshold affect presentation of QoS charts
# and not used for monitors/alerts
// <<Stencil::Block(tfLatencyThresholdsPercentage)>>
Qos_red_line    = 98
Qos_yellow_line = 99
// <</Stencil::Block>>

# Replace the following values with environments
# for your service you want to ignore alerts for.
// <<Stencil::Block(tfExcludeEnvironments)>>
exclude_environments = "development"
// <</Stencil::Block>>

# You can override your service's reportingTeam name in case your DataDog team tag is different by:
# dd_team_tag = "xxx-xxxx"
// <<Stencil::Block(tfCustomVars)>>

// <</Stencil::Block>>
