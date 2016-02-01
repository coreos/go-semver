# go-semver - Semantic Versioning Library

[![Build Status](https://travis-ci.org/coreos/go-semver.png)](https://travis-ci.org/coreos/go-semver)

go-semver is a [semantic versioning](http://semver.org) library for Go. It lets you parse and compare two semantic version strings.

## Usage

```go
vA, err := semver.NewVersion("1.2.3")
vB, err := semver.NewVersion("3.2.1")

fmt.Printf("%s < %s == %t\n", vA, vB, vA.LessThan(*vB))
```

## Example Application

```bash
$ go run example.go 1.2.3 3.2.1
1.2.3 < 3.2.1 == true

$ go run example.go 5.2.3 3.2.1
5.2.3 < 3.2.1 == false
```

## Comparison operators

```go
// Return 1 if a > b, -1 if a < b, 0 if a == b
func (a *Version) Cmp(b Version) int
```
```go
// Return true if a < b
func (a *Version) LessThan(b Version) bool
```
```go
// Return true if a == b
func (a *Version) Equal(b Version) bool
```
```go
// Return true if a > b
func (a *Version) GreaterThan(b Version) bool
```
```go
// Return true if a <= b
func (a *Version) LessOrEqual(b Version) bool
```
```go
// Return true if a >= b
func (a *Version) GreaterOrEqual(b Version) bool
```

As per [specification](http://semver.org), each of these ignore build metadata:
```go
vA, err := semver.NewVersion("1.0.0+exp.sha.5114f85")
vB, err := semver.NewVersion("1.0.0+20130313144700")

fmt.Printf("%s == %s: %t\n", vA, vB, vA.Equal(*vB)) // true
```
