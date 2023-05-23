package utils

import (
	"flag"

	"k8s.io/klog/v2"
)

// SetupKlog configures the klog logger with the specified level
func SetupKlog(level string) {
	flags := flag.NewFlagSet("klog", flag.ContinueOnError)
	klog.InitFlags(flags)
	_ = flags.Set("v", level)
	_ = flags.Parse(nil)
}
