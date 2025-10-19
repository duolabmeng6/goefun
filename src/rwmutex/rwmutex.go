// Copyright 2018 gf Author(https://github.com/gogf/gf/v2/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf/v2/.

// Package rwmutex 为 sync.RWMutex 提供可切换的并发安全特性
package rwmutex

import "sync"

// RWMutex is a sync.RWMutex with a switch of concurrent safe feature.
type RWMutex struct {
    sync.RWMutex
    safe bool
}

// New creates and returns a new *RWMutex.
// The parameter <safe> is used to specify whether using this mutex in concurrent-safety,
// which is false in default.
func New(safe ...bool) *RWMutex {
    mu := new(RWMutex)
    if len(safe) > 0 {
        mu.safe = safe[0]
    } else {
        mu.safe = false
    }
    return mu
}

func (mu *RWMutex) IsSafe() bool {
    return mu.safe
}

func (mu *RWMutex) Lock() {
    if mu.safe {
        mu.RWMutex.Lock()
    }
}

func (mu *RWMutex) Unlock() {
    if mu.safe {
        mu.RWMutex.Unlock()
    }
}

func (mu *RWMutex) RLock() {
    if mu.safe {
        mu.RWMutex.RLock()
    }
}

func (mu *RWMutex) RUnlock() {
    if mu.safe {
        mu.RWMutex.RUnlock()
    }
}
