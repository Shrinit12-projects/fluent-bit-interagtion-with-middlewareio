package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// LogEntry represents a structured log entry with various fields for monitoring
type LogEntry struct {
	Timestamp    string  `json:"timestamp"`
	Level        string  `json:"level"`
	Service      string  `json:"service"`
	Message      string  `json:"message"`
	UserID       string  `json:"user_id,omitempty"`
	Endpoint     string  `json:"endpoint,omitempty"`
	ResponseTime int     `json:"response_time_ms,omitempty"`
	StatusCode   int     `json:"status_code,omitempty"`
	Region       string  `json:"region,omitempty"`
	Component    string  `json:"component,omitempty"`
}

// Configuration variables for log generation and rotation
var (
	// Sample data for generating realistic logs
	users     = []string{"user_001", "user_002", "user_003", "user_004", "user_005"}
	endpoints = []string{"/api/login", "/api/users", "/api/orders", "/api/products", "/api/payments"}
	regions   = []string{"us-east-1", "us-west-2", "eu-west-1", "ap-south-1"}
	components = []string{"auth-service", "user-service", "order-service", "payment-service", "notification-service"}
	services  = []string{"web-server", "api-gateway", "database", "cache", "queue"}
	
	// Log rotation configuration
	logFile   = "/var/log/app.log"           // Main log file path
	maxSize   = int64(10 * 1024 * 1024)      // 10MB - rotate when file exceeds this size
	maxFiles  = 5                            // Keep 5 historical log files (app.log.1 to app.log.5)
)

// rotateLog handles log file rotation when the current log file exceeds maxSize
// It shifts existing rotated files (app.log.1 -> app.log.2, etc.) and moves current log to app.log.1
func rotateLog() {
	// Check if current log file exists and exceeds size limit
	info, err := os.Stat(logFile)
	if err != nil || info.Size() < maxSize {
		return // No rotation needed
	}

	// Shift existing rotated files: app.log.4 -> app.log.5, app.log.3 -> app.log.4, etc.
	for i := maxFiles - 1; i > 0; i-- {
		old := fmt.Sprintf("%s.%d", logFile, i)
		new := fmt.Sprintf("%s.%d", logFile, i+1)
		os.Rename(old, new) // Oldest file (app.log.5) gets overwritten
	}

	// Move current active log file to app.log.1
	os.Rename(logFile, logFile+".1")
}

// writeLog writes a log entry to the file, handling rotation automatically
func writeLog(entry LogEntry) {
	// Check and perform log rotation if needed
	rotateLog()

	// Open log file for appending (create if doesn't exist)
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Set current timestamp and write JSON log entry
	entry.Timestamp = time.Now().Format(time.RFC3339)
	jsonLog, _ := json.Marshal(entry)
	file.Write(append(jsonLog, '\n'))
}

// generateLogs creates realistic log entries with various types:
// - API request logs with user activity, performance metrics
// - Component health logs with error/warning/info levels
// - Debug logs for system processing information
func generateLogs() {
	rand.Seed(time.Now().UnixNano())
	
	// Generate API request log with realistic user interaction data
	user := users[rand.Intn(len(users))]
	endpoint := endpoints[rand.Intn(len(endpoints))]
	responseTime := rand.Intn(500) + 50  // 50-550ms response time
	statusCode := []int{200, 201, 400, 401, 404, 500}[rand.Intn(6)]  // Mix of success/error codes
	
	writeLog(LogEntry{
		Level:        "INFO",
		Service:      "api-gateway",
		Message:      "API request processed",
		UserID:       user,
		Endpoint:     endpoint,
		ResponseTime: responseTime,
		StatusCode:   statusCode,
		Region:       regions[rand.Intn(len(regions))],
	})

	// Generate component health logs with realistic error rates
	component := components[rand.Intn(len(components))]
	service := services[rand.Intn(len(services))]
	
	if rand.Float32() < 0.1 { // 10% error rate - realistic for production systems
		writeLog(LogEntry{
			Level:     "ERROR",
			Service:   service,
			Message:   fmt.Sprintf("%s encountered an error", component),
			Component: component,
			Region:    regions[rand.Intn(len(regions))],
		})
	} else if rand.Float32() < 0.2 { // 20% warning rate - performance degradation
		writeLog(LogEntry{
			Level:     "WARN",
			Service:   service,
			Message:   fmt.Sprintf("%s performance degraded", component),
			Component: component,
			Region:    regions[rand.Intn(len(regions))],
		})
	} else { // 70% normal operation
		writeLog(LogEntry{
			Level:     "INFO",
			Service:   service,
			Message:   fmt.Sprintf("%s operating normally", component),
			Component: component,
			Region:    regions[rand.Intn(len(regions))],
		})
	}

	// Generate debug logs occasionally (30% chance) for system processing info
	if rand.Float32() < 0.3 {
		writeLog(LogEntry{
			Level:   "DEBUG",
			Service: "debug-service",
			Message: fmt.Sprintf("Processing batch of %d items", rand.Intn(100)+1),
			Region:  regions[rand.Intn(len(regions))],
		})
	}
}

// main function starts the enhanced logging service with automatic log rotation
func main() {
	log.Println("Starting enhanced Go logging service with log rotation...")
	log.Printf("Log rotation: %dMB max size, %d files retained", maxSize/(1024*1024), maxFiles)

	// Continuous log generation with random intervals for realistic traffic patterns
	for {
		generateLogs()
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second) // 1-3 second intervals
	}
}
