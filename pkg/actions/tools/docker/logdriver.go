package docker

import (
	"github.com/carapace-sh/carapace"
)

// ActionLogDrivers completes log drivers
//
//	splunk (Writes log messages to splunk using the HTTP Event Collector.)
//	syslog (Writes logging messages to the syslog facility. The syslog daemon must be run...)
func ActionLogDrivers() carapace.Action {
	return carapace.ActionValuesDescribed(
		"none", "No logs are available for the container and docker logs does not return any output.",
		"local", "Logs are stored in a custom format designed for minimal overhead.",
		"json-file", "The logs are formatted as JSON. The default logging driver for Docker.",
		"syslog", "Writes logging messages to the syslog facility. The syslog daemon must be running on the host machine.",
		"journald", "Writes log messages to journald. The journald daemon must be running on the host machine.",
		"gelf", "Writes log messages to a Graylog Extended Log Format (GELF) endpoint such as Graylog or Logstash.",
		"fluentd", "Writes log messages to fluentd (forward input). The fluentd daemon must be running on the host machine.",
		"awslogs", "Writes log messages to Amazon CloudWatch Logs.",
		"splunk", "Writes log messages to splunk using the HTTP Event Collector.",
		"etwlogs", "Writes log messages as Event Tracing for Windows (ETW) events. Only available on Windows platforms.",
		"gcplogs", "Writes log messages to Google Cloud Platform (GCP) Logging.",
		"logentries", "Writes log messages to Rapid7 Logentries.",
	)
}
