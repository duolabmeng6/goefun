// * cable <https://github.com/jahnestacado/cable>
// * Copyright (c) 2018 Ioannis Tzanellis
// * Licensed under the MIT License (MIT).

// Package cable implements utility functions for scheduling/limiting function calls
package cable

import (
	"sync"
	"time"
)

// ThrottleOptions is used to further configure the behavior of a throttled-function
type ThrottleOptions struct {
	Immediate bool
}

// Throttle returns a function that no matter how many times it is invoked,
// it will only execute once within the specified interval
func Throttle(f func(), interval time.Duration, options ThrottleOptions) func() {
	var last time.Time
	noop := func() {}
	var access sync.Mutex
	cancel := noop

	immediateDone := false
	handleImmediate := func() {
		if options.Immediate && !immediateDone {
			f()
			immediateDone = true
		}
	}

	return func() {
		handleImmediate()
		now := time.Now()
		access.Lock()
		delta := now.Sub(last)
		cancel()
		if delta > interval || last.IsZero() {
			last = now
			f()
			access.Unlock()
		} else {
			cancel = SetTimeout(func() {
				access.Lock()
				last = now
				f()
				access.Unlock()
			}, interval)
			access.Unlock()
		}
	}
}

// DebounceOptions is used to further configure the behavior of a debounced-function
type DebounceOptions struct {
	Immediate bool
}

// Debounce returns a function that no matter how many times it is invoked,
// it will only execute after the specified interval has passed from its last invocation
func Debounce(f func(), interval time.Duration, options DebounceOptions) func() {
	handleImmediateCall := func() {
		if options.Immediate {
			f()
		}
	}
	cancel := handleImmediateCall
	return func() {
		cancel()
		cancel = SetTimeout(f, interval)
	}
}

// SetTimeout postpones the execution of function f for the specified interval.
// It returns a cancel function which when invoked earlier than the specified interval, it will
// cancel the execution of function f. Note that function f is executed in a different goroutine
func SetTimeout(f func(), interval time.Duration) func() {
	var isCanceled bool
	var access sync.Mutex
	go (func() {
		time.Sleep(interval)
		access.Lock()
		defer access.Unlock()
		if !isCanceled {
			f()
		}
	})()

	cancel := func() {
		access.Lock()
		isCanceled = true
		access.Unlock()
	}
	return cancel
}

// SetInterval executes function f repeatedly with a fixed time delay(interval) between each call
// until function f returns false. It returns a cancel function which can be used to cancel aswell
// the execution of function f
func SetInterval(f func() bool, interval time.Duration) func() {
	var access sync.Mutex
	shouldContinue := true
	go (func() {
		for range time.Tick(interval) {
			access.Lock()
			if !shouldContinue {
				access.Unlock()
				break
			}
			shouldContinue = f()
			access.Unlock()
		}
	})()

	cancel := func() {
		access.Lock()
		shouldContinue = false
		access.Unlock()
	}

	return cancel
}
