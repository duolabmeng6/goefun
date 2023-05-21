// * cable <https://github.com/jahnestacado/cable>
// * Copyright (c) 2018 Ioannis Tzanellis
// * Licensed under the MIT License (MIT).

package cable

import (
	"sync"
	"testing"
	"time"
)

func Test_SetTimeout(t *testing.T) {
	timeoutInterval1 := 100 * time.Millisecond
	calledAt := time.Now()
	SetTimeout(func() {
		executedAt := time.Now()
		delta := executedAt.Sub(calledAt)
		if delta <= timeoutInterval1 {
			t.Errorf("SetTimeout callback was called earlier: %d, want >: %d.", delta, timeoutInterval1)
		}

	}, timeoutInterval1)
	time.Sleep(200 * time.Millisecond)

	timeoutInterval2 := 50 * time.Millisecond
	isCanceled := true
	cancel := SetTimeout(func() {
		isCanceled = false
	}, timeoutInterval2)

	cancel()
	time.Sleep(100 * time.Millisecond)

	if !isCanceled {
		t.Errorf("SetTimeout cancel callback execution failed")
	}
}

func Test_SetInterval(t *testing.T) {
	interval := time.Duration(20)
	maxTimesInvoked := 5
	timeWindow := 1 * time.Millisecond
	assertAfter := (interval * time.Duration(maxTimesInvoked+1) * time.Millisecond) - timeWindow
	var access sync.Mutex

	var timesInvoked1 int
	SetInterval(func() bool {
		access.Lock()
		defer access.Unlock()
		timesInvoked1++
		if timesInvoked1 == maxTimesInvoked {
			return false
		}
		return true
	}, interval*time.Millisecond)

	time.Sleep(assertAfter)

	access.Lock()
	if timesInvoked1 != 5 {
		t.Errorf(`SetInterval with internal cancelation finished earlier/later.
			 Callback invoked times: %d, want: %d.`, timesInvoked1, maxTimesInvoked)
	}
	access.Unlock()

	var timesInvoked2 int
	totalSetIntervalDuration := (interval*time.Duration(maxTimesInvoked))*time.Millisecond + timeWindow
	cancelSetInterval := SetInterval(func() bool {
		access.Lock()
		defer access.Unlock()
		timesInvoked2++
		return true
	}, interval*time.Millisecond)

	SetTimeout(func() {
		cancelSetInterval()
	}, totalSetIntervalDuration)

	time.Sleep(assertAfter)
	access.Lock()
	if timesInvoked2 != 5 {
		t.Errorf(`SetInterval with external cancelation finished earlier/later.
			 Callback invoked times: %d, want: %d.`, timesInvoked2, maxTimesInvoked)
	}
	access.Unlock()
}

func Test_Throttle(t *testing.T) {
	throttleInterval := 33 * time.Millisecond
	executionInterval := 5 * time.Millisecond
	setIntervalMaxDuration := 200 * time.Millisecond
	var access sync.Mutex

	var timesInvoked1 int
	throttledFunc1 := Throttle(func() {
		access.Lock()
		timesInvoked1++
		access.Unlock()
	}, throttleInterval, ThrottleOptions{})

	startedAt1 := time.Now()
	SetInterval(func() bool {
		delta := time.Now().Sub(startedAt1)
		throttledFunc1()
		if delta > setIntervalMaxDuration {
			return false
		}
		return true
	}, executionInterval)

	time.Sleep(setIntervalMaxDuration + throttleInterval + executionInterval)
	expectedInvocations1 := 7
	access.Lock()
	if timesInvoked1 != expectedInvocations1 {
		t.Errorf("Throttled callback has not been invoked the expected amount of times: %d, want: %d.",
			timesInvoked1, expectedInvocations1)
	}
	access.Unlock()

	var timesInvoked2 int
	throttledFunc2 := Throttle(func() {
		access.Lock()
		timesInvoked2++
		access.Unlock()
	}, throttleInterval, ThrottleOptions{Immediate: true})

	startedAt2 := time.Now()
	SetInterval(func() bool {
		delta := time.Now().Sub(startedAt2)
		throttledFunc2()
		if delta > setIntervalMaxDuration {
			return false
		}
		return true
	}, executionInterval)

	time.Sleep(setIntervalMaxDuration + throttleInterval + executionInterval)
	expectedInvocations2 := expectedInvocations1 + 1
	access.Lock()
	if timesInvoked2 != expectedInvocations2 {
		t.Errorf("Throttled callback has not been invoked the expected amount of times: %d, want: %d.",
			timesInvoked2, expectedInvocations2)
	}
	access.Unlock()
}

func Test_Debounce(t *testing.T) {
	debounceInterval := 30 * time.Millisecond
	executionInterval := 5 * time.Millisecond
	setIntervalMaxDuration := 200 * time.Millisecond
	var timesInvoked1 int
	var timesInvoked2 int
	var startedAt time.Time

	maxExpectedInvocations := 1
	debouncedFunc := Debounce(func() {
		timesInvoked1++
		if timesInvoked1 != maxExpectedInvocations {
			t.Errorf("Debounced callback has not been invoked the expected maximum amount of times: %d, want: %d.",
				timesInvoked1, maxExpectedInvocations)
		}
		if time.Now().Sub(startedAt) <= setIntervalMaxDuration {
			t.Errorf("Debounced callback has not been invoked sooner than expected")
		}
	}, debounceInterval, DebounceOptions{})

	maxExpectedtimesInvoked2 := 2
	debouncedImmediateFunc := Debounce(func() {
		timesInvoked2++
		delta := time.Now().Sub(startedAt)
		if timesInvoked2 > maxExpectedtimesInvoked2 {
			t.Errorf("Debounced immediate callback has not been invoked the expected maximum amount of times: %d, want <=: %d.",
				timesInvoked2, maxExpectedtimesInvoked2)
		}
		if timesInvoked2 == 1 && delta >= setIntervalMaxDuration {
			t.Errorf("Debounced immediate callback has been invoked later than expected")
		}

		if timesInvoked2 == 2 && delta <= setIntervalMaxDuration {
			t.Errorf("Debounced immediate callback has been invoked earlier than expected")
		}

	}, debounceInterval, DebounceOptions{Immediate: true})

	startedAt = time.Now()
	SetInterval(func() bool {
		delta := time.Now().Sub(startedAt)
		debouncedFunc()
		debouncedImmediateFunc()
		if delta > setIntervalMaxDuration {
			return false
		}
		return true
	}, executionInterval)

	time.Sleep(5 * time.Second)
}
