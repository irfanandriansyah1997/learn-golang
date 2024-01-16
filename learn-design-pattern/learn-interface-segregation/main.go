package main

import (
	"fmt"
)

type Document struct{}

type Machine interface {
	Print(d Document)
	Scan(d Document)
}

// Issue nya kita punya interface tapi di beberapa implementasi
// beberapa method ngga bisa digunain
// dan cuma bisa diisi panic

// type ModernPrinter struct{}

// func (s ModernPrinter) Print(d Document) {
// 	fmt.Println(s)
// }

// func (s ModernPrinter) Scan(d Document) {
// 	fmt.Println(s)
// }

// type OldPrinter struct{}

// func (s OldPrinter) Print(d Document) {
// 	fmt.Println(s)
// }

// func (s OldPrinter) Scan(d Document) {
// 	panic("ngga bisa digunain ") // Bakal break kalau misal dipake langsung ❌
// }

// Cara Resolve

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	if m.printer != nil {
		m.printer.Print(d)
		return
	}

	fmt.Println("Not provide Print method")
}

func (m MultiFunctionMachine) Scan(d Document) {
	if m.scanner != nil {
		m.scanner.Scan(d)
		return
	}

	fmt.Printf("Not provide Scan method\n")
}

// Cara Implementasi

type EpsonPrinter struct{}

func (e EpsonPrinter) Print(d Document) {
	fmt.Println("Epson", d)
}

type CanonPrinter struct{}

func (c CanonPrinter) Print(d Document) {
	fmt.Println("Canon", d)
}

func (c CanonPrinter) Scan(d Document) {
	fmt.Println("Canon", d)
}

func main() {
	epson := EpsonPrinter{}
	canon := CanonPrinter{}

	device1 := MultiFunctionMachine{
		printer: epson,
		scanner: nil,
	}
	device2 := MultiFunctionMachine{
		printer: canon,
		scanner: canon,
	}

	device1.Print(Document{})
	device1.Scan(Document{})

	device2.Print(Document{})
	device2.Scan(Document{})
}
