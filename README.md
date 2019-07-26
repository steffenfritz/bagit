# gobagit
gobagit is a command line tool to work with the BagIt format

Wikipedia: [https://en.wikipedia.org/wiki/BagIt](https://en.wikipedia.org/wiki/BagIt) 

IETF: [https://tools.ietf.org/html/rfc8493](https://tools.ietf.org/html/rfc8493)


[![Build Status](https://travis-ci.org/steffenfritz/bagit.svg?branch=master)](https://travis-ci.org/steffenfritz/bagit)
[![codecov](https://codecov.io/gh/steffenfritz/bagit/branch/master/graph/badge.svg)](https://codecov.io/gh/steffenfritz/bagit)
[![Go Report Card](https://goreportcard.com/badge/github.com/steffenfritz/bagit)](https://goreportcard.com/report/github.com/steffenfritz/bagit)


Version: 0.4.0

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

    -create PATH_TO_DIR
        Create bag. Expects path to source directory
    -fetch string
        Adds optional fetch file to bag. Expects path to fetch.txt file and switch manifetch
    -hash ALGORITHM
        Hash algorithm used for manifest file when creating a bag [sha1, sha256, sha512, md5] (default "sha512")
    -header PATH_TO_FILE
        Additional headers for bag-info.txt. Expects path to json file
    -manifetch string
        Path to manifest file for optional fetch.txt file. Mandatory if fetch switch is used
    -output PATH_TO_DIR
        Output directory for bag. Used with create flag (default "bag_2019-06-11T225839")
    -tagmanifest ALGORITHM
        Hash algorithm used for tagmanifest file when creating a bag [sha1, sha256, sha512, md5] (default "sha512")
    -tar
        Create a tar archive when creating a bag
    -v    Verbose output
    -validate PATH_TO_BAG
        Validate bag. Expects path to bag
    -version
        Print version


# Installation

## From source

    go get github.com/steffenfritz/bagit/cmd/gobagit


## Binary


Download a pre-built binary from the releases page.


# Limitations

1. Only one payload manifest file is supported in version 0.4.0
2. No additional tag directories are supported in version 0.4.0
3. Issues page [https://github.com/steffenfritz/bagit/issues](https://github.com/steffenfritz/bagit/issues)
4. When creating a bag fetch does NOT validate if the provided hashes match the hash algorithm for the bag's manifest
