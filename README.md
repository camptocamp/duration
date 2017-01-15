duration
===========

Simple binary to convert a duration string to seconds.

A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '1.5h' or '2h45m'.
Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.

# Usage

example:
``` bash
docker run --rm camptocamp/duration 2h30m
```

this will output:
```bash
9000
```

# Build

You can use the centurylink/golang-builder to build this utility

``` bash
docker run --rm -v $(pwd):/src centurylink/golang-builder
```