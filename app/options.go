package app

import "dtweave.io/goframe/app/flag"

type Options interface {
	// Flags returns flags for sub gateway
	Flags() (fss flag.NamedFlagSets)
	// Validate validate params
	Validate() []error
	// Complete set default value
	Complete() error
}
