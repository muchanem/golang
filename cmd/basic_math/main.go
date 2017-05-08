package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	// You can assiagn variables of the same type in this way
	int1, int2, int3 := 13, 14, 15
	// Just adding
	intsum := int1 + int2 + int3
	// Basic Println just prints sum
	fmt.Println("Sum:", intsum)
	// assiaging floats multiple at a time
	float1, float2, float3 := 13.1, 14.2, 15.3
	// Just adding
	floatsum := float1 + float2 + float3
	// printing value (unpresice unless you use math/big)
	fmt.Println("Sum:", floatsum)

	// Assigning using var becuase I am not setting them yet
	var b1, b2, b3, bigsum big.Float
	// Setting value with setfloat64
	b1.SetFloat64(13.1)
	b2.SetFloat64(14.2)
	b1.SetFloat64(15.3)
	//adding with .add use & to show it is a pointer then adding b3 seperatly becuase it only accepts to args
	bigsum.Add(&b1, &b2).Add(&bigsum, &b3)

	//printf with precision
	fmt.Printf("BigSum = %.10g\n", &bigsum)
	// set radius
	cirleradius := 15.5
	// find circumference
	circumference := cirleradius * math.Pi
	// print and sqelch to 2 decimal places
	fmt.Printf("Circumference: %.2f\n", circumference)
}
