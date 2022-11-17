package gover

import "runtime"

func version() string {
	return runtime.Version()
}