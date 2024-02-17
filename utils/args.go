package utils

import (
	"flag"
)

// Args Hold all arguments.
type Args struct {
	URL    string
	KEY    string
	IV     string
	Encode bool
}

// Usage Return usage
func Usage() { flag.Usage() }

// ParseFlags Parse all available flags.
func ParseFlags() *Args {
	VarsArgs := &Args{}
	flag.StringVar(&VarsArgs.URL, "url", "", "DATA/URL To Encode/Decode (AES Encrypted)")
	flag.StringVar(&VarsArgs.KEY, "key", "", "AES Encryption Key")
	flag.StringVar(&VarsArgs.IV, "iv", "", "AES Encryption IV")
	flag.BoolVar(&VarsArgs.Encode, "encode", false, "Encode a payload")
	flag.Parse()
	return VarsArgs
}
