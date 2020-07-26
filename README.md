GoCharts
========

[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

GoCharts is a library to assist with building charts:

* [C3](https://c3js.org/) - [code](charts/c3)
* [D3](https://d3js.org/) - [code](charts/d3)
* [Rickshaw](https://github.com/shutterstock/rickshaw) - [code](charts/rickshaw)
* [wcharczuk/go-chart](https://github.com/wcharczuk/go-chart) - [code](charts/wchart)

[`quicktemplate`](https://github.com/valyala/quicktemplate) is used for rendering some of the charts.

An example chart is the Rickshaw chart shown below:

![](charts/rickshaw/graph_example_2.png)

Supporting data objects are also provided including:

* [Table](data/table) - Easy manipulation including [writing to CSV and XLSX](data/table/write.go).
* [Static Time Series](data/statictimeseries) - for building charts.
* [Roadmap](data/roadmap) - For creating Roadmap slides, such as from Aha!

## Installation

```bash
$ go get github.com/grokify/gocharts/...
```

## Usage

See the example here:

[charts/rickshaw/examples/report.go](charts/rickshaw/examples/report.go)

 [build-status-svg]: https://api.travis-ci.org/grokify/gocharts.svg?branch=master
 [build-status-link]: https://travis-ci.org/grokify/gocharts
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/gocharts
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/gocharts
 [docs-godoc-svg]: https://img.shields.io/badge/reference-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/grokify/gocharts
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/gocharts/blob/master/LICENSE.md