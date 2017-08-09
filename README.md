# epp - Environment preprocessor

A small templating engine that allows you to use environmental variables. For
example:
```yaml
{% if HOME %}
...
{% endif %}
```

Will check if `$HOME` is set.

## NOTE about two epp versions

**NOTE** that there are currently _two_ versions of epp that you can use. The regular one, which uses the above `{{ HOME }}` syntax to expand environment variables, and the `{{ env "HOME" }}` syntax for the [`v2.0.0-rc1`](https://github.com/blendle/epp/releases/tag/v2.0.0-rc1) version.

The 2.0 variant is used in our Charts (to be in line with the templating style in Helm itself), while the legacy variant is used in our `config/deploy/resources/*.yml` files, for legacy reasons.

In the future, we want to merge the two, and have all situations use the new syntax style.

## Installation
```
$ go get github.com/blendle/epp
```

Or download the latest release.

## Usage
```
$ epp inputfile -o output
$ epp - < input > output
$ epp inputfile # Write to stdout
```
