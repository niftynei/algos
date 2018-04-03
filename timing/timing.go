package timing

import (
	"fmt"
	"time"
)

func Timer(start time.Time, desc string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", desc, elapsed)
}
