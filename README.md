# cvgen

[![Build Status](https://travis-ci.org/noxt/cvgen.svg?branch=master)](https://travis-ci.org/noxt/cvgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/noxt/cvgen)](https://goreportcard.com/report/github.com/noxt/cvgen)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/noxt/cvgen/master/LICENSE.md)


Build CV sites from YAML files!

## Configuration

```yaml
# cvgen.yaml

template:
  repo: http://github.com/noxt/cvgen-templates/
  name: orbit
```

## Commands

* `init` Setup config file
* `template`
    * `install` Install templates from config file
* `build` Build CV site