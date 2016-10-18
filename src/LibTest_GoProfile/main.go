package main

import (
	"github.com/pkg/profile"
)

func main() {
	mode := "mem" //flag.String("profile.mode", "", "enable profiling mode, one of [cpu, mem, block, trace]")

	switch mode {
	case "cpu":
		defer profile.Start(profile.CPUProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	case "mem":
		defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	case "block":
		defer profile.Start(profile.BlockProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	case "trace":
		defer profile.Start(profile.TraceProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	default:
		// do nothing
	}
}
