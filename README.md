# Overview

This repo serves as a tutorial of how to use golang interfaces in order to achieve a modular and testable code.
the code is divided into 3 parts:
1. phase 1 - the basic implementation of a s3 downloader. The s3 downloader is a simple program that downloads an object from aws s3 bucket and saves it to a file.
2. phase 2 - the same program but with interfaces and unit tests.
3. phase 3 - the same program as in phase 2 with more code refactoring and better tests coverage.


## Running the tests 
### phase 2
```sh
make test_phase_2
```
### phase 3
```sh
make test_phase_3
```

## Inspecting code coverage
```sh
make test_with_coverage
```