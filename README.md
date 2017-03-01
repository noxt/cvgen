# cvgen

[![Build Status](https://travis-ci.org/noxt/cvgen.svg?branch=master)](https://travis-ci.org/noxt/cvgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/noxt/cvgen)](https://goreportcard.com/report/github.com/noxt/cvgen)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/noxt/cvgen/master/LICENSE.md)

Build CV sites from YAML files!


## Configuration

```yaml
# config.yml

template:
  repo_url: https://github.com/noxt/cvgen-templates
  path: orbit/
  files: [index.html]
output_dir: output
```


## Usage

`cvgen *command*`


## Commands List

| Name | Description |
|---|---|
| `init` | Setup config file |
| `template install` | Install templates from config file |
| `build` | Compile template |