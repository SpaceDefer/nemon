package bigBrother

import "time"

const heartbeatInterval = 10 * time.Second

type ApplicationInfo struct {
	Name     string
	Location string
}
