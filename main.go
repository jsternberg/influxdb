package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jsternberg/influxdb/v2"
)

func main() {
	if err := influxdb.Run(context.TODO()); err != nil {
		fmt.Fprintf(os.Stderr, "fatal error: %s.\n", err)
		os.Exit(1)
	}
}
