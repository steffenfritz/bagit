# gobagit
gobagit is a command line tool to work with the BagIt format

Wikipedia: [https://en.wikipedia.org/wiki/BagIt](https://en.wikipedia.org/wiki/BagIt) 

IETF: [https://tools.ietf.org/html/rfc8493](https://tools.ietf.org/html/rfc8493)


[![Build status](https://ci.appveyor.com/api/projects/status/vscholjbph8umbd3?svg=true)](https://ci.appveyor.com/project/steffenfritz/bagit)
[![codecov](https://codecov.io/gh/steffenfritz/bagit/branch/master/graph/badge.svg)](https://codecov.io/gh/steffenfritz/bagit)
[![Go Report Card](https://goreportcard.com/badge/github.com/steffenfritz/bagit)](https://goreportcard.com/report/github.com/steffenfritz/bagit)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=steffenfritz_bagit&metric=alert_status)](https://sonarcloud.io/dashboard?id=steffenfritz_bagit)

Version: 0.5.0

# Usage examples


Create a simple bag:

    gobagit -create testinputdir


Create a bag with some possible commandline options

    gobagit -create testinputdir -output outbagdir -tar -hash sha1 -header headerfile.json


Create a bag with all possible commandline options


    gobagit -create testinputdir -output outbagdir -tar -hash sha1 -header headerfile.json -fetch fetch.txt -manifetch fetchmanifest.txt -tagmanifest sha512

Validate a bag

    gobagit -validate bagdir/

Pass additional headers as a json file, no nesting supported. Example: 

    {
        "Source-Organization": "FOO University",
        "Contact-Email":"steffen@fritz.wtf"
    }


Commandline options:

       --create, -C
               Create bag. Expects path to source directory

       --fetch, -F
               Adds optional fetch file to bag. Expects path to fetch.txt file and flag manifetch

       --hash, -H
               Hash algorithm used for manifest file when creating a bag

       --header, -J
               Additional headers for bag-info.txt. Expects path to json file

       --manifetch, -M
               Path to manifest file for optional fetch.txt file. Mandatory if fetch switch is used

       --output, -O
               Output directory for bag. Used with create flag

       --tagmanifest, -t
               Hash algorithm used for tag manifest file

       --tar, -T
               Create a tar archive when creating a bag

       --validate, -V
               Validate bag. Expects path to bag

       --verbose, -v
               Verbose output
       
       --version
               Print version



# Installation

## From source

    go get github.com/steffenfritz/bagit/cmd/gobagit


## Binary


Download a pre-built binary from the releases page.

# Tests and Validation

The test cases include validation tests of two bags, created with LoC's bagit.py (https://libraryofcongress.github.io/bagit-python/)

One bag is valid, the other isn't. gobagit has to validate both cases to pass the testing.

Bags created with gobagit are validated using bagit.py. The created bags have to be valid to pass the testing.

# Limitations

1. Only one payload manifest file is supported in version 0.4.0
2. No additional tag directories are supported in version 0.4.0
3. Issues page [https://github.com/steffenfritz/bagit/issues](https://github.com/steffenfritz/bagit/issues)
4. When creating a bag fetch does NOT validate if the provided hashes match the hash algorithm for the bag's manifest
