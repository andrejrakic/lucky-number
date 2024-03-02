package main

import (
	// Import the random package using the new location under the
	// `internal` directory.
	"github.com/andrejrakic/lucky-number/internal/random"

	// Import the color package.
	"github.com/fatih/color"
)

func main() {
	// Call the random.Number() function to get the random number. Notice that
	// we use the package name as the accessor, just like we do for the standard
	// library packages.
	// Use it to print the message in green.
	green := color.New(color.FgGreen)
	green.Printf("Your lucky number is %d!\n", random.Number())
}
