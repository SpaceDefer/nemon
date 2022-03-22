package bigBrother

import (
	"fmt"
	"time"
)

const heartbeatInterval = 10 * time.Second

type ApplicationInfo struct {
	Name     string
	Location string
}

func checkError(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}
