package profiler

import (
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"sync"
	"time"
)

// OperationTiming represents the timing for a specific operation, including the operation name and its duration in milliseconds.
type OperationTiming struct {
	Operation string  `json:"Operation"` // The name of the operation being timed
	TimeMs    float64 `json:"TimeMs"`    // The time taken by the operation in milliseconds
}

// Profiler is a structure used to track the execution time of different operations.
// It can be enabled or disabled and supports logging with zap.Logger.
type Profiler struct {
	enabled bool              // Indicates whether the profiler is enabled or not
	timings []OperationTiming // A slice that stores the timings of the tracked operations
	logger  *zap.Logger       // Logger for logging the timings (optional)
	mutex   sync.Mutex        // Mutex to ensure thread-safe access to the timings slice
}

// NewProfiler creates and returns a new Profiler instance. The profiler can be enabled or disabled, and
// it can optionally use a zap.Logger for logging the timings.
func NewProfiler(enabled bool, logger *zap.Logger) *Profiler {
	return &Profiler{
		enabled: enabled,
		timings: []OperationTiming{},
		logger:  logger,
	}
}

// recordAndLogTiming is a helper method that calculates the duration of an operation,
// logs it (if logging is enabled), and stores it in the timings slice.
func (p *Profiler) recordAndLogTiming(operationName string, start time.Time) {
	if !p.enabled {
		return
	}

	duration := time.Since(start).Seconds() * 1000 // Calculate duration in milliseconds

	// Lock the mutex before writing to the timings slice
	p.mutex.Lock()
	p.timings = append(p.timings, OperationTiming{
		Operation: operationName,
		TimeMs:    duration,
	})
	p.mutex.Unlock()

	// Log the timing if the logger is available
	if p.logger != nil {
		p.logger.Debug("Operation timing recorded",
			zap.String("operation", operationName),
			zap.Float64("duration_ms", duration))
	}
}

// DeferTiming is a method for timing operations using Go's `defer` mechanism. It returns a function
// that, when called (usually in a `defer` statement), records the time taken for an operation.
func (p *Profiler) DeferTiming(operationName string) func() {
	start := time.Now()
	return func() {
		p.recordAndLogTiming(operationName, start)
	}
}

// TimeOperation times a function without a return value. It takes the name of the operation
// and a function to time, logs the duration, and stores it in the timings slice.
func (p *Profiler) TimeOperation(operationName string, fn func()) {
	if !p.enabled {
		fn()
		return
	}

	start := time.Now()
	fn()
	p.recordAndLogTiming(operationName, start)
}

// TimeOperationWithReturn times a function that returns a value. It takes the name of the operation
// and a function to time, logs the duration, stores it in the timings slice, and returns the result of the function.
func (p *Profiler) TimeOperationWithReturn(operationName string, fn func() interface{}) interface{} {
	if !p.enabled {
		return fn()
	}

	start := time.Now()
	result := fn()
	p.recordAndLogTiming(operationName, start)
	return result
}

// StartTiming starts the manual timing of an operation and returns the current time.
// Use this method when you need to control both the start and stop of the timing manually.
func (p *Profiler) StartTiming() time.Time {
	return time.Now() // Return the current time as the start time
}

// StopTiming stops the manual timing of an operation by taking the start time as a parameter,
// calculating the duration, and recording the result.
func (p *Profiler) StopTiming(operationName string, start time.Time) {
	p.recordAndLogTiming(operationName, start)
}

// GetTimings returns a slice of all the recorded timings. It is thread-safe as the mutex
// is locked during the read operation.
func (p *Profiler) GetTimings() []OperationTiming {
	p.mutex.Lock()         // Lock the mutex before reading the timings slice
	defer p.mutex.Unlock() // Ensure unlocking after reading
	return p.timings
}

// GetTotalExecutionTime returns the total time of all operations tracked by the profiler in milliseconds.
// It sums the timings in the slice and ensures thread safety with a mutex.
func (p *Profiler) GetTotalExecutionTime() float64 {
	total := 0.0
	p.mutex.Lock()         // Lock the mutex before reading the timings slice
	defer p.mutex.Unlock() // Ensure unlocking after reading
	if len(p.timings) > 0 {
		total, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", p.timings[len(p.timings)-1].TimeMs), 64)
	}
	return total
}

// Reset clears all recorded timings. It ensures thread safety by locking the mutex during the operation.
func (p *Profiler) Reset() {
	p.mutex.Lock()         // Lock the mutex before clearing the timings slice
	defer p.mutex.Unlock() // Ensure unlocking after clearing
	p.timings = []OperationTiming{}
}
