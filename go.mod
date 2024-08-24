module github.com/johnnewcombe/telstar-util

go 1.22

// use the local library rather than the one in git
replace github.com/johnnewcombe/telstar-library => ../telstar-library

require github.com/spf13/cobra v1.8.1

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
