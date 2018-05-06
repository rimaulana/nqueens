# Bitwise Approach in Solving N-Queens Problem in Golang

[![Go Report Card](https://goreportcard.com/badge/github.com/rimaulana/nqueens)](https://goreportcard.com/report/github.com/rimaulana/nqueens) [![CircleCI](https://img.shields.io/circleci/project/github/rimaulana/nqueens.svg)](https://circleci.com/gh/rimaulana/nqueens/tree/master) [![codecov](https://codecov.io/gh/rimaulana/nqueens/branch/master/graph/badge.svg)](https://codecov.io/gh/rimaulana/nqueens) [![codebeat badge](https://codebeat.co/badges/c217e8a8-b808-4b35-aee0-a0705874289d)](https://codebeat.co/projects/github-com-rimaulana-nqueens-master) [![Maintainability](https://api.codeclimate.com/v1/badges/4a663411cfea93342333/maintainability)](https://codeclimate.com/github/rimaulana/nqueens/maintainability) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

This program is intended to solve n-queens problem through bitwise approach using go programming language. This algorithm used on these codes is based on a paper by Martin Richards titled Backtracking Algorithms in MCPL using Bit Patterns and Recursion.

## Installation
No need to install, just download the pre-compiled binary on our [releases](https://github.com/rimaulana/nqueens/releases) tab.

## Source Code
In order to get these source codes, you need to have at least go1.9 installed on your system and then you can easily get it with
```bash
go get -u github.com/rimaulana/nqueens
```

## Compile from Source Code
To be able to compile the source code using Makefile we provided, you will at least need shell (git shell or linux shell) that can run make. you can simply run
```bash
make release
```
that command will create a folder release and inside there will be three binary for windows, linux and darwin operating system and all will be for amd64 processor architecture. However, if you prefer to custom compile these codes, you can use go compiler.
```bash
go build
```

## Running the Program
There are several flag or params you can supply when executing the program. they are:

* size (int)(n), this flag will tell the program to find the solution(s) of n x n board with n number of queens.
* sample (int)(m), this flag will tell the program to print m sample of board, if you just interested on the number of solution just don't add this flag when executing the program to deactivate this option.

Sample of command to find solution of 4 Queens on 4x4 boards with 2 samples is.
```bash
./nqueens-v1.0.0-linux-amd64 --size=4 --sample=2
```
That command will give us output like the following
```text
┌─┬─┬─┬─┐
│ │Q│ │ │
├─┼─┼─┼─┤
│ │ │ │Q│
├─┼─┼─┼─┤
│Q│ │ │ │
├─┼─┼─┼─┤
│ │ │Q│ │
└─┴─┴─┴─┘
┌─┬─┬─┬─┐
│ │ │Q│ │
├─┼─┼─┼─┤
│Q│ │ │ │
├─┼─┼─┼─┤
│ │ │ │Q│
├─┼─┼─┼─┤
│ │Q│ │ │
└─┴─┴─┴─┘
2 boards are drawn in 201.767µs
There are 2 solutions for board size 4
Solutions are generated in 7.922µs
```
# Reference
Martin Richards, Backtracking Algorithms in MCPL using Bit Patterns and Recursion, Computer Laboratory, University of Cambridge, February 23, 2009, http://www.cl.cam.ac.uk/users/mr/