// Package
package main

// Imports
import (
	"fmt"
	. "glu-software/internal"
	"glu-software/routes"
	"net/http"
	"os"
	"strconv"
)

// Data
var addr string

// Early
func init() {

	// Parse command line arguments
	parseCommandLineArgs()

	// Print a success message to standard out
	fmt.Printf("http://%s\n", addr)
}

// Entry point
func main() {

	// Serve
	http.ListenAndServe(addr, routes.Router)
}

func parseCommandLineArgs() {

	// Error check
	if len(os.Args) != 2 {

		// Error
		printUsage(1)
	}

	// Uninitialized data
	var port_no int
	var hostname string

	// Initialized data
	var err error = nil

	// Store the hostname
	hostname, err = os.Hostname()

	// Error checking
	CheckErr(err)

	// Store the port number
	port_no, err = strconv.Atoi(os.Args[1])

	// Error checking
	CheckErr(err)

	// Bounds check
	if port_no < 1 || port_no > 65535 {

		// Error
		printUsage(1)
	}

	// Construct the address string
	addr = fmt.Sprintf("%s:%d", hostname, port_no)

	// Done
	return
}

func printUsage(exit int) {

	// Print a usage message
	fmt.Printf("Usage: glu_app <port_no>\n")

	// Abort
	os.Exit(exit)
}
