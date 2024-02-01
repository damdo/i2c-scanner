// Code heavily inspired by: https://pkg.go.dev/periph.io/x/periph/conn/i2c
package main

import (
	"fmt"
	"log"
	"strings"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

func main() {
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Enumerate all I2C buses available and the corresponding pins.
	fmt.Print("I2C buses available:\n")
	for _, ref := range i2creg.All() {
		fmt.Printf("- name: %s\n", ref.Name)

		if ref.Number != -1 {
			fmt.Printf("  number: %d\n", ref.Number)
		}

		if len(ref.Aliases) != 0 {
			fmt.Printf("  aliases: %s\n", strings.Join(ref.Aliases, " "))
		}

		b, err := ref.Open()
		if err != nil {
			fmt.Printf("  Failed to open bus: %v", err)
		}

		fmt.Println("  pins:")
		if p, ok := b.(i2c.Pins); ok {
			fmt.Printf("   * SDA: %s\n", p.SDA())
			fmt.Printf("   * SCL: %s\n", p.SCL())
		}

		if err := b.Close(); err != nil {
			fmt.Printf("  Failed to close bus: %v", err)
		}

		fmt.Println()
	}
}
