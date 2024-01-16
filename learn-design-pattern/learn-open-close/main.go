package main

import "fmt"

type Color int
type Size int
type Product struct {
	name  string
	color Color
	size  Size
}

const (
	red Color = iota
	green
	blue
)

const (
	small Size = iota
	medium
	large
)

// Sebelum pakai open close
// perlu tambah method kalau ada case case baru yang mana method
// di filter struct makin complex ❌

// type Filter struct{}

// func (f *Filter) filterByColor(products []Product, color Color) []*Product {
// 	result := make([]*Product, 0)

// 	for i, v := range products {
// 		if v.color == color {
// 			result = append(result, &products[i])
// 		}
// 	}

// 	return result
// }

// func (f *Filter) filterBySize(products []Product, size Size) []*Product {
// 	result := make([]*Product, 0)

// 	for i, v := range products {
// 		if v.size == size {
// 			result = append(result, &products[i])
// 		}
// 	}

// 	return result
// }

// func (f *Filter) filterBySizeAndColor(products []Product, color Color, size Size) []*Product {
// 	result := make([]*Product, 0)

// 	for i, v := range products {
// 		if v.color == color && v.size == size {
// 			result = append(result, &products[i])
// 		}
// 	}

// 	return result
// }

// Solusinya buat struct yang agnostic lalu di exteds dengan filter

type Specification interface {
	isMatch(product *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) isMatch(p *Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) isMatch(p *Product) bool {
	return p.size == spec.size
}

type AndSpecification struct {
	firstSpec, secondSpec Specification
}

func (spec AndSpecification) isMatch(p *Product) bool {
	return spec.firstSpec.isMatch(p) && spec.secondSpec.isMatch(p)
}

type OrSpecification struct {
	firstSpec, secondSpec Specification
}

func (spec OrSpecification) isMatch(p *Product) bool {
	return spec.firstSpec.isMatch(p) || spec.secondSpec.isMatch(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if spec.isMatch(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}

func main() {
	apple := Product{name: "Apple", color: green, size: small}
	tree := Product{name: "Tree", color: green, size: large}
	house := Product{name: "House", color: blue, size: large}

	products := []Product{apple, tree, house}
	bf := BetterFilter{}

	fmt.Println("Green products:")
	greenSpec := ColorSpecification{green}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Println("Large products:")
	largeSpec := SizeSpecification{large}
	for _, v := range bf.Filter(products, largeSpec) {
		fmt.Printf(" - %s is large\n", v.name)
	}

	fmt.Println("Green Large products:")
	greenLargeSpec := AndSpecification{largeSpec, greenSpec}
	for _, v := range bf.Filter(products, greenLargeSpec) {
		fmt.Printf(" - %s is large & green\n", v.name)
	}

	fmt.Println("Green or Large products:")
	greenOrLargeSpec := OrSpecification{largeSpec, greenSpec}
	for _, v := range bf.Filter(products, greenOrLargeSpec) {
		fmt.Printf(" - %s is large / green\n", v.name)
	}
}
