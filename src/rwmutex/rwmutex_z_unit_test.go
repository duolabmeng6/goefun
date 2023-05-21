// Copyright 2018 gf Author(https://github.com/gogf/gf/v2/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf/v2/.

package rwmutex_test

import (
	"testing"
	"time"

	"github.com/duolabmeng6/goefun/src/rwmutex"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestRwmutexIsSafe(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		lock := rwmutex.New()
		gtest.Assert(lock.IsSafe(), false)

		lock = rwmutex.New(false)
		gtest.Assert(lock.IsSafe(), false)

		lock = rwmutex.New(false, false)
		gtest.Assert(lock.IsSafe(), false)

		lock = rwmutex.New(true, false)
		gtest.Assert(lock.IsSafe(), true)

		lock = rwmutex.New(true, true)
		gtest.Assert(lock.IsSafe(), true)

		lock = rwmutex.New(true)
		gtest.Assert(lock.IsSafe(), true)
	})
}

func TestSafeRwmutex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		safeLock := rwmutex.New(true)
		array := garray.New(true)

		go func() {
			safeLock.Lock()
			array.Append(1)
			time.Sleep(100 * time.Millisecond)
			array.Append(1)
			safeLock.Unlock()
		}()
		go func() {
			time.Sleep(10 * time.Millisecond)
			safeLock.Lock()
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			array.Append(1)
			safeLock.Unlock()
		}()
		time.Sleep(50 * time.Millisecond)
		gtest.Assert(array.Len(), 1)
		time.Sleep(80 * time.Millisecond)
		gtest.Assert(array.Len(), 3)
		time.Sleep(100 * time.Millisecond)
		gtest.Assert(array.Len(), 3)
		time.Sleep(100 * time.Millisecond)
		gtest.Assert(array.Len(), 4)
	})
}

func TestSafeReaderRwmutex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		safeLock := rwmutex.New(true)
		array := garray.New(true)

		go func() {
			safeLock.RLock()
			array.Append(1)
			time.Sleep(100 * time.Millisecond)
			array.Append(1)
			safeLock.RUnlock()
		}()
		go func() {
			time.Sleep(10 * time.Millisecond)
			safeLock.RLock()
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			array.Append(1)
			time.Sleep(100 * time.Millisecond)
			array.Append(1)
			safeLock.RUnlock()
		}()
		go func() {
			time.Sleep(50 * time.Millisecond)
			safeLock.Lock()
			array.Append(1)
			safeLock.Unlock()
		}()
		time.Sleep(50 * time.Millisecond)
		gtest.Assert(array.Len(), 2)
		time.Sleep(100 * time.Millisecond)
		gtest.Assert(array.Len(), 3)
		time.Sleep(100 * time.Millisecond)
		gtest.Assert(array.Len(), 4)
		time.Sleep(100 * time.Millisecond)
		gtest.Assert(array.Len(), 6)
	})
}

func TestUnsafeRwmutex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		unsafeLock := rwmutex.New()
		array := garray.New(true)

		go func() {
			unsafeLock.Lock()
			array.Append(1)
			time.Sleep(100 * time.Millisecond)
			array.Append(1)
			unsafeLock.Unlock()
		}()
		go func() {
			time.Sleep(10 * time.Millisecond)
			unsafeLock.Lock()
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			array.Append(1)
			unsafeLock.Unlock()
		}()
		time.Sleep(50 * time.Millisecond)
		gtest.Assert(array.Len(), 2)
		time.Sleep(100 * time.Millisecond)
		gtest.Assert(array.Len(), 3)
		time.Sleep(50 * time.Millisecond)
		gtest.Assert(array.Len(), 3)
		time.Sleep(100 * time.Millisecond)
		gtest.Assert(array.Len(), 4)
	})
}
