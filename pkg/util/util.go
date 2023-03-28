package util

import "strings"

const (
	maxLogLength = 3900
)

func FilterLogs(logs string) string {
	startIndex := strings.Index(logs, "Running on")
	if startIndex != -1 {
		return logs[startIndex:]
	}
	return logs
}

func TrimLogs(logs string) string {
	logs = strings.TrimSpace(logs)
	if len(logs) > maxLogLength {
		return string(logs[len(logs)-maxLogLength])
	}
	return logs
}
