# go-between [![Build Status][travis-img]][travis-url] [![Doc][godoc-img]][godoc-url]
> get string between 2 strings

[travis-img]: https://semaphoreci.com/api/v1/projects/59441dc0-ebbd-402f-a746-89519a6aa108/657454/badge.svg
[travis-url]: https://semaphoreci.com/ddo/go-between

[godoc-img]: https://img.shields.io/badge/godoc-Reference-brightgreen.svg?style=flat-square
[godoc-url]: https://godoc.org/github.com/ddo/go-between

## Installation

```bash
go get github.com/ddo/go-between
```

## Usage

```js
import . "between"

Between("hello world", "hel", "ld"); // => "lo wor"
Between("i'm so cool", "m", "oo"); // => " so c"
Between("world hello world", "hel", "ld"); // => "lo wor"
```