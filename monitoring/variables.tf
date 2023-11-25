variable "alert_on_panics" {
  type        = bool
  default     = true
  description = "Enables/Disables the panics monitor defined based on the logs"
}

variable "cpu_high_threshold" {
  type    = number
  default = 80
}

# window in minutes
variable "cpu_high_window" {
  type    = number
  default = 30
}

# Number of restarts per 15m to trigger a P2 alert.
variable "pod_restart_low_count_threshold" {
  type    = number
  default = 0
}

# Number of restarts per 30m to be considered a P1 incident.
variable "pod_restart_high_count_threshold" {
  type    = number
  default = 3
}

variable "exclude_environments" {
  type    = string
  default = "development"
}
