# Preamble (see https://tech.davis-hansson.com/p/make/)
SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

.PHONY: d2md
.SECONDEXPANSION:
d2md: $$(patsubst %.d2.md, %.html, $$(shell find -name '*.d2.md'))

bin/d2md: d2md/cmd/main.go
	cd d2md/cmd
	go build -o ../../bin/d2md

%.html: %.d2.md bin/d2md
	bin/d2md $*.d2.md > $@
