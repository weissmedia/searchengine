package profiler

import (
	"go.uber.org/zap"
	"time"
)

// OperationTiming Struct to store timing information
type OperationTiming struct {
	Operation string  // Description of the operation
	TimeMs    float64 // Time in milliseconds
}

type Profiler struct {
	enabled bool
	timings []OperationTiming
	logger  *zap.Logger // Adding logger to the profiler
}

// NewProfiler creates a new Profiler instance with optional logging
func NewProfiler(enabled bool, logger *zap.Logger) *Profiler {
	return &Profiler{
		enabled: enabled,
		timings: []OperationTiming{},
		logger:  logger,
	}
}

// Helper method to calculate the duration, log it, and add it to the timings
func (p *Profiler) recordTiming(operationName string, start time.Time) {
	if !p.enabled {
		return
	}

	duration := time.Since(start).Seconds() * 1000 // Calculate duration in milliseconds
	timing := OperationTiming{
		Operation: operationName,
		TimeMs:    duration,
	}
	p.timings = append(p.timings, timing)

	// Log the timing if the logger is available
	if p.logger != nil {
		p.logger.Info("Operation timing recorded",
			zap.String("operation", operationName),
			zap.Float64("duration_ms", duration))
	}
}

// DeferTiming for the defer-based timing approach
func (p *Profiler) DeferTiming(operationName string) func() {
	start := time.Now()
	return func() {
		p.recordTiming(operationName, start)
	}
}

// TimeOperation for the normal function call with timing
func (p *Profiler) TimeOperation(operationName string, fn func()) {
	start := time.Now()
	fn()
	p.recordTiming(operationName, start)
}

// StartTiming for manual timing with start time passed in
func (p *Profiler) StartTiming(operationName string, start time.Time) {
	p.recordTiming(operationName, start)
}

// GetTimings returns the recorded timings
func (p *Profiler) GetTimings() []OperationTiming {
	return p.timings
}

// Reset clears all recorded timings
func (p *Profiler) Reset() {
	p.timings = []OperationTiming{}
}
