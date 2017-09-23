# epp - Easy preprocessor

A small templating engine that allows you to process templated strings or files.

For example:

```yaml
{{ if (env "HOME") }}
...
{{ end }}
```

Will check if `$HOME` is set.

You can use all of the helpers provided by the [Sprig][sprig] library

[sprig]: http://masterminds.github.io/sprig/

## Installation

```
$ go get github.com/blendle/epp
```

Or [download one of the stable releases][download].

[download]: https://github.com/blendle/epp/releases

## Usage

```
$ epp inputfile -o output
$ epp - < input > output
$ epp inputfile # Write to stdout
```
