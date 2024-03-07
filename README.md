[![SLSA Go releaser](https://github.com/jarrodhroberson/cuid2/actions/workflows/go-ossf-slsa3-publish.yml/badge.svg)](https://github.com/jarrodhroberson/cuid2/actions/workflows/go-ossf-slsa3-publish.yml)

# GOCUID2
Go port of CUID2 unique identifier from studying JavaScript and Java implemenations.

## About the project

This is my attempt at writing a CUID2 compliant id generator for Go.

### Status

This project is in the "just got it working and looking for feedback" status.

I based my implemenation on the original paralleldrive JavaScript one and to cross reference some JS stuff I found slighly confusing
As a former Java main for 30 years, I looked at the Java version I found.

Both are listed below.

### See also

* [Original Implemenation of CUID2 in JavaScript](https://github.com/paralleldrive/cuid2)
* [Java version of CUID2 I cross referenced when porting the paralleldrive code](https://github.com/thibaultmeyer/cuid-java)

## Getting started

Below we describe the conventions or tools specific to golang project.

### Layout

```
├── .gitignore
├── cuid2.go
├── cuid2_test.go
├── go.mod
├── go.sum
├── LICENSE
├── README.md
```

A brief description of the layout:
 
* `.gitignore` varies per project, but all projects need to ignore `bin` directory.
* `.cuid2.go` is the package.
* `.cuid2_test.go` is the tests I created for the package.

## Notes

* This looks to me to generate the same thing as the original paralleldrive JavaScript version.
* If it does not please tell my why in an issue.
