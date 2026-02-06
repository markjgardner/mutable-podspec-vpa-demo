package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	// Read configuration from environment variables
	memoryMB := getEnvInt("MEMORY_ALLOC_MB", 512)
	waitSeconds := getEnvInt("WAIT_SECONDS", 30)

	fmt.Printf("Starting VPA Demo Application\n")
	fmt.Printf("Configuration:\n")
	fmt.Printf("  - Memory to allocate: %d MB\n", memoryMB)
	fmt.Printf("  - Wait time before release: %d seconds\n", waitSeconds)

	// Allocate memory
	fmt.Printf("Allocating %d MB of memory...\n", memoryMB)
	data := allocateMemory(memoryMB)
	
	// Print memory stats after allocation
	printMemoryStats()

	// Wait for the configured duration
	fmt.Printf("Waiting for %d seconds before releasing memory...\n", waitSeconds)
	time.Sleep(time.Duration(waitSeconds) * time.Second)

	// Release memory by clearing the reference
	fmt.Println("Releasing memory...")
	data = nil // Clear reference to allow GC
	runtime.GC() // Force garbage collection
	// Ensure data is used to avoid compiler optimization
	_ = data
	
	// Print memory stats after release
	printMemoryStats()

	// Continue running with minimal resources
	fmt.Println("Memory released. Running with minimal resources...")
	fmt.Println("Application will continue running indefinitely. Press Ctrl+C to stop.")
	
	// Keep the application running
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			fmt.Printf("Still running... (timestamp: %s)\n", time.Now().Format(time.RFC3339))
			printMemoryStats()
		}
	}
}

// allocateMemory allocates the specified amount of memory in megabytes
func allocateMemory(mb int) []byte {
	// Allocate memory as a byte slice
	size := mb * 1024 * 1024 // Convert MB to bytes
	data := make([]byte, size)
	
	// Write to the memory to ensure it's actually allocated (not just reserved)
	for i := 0; i < len(data); i += 4096 {
		data[i] = 1
	}
	
	fmt.Printf("Successfully allocated %d MB of memory\n", mb)
	return data
}

// getEnvInt reads an integer from environment variable with a default value
func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	
	intValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("Warning: Invalid value for %s (%s), using default %d\n", key, value, defaultValue)
		return defaultValue
	}
	
	return intValue
}

// printMemoryStats prints current memory usage statistics
func printMemoryStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("Memory Stats:\n")
	fmt.Printf("  - Allocated: %d MB\n", m.Alloc/1024/1024)
	fmt.Printf("  - Total Allocated: %d MB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("  - System Memory: %d MB\n", m.Sys/1024/1024)
	fmt.Printf("  - Num GC: %d\n", m.NumGC)
}
