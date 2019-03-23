package config

import (
	"github.com/acoshift/configfile"
)

var cfg = configfile.NewEnvReader()

// Load loads directory or yaml
func Load(filename string) {
	cfg = configfile.NewReader(filename)
}

// Alias functions
var (
	BytesDefault = cfg.BoolDefault
	Bytes = cfg.Bytes
	MustBytes = cfg.MustBytes
	StringDefault = cfg.StringDefault
	String = cfg.String
	MustString = cfg.MustString
	IntDefault = cfg.IntDefault
	Int = cfg.Int
	MustInt = cfg.MustInt
	Int64Default = cfg.Int64Default
	Int64 = cfg.Int64
	MustInt64 = cfg.MustInt64
	BoolDefault = cfg.BoolDefault
	Bool = cfg.Bool
	MustBool = cfg.MustBool
	DurationDefault = cfg.DurationDefault
	Duration = cfg.Duration
	MustDuration = cfg.MustDuration
)
