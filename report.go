package main

import (
	"fmt"
)

func JSONOutput(t *SSLTest) {
	fmt.Printf("%s\n", t.ToJSON())
}

// generate a human readable report in stdout
func HumanOutput(t *SSLTest) {
	fmt.Println("========== SSL Report ==========")

	fmt.Printf("Hostname: %s\n\n", t.Addr)

	fmt.Println("TLS Versions\n------------")
	for v, s := range t.Versions {
		fmt.Printf("%-6s: %t\n", v, s)
	}

	fmt.Println("\nTLS Ciphers\n-----------")
	for c, s := range t.Ciphers {
		fmt.Printf("%-42s: %t\n", c, s)
	}

	fmt.Printf("\nStarted At : %s\n", t.StartedAt.Format("2006/01/02 15:04:05 MST"))
	fmt.Printf("Finished At: %s\n", t.FinishedAt.Format("2006/01/02 15:04:05 MST"))

	fmt.Println("================================")
}
