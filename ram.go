// ram.go
package main

// #include <unistd.h>
// #include <sys/sysctl.h>
// #include <sys/types.h>
import "C"

import (
	"fmt"

	"git.thwap.org/rockhopper/gout"
)

func RamUsage() {
	maxRam := int64(C.sysconf(C._SC_PHYS_PAGES) * C.sysconf(C._SC_PAGE_SIZE))
	freeRam := int64(C.sysconf(C._SC_AVPHYS_PAGES) * C.sysconf(C._SC_PAGE_SIZE))
	usedRam := (maxRam - freeRam)
	ramPercUsed := (float64(usedRam) / float64(maxRam)) * 100.0
	fmt.Printf("%s: %s\n", gout.Bold(gout.White("RAM")), progress(int(ramPercUsed)))
}