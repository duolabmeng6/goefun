<p align="center">
  <p align="center">
  <a href="https://travis-ci.org/jahnestacado/cable"><img alt="build" 
  src="https://travis-ci.org/jahnestacado/cable.svg?branch=master"></a>
    <a href="https://github.com/jahnestacado/cable/blob/master/LICENSE"><img alt="Software License" src="https://img.shields.io/github/license/mashape/apistatus.svg?style=flat-square"></a>
    <a href="https://goreportcard.com/report/github.com/jahnestacado/cable"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/jahnestacado/cable?style=flat-square&fuckgithubcache=1"></a>
    <a href="https://godoc.org/github.com/jahnestacado/cable">
        <img alt="Docs" src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square">
    </a>
    <a href="https://codecov.io/gh/jahnestacado/cable">
  <img src="https://codecov.io/gh/jahnestacado/cable/branch/master/graph/badge.svg" />
</a>
  <img src="https://github.com/jahnestacado/cable/blob/master/resources/cable-img.webp?raw=true" /img>
  </p>
</p>

# Cable
Utility belt package for scheduling/limiting function calls (throttle, debounce, setTimeout, setInterval)

## Install
```go get github.com/jahnestacado/cable```

## API

#### func  Throttle

```go
func Throttle(f func(), interval time.Duration) func()
```
Throttle returns a function that no matter how many times it is invoked, it will
only execute once within the specified interval

#### type ThrottleOptions

```go
type ThrottleOptions struct {
	Immediate bool
}
```

ThrottleOptions is used to further configure the behavior of a
throttled-function

#### func  Debounce

```go
func Debounce(f func(), interval time.Duration, options DebounceOptions) func()
```
Debounce returns a function that no matter how many times it is invoked, it will
only execute after the specified interval has passed from its last invocation

#### type DebounceOptions

```go
type DebounceOptions struct {
	Immediate bool
}
```

DebounceOptions is used to further configure the behavior of a
debounced-function

#### func  SetInterval

```go
func SetInterval(f func() bool, interval time.Duration) func()
```
SetInterval executes function f repeatedly with a fixed time delay(interval)
between each call until function f returns false. It returns a cancel function
which can be used to cancel aswell the execution of function f

#### func  SetTimeout

```go
func SetTimeout(f func(), interval time.Duration) func()
```
SetTimeout postpones the execution of function f for the specified interval. It
returns a cancel function which when invoked earlier than the specified
interval, it will cancel the execution of function f. Note that function f is
executed in a different goroutine

[GoDoc for cable.go](https://godoc.org/github.com/jahnestacado/cable)

## License
Copyright (c) 2018 Ioannis Tzanellis<br>
[Released under the MIT license](https://github.com/jahnestacado/cable/blob/master/LICENSE) 




