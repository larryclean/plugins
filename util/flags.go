package util

import (
	goflag "flag"
	"github.com/spf13/pflag"
	"strings"
)

// WordSepNormalizeFunc changes all flags that contain "_" separators
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
	}
	return pflag.NormalizedName(name)
}

// WarnWordSepNormalizeFunc changes and warns for flags that contain "_" separators
func WarnWordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		nname := strings.Replace(name, "_", "-", -1)
		//glog.Warningf("%s is DEPRECATED and will be removed in a future version. Use %s instead.", name, nname)

		return pflag.NormalizedName(nname)
	}
	return pflag.NormalizedName(name)
}

// InitFlags normalizes and parses the command line flags
func InitFlags() {
	pflag.CommandLine.SetNormalizeFunc(WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.Parse()
}
