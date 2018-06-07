package util

// Copyright 2018 sunny authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Source code and project home:
//
// https://github.com/sunnyregion/util
//
// Installation:
//
// go get  github.com/sunnyregion/util
//
// Example:
//
//		import "github.com/sunnyregion/util"
// 这个文件我是为了处理GO语言并发编程之互斥锁、读写锁详解
import (
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	Val map[string]string
	Mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string, str strintg) {
	c.Mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.Val[key] += str
	c.Mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) string {
	c.Mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.Mux.Unlock()
	return c.Val[key]
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Clear(key string) {
	c.Mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.Mux.Unlock()
	c.Val[key] = ``
}
