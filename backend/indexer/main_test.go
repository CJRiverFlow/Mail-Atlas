/*
	Indexer test, verify the following before
	* mail file has been unziped in the root repor directory
	* build and start the Zincsearch database container
	* Graphviz installed https://www.graphviz.org/download/

	#PROFILING
	Run in terminal:
	- go test -cpuprofile ./evaluation/cpu-1.prof -memprofile ./evaluation/mem-1.prof -bench .
	- go tool pprof --web ./evaluation/cpu-1.prof
	- go tool pprof --svg ./evaluation/cpu-1.prof > cpu-1.svg
	- go tool pprof --top cpu-1.prof
*/
package main

import (
	"os"
	"testing"
)

func TestIndexer(t *testing.T) {
	testDir := "../../enron_mail_20110402/maildir/gang-l"
	os.Args = []string{"", "--path", testDir}
	main()
}
