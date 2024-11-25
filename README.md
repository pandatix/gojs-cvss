<div align="center">
	<h1>GoJS-CVSS</h1>
    <p>Transpilation of <a href="https://github.com/pandatix/go-cvss">github.com/pandatix/go-cvss</a> into JS using <a href="https://github.com/gopherjs/gopherjs">gopherjs</a></p>
	<a href="https://pkg.go.dev/github.com/pandatix/gojs-cvss"><img src="https://shields.io/badge/-reference-blue?logo=go&style=for-the-badge" alt="reference"></a>
	<a href="https://goreportcard.com/report/github.com/pandatix/gojs-cvss"><img src="https://goreportcard.com/badge/github.com/pandatix/gojs-cvss?style=for-the-badge" alt="go report"></a>
	<a href="https://coveralls.io/github/pandatix/gojs-cvss?branch=main"><img src="https://img.shields.io/coverallsCoverage/github/pandatix/gojs-cvss?style=for-the-badge" alt="Coverage Status"></a>
	<br>
	<a href=""><img src="https://img.shields.io/github/license/pandatix/gojs-cvss?style=for-the-badge" alt="License"></a>
	<a href="https://github.com/pandatix/gojs-cvss/actions?query=workflow%3Aci+"><img src="https://img.shields.io/github/actions/workflow/status/pandatix/gojs-cvss/ci.yaml?style=for-the-badge&label=CI" alt="CI"></a>
	<a href="https://github.com/pandatix/gojs-cvss/actions/workflows/codeql-analysis.yaml"><img src="https://img.shields.io/github/actions/workflow/status/pandatix/gojs-cvss/codeql-analysis.yaml?style=for-the-badge&label=CodeQL" alt="CodeQL"></a>
	<br>
	<a href="https://securityscorecards.dev/viewer/?uri=github.com/pandatix/gojs-cvss"><img src="https://img.shields.io/ossf-scorecard/github.com/pandatix/gojs-cvss?label=openssf%20scorecard&style=for-the-badge" alt="OpenSSF Scoreboard"></a>
</div>

> [!WARNING]
> This project is not maintained anymore due to its transpiled size, lack of tests and no need for such a library.
> Please use [js-cvss](https://github.com/pandatix/js-cvss) or an alternative.

It currently supports :
 - [ ] [CVSS 2.0](https://www.first.org/cvss/v2/guide)
 - [ ] [CVSS 3.0](https://www.first.org/cvss/v3.0/specification-document)
 - [ ] [CVSS 3.1](https://www.first.org/cvss/v3.1/specification-document)
 - [X] [CVSS 4.0](https://www.first.org/cvss/v4.0/specification-document)

## How to use

### API

The API provides multiple ways to handle CVSS v4.0:
 - as an in-browser global script: `<script src="gojs-cvss.js"></script>`
 - as a Node module: `m = require("./gojs-cvss.js")`

From the global/module import, you have:
 - the `CVSS40` type that you can create either with `vec = new CVSS40()` or `vec = m.cvss40.New()`
 - the `Rating` method that cou can call with `Rating(...)` or `m.Rating(...)`, takes a Number and returns a 2-objects slice composed of the rating (`String` or `null`) and an error message (`String` or `null`) if the value is out of bounds

The `CVSS40` object has multiple methods:
 - `Parse(vector)`: takes a CVSS v4.0 string representation to parse it and returns an error message (`String` or `null`)
 - `Vector()`: returns the string representation (`String`) of the CVSS v4.0 object.
 - `Get(metric)`: takes the metric (`String`) to get and returns a 2-objects slice composed of the metric value (`String` or `null`) and an error message (`String` or `null`)
 - `Set(metric, value)`: takes the metric and the value (`String`) to set it to, and returns an error message (`String` or `null`)
 - `Score()`: returns the score of the CVSS v4.0 object (`Number`)
 - `Nomenclature()`: returns the nomenclature of the CVSS v4.0 object (`String`)

The [examples](examples) are equivalent and shows how to parse a CVSS v4.0 string representation, change a metric value, compute its score+nomenclature and print the new vector.
