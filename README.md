# bubl
[![Go Report Card](https://goreportcard.com/badge/github.com/alexbarksdale/bubl)](https://goreportcard.com/report/github.com/alexbarksdale/bubl)

Bubl is a flexible and easy-to-use boilerplate generator. Bubl lets you create a template from any file or directory on your system. No need to rely on third party services to create a template that fits your needs, create it yourself.

![bubl](https://i.imgur.com/48sF3mT.gif)

## Table of Contents
* [Getting Started](#getting-started)
  * [Prerequisites](#prereq)
  * [Installation - Global](#installation-global)
  * [Installation - Local](#installation-local)
* [Commands](#commands)

## Getting Started
<a name="prereq"></a>
### Prerequisites
1. If you haven't already, install [Go](https://golang.org/). Make sure you installed it correctly by running: `$ go`
2. If you would like to use bubl anywhere on your machine make sure that your `$GOBIN` is exported into your path. i.e `.bashrc` or  `.zshrc`

<a name="installation-global"></a>
### Installation - Global
1. If you have your `$GOBIN` exported into your path, follow these instructions for the best results.
  
```
$ go get github.com/alexbarksdale/bubl

$ cd $GOPATH/src/github.com/alexbarksdale/bubl

$ go install
```
2. Test to make sure things were installed correctly by navigating into another directory and running `$ bubl` in your terminal.
3. If you're unable to run bubl globally, check to make sure your `$GOBIN` is exported correctly. Otherwise, you can still run this project in a local directory, but that's quite inconvenient.

<a name="installation-local"></a>
### Installation - Local
**TLDR:** This way of using the program is very inconvenient.
  
```
$ git clone https://github.com/alexbarksdale/bubl.git

$ cd bubl

$ go run main.go
```

<a name="commands"></a>
# Commands

```Usage: bubl <command>

$ bubl create <template-path> <bubl-alias>
	Create a bubble by providing a path to your template 
	and an alias to identify your bubble.

$ bubl gen (options) <bubl-alias>
	Generate a template from a bubble to your current directory.

	OPTIONS:
	-bundle
		Bundle together an arbitrary amount of bubbles to generate.

$ bubl pop <bubl-alias>
	Remove a bubble template.

$ bubl list
	List out created bubbles.
```
