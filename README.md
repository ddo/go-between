# go-between [![Build Status][travis-img]][travis-url] [![Doc][godoc-img]][godoc-url]
> get string between 2 strings

[travis-img]: https://img.shields.io/travis/ddo/go-between.svg?style=flat-square
[travis-url]: https://travis-ci.org/ddo/go-between

[godoc-img]: https://img.shields.io/badge/godoc-Reference-brightgreen.svg?style=flat-square
[godoc-url]: https://godoc.org/github.com/ddo/go-between

## Installation

```
go get github.com/ddo/go-between
```

## Usage

```js
import . "between"

Between("hello world", "hel", "ld"); // => "lo wor"
Between("i'm so cool", "m", "oo"); // => " so c"
Between("world hello world", "hel", "ld"); // => "lo wor"
```