package main

import (
	"errors"
	"fmt"
)

// Telemetry represents the data read from a site device.
type Telemetry struct {
	DeviceID string
	PowerKW  float64 // Power generated (positive) or consumed (negative)
	Status   string  // "ONLINE", "FAULT", "OFFLINE"
}

// Device defines the contract for any device connected to the Tesla Site Controller.
type Device interface {
	FetchTelemetry() (Telemetry, error)
}

// ============================================================================
// EXERCISE 1: Implement Megapack Battery Storage (Pointer Receiver)
// ============================================================================
// 1. Define a struct `Megapack` with fields:
//    - ID (string)
//    - StateOfCharge (float64) - value between 0.0 and 100.0
//    - PowerKW (float64)
//
// 2. Implement the `Device` interface on `*Megapack` (pointer receiver).
//    - Requirements:
//      - If StateOfCharge < 0.0 or StateOfCharge > 100.0, return an empty Telemetry and an error
//        explaining that StateOfCharge is invalid.
//      - If StateOfCharge is valid, return a Telemetry struct with status "ONLINE" and the PowerKW.

// TODO: Define your Megapack struct and its FetchTelemetry() method here.



// ============================================================================
// EXERCISE 2: Implement SolarInverter (Value Receiver)
// ============================================================================
// 1. Define a struct `SolarInverter` with fields:
//    - ID (string)
//    - Irradiance (float64) - W/m^2 solar intensity
//    - Efficiency (float64) - scaling factor (e.g. 0.15 for 15% efficiency)
//
// 2. Implement the `Device` interface on `SolarInverter` (value receiver).
//    - Requirements:
//      - Calculate power generated: PowerKW = Irradiance * Efficiency.
//      - If Irradiance < 0.0, return a Telemetry struct with status "OFFLINE", PowerKW = 0.0, and NO error.
//      - Otherwise, return a Telemetry struct with status "ONLINE" and the calculated PowerKW.

// TODO: Define your SolarInverter struct and its FetchTelemetry() method here.



// ============================================================================
// EXERCISE 3: Interface Aggregator
// ============================================================================
// Implement `GetTotalPower` which aggregates the total power from a slice of Devices.
// - Call FetchTelemetry() on each device.
// - If a device is "ONLINE", add its PowerKW to the total sum.
// - If a device returns an error or status is "FAULT" or "OFFLINE", log it (using fmt.Println),
//   but DO NOT abort. Continue aggregating power from other operational devices.
// - Return the accumulated power.
func GetTotalPower(devices []Device) float64 {
	// TODO: Implement the aggregator logic.
	return 0.0
}

// ============================================================================
// EXERCISE 4: Conceptual Questions (Answer in your head or in the comments)
// ============================================================================
// Q1: Why does passing a slice of concrete type `[]Megapack` to a function expecting `[]Device`
//     fail to compile? How do we resolve this in Go?
//
// Q2: Consider the following snippet:
//     var mp *Megapack = nil
//     var dev Device = mp
//     What happens if we check `if dev == nil`? Will it evaluate to true or false? Why?
// ============================================================================

func main() {
	fmt.Println("--- Running Day 1: Interfaces Practice Tests ---")
	
	// Test 1: Megapack (Valid)
	// var _ Device = &Megapack{} // Uncomment to verify interface implementation compiled.
	
	// Test 2: SolarInverter (Valid)
	// var _ Device = SolarInverter{} // Uncomment to verify interface implementation compiled.

	// Run validation logic
	runTests()
}

// Do not modify runTests() yet. It is used to check your answers.
func runTests() {
	var mp Device
	var inv Device

	// Let's use reflection/assertions to check if methods are implemented
	fmt.Println("Verifying implementation of Megapack and SolarInverter...")

	// 1. Verify Megapack interface matching
	// We use a helper function to avoid compiler errors if not yet defined
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("❌ Panic occurred during testing: %v\nCheck that you implemented all structs and methods correctly.\n", r)
		}
	}()

	// We construct a mock wrapper or cast to test Megapack
	// (using reflection to check if Megapack exists and matches)
	fmt.Println("Please implement the Megapack and SolarInverter structs first, then uncomment the test assertions in main().")
}
