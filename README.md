# bagit
bagit is a library and command line tool to work with the BagIt format

Wikipedia: [https://en.wikipedia.org/wiki/BagIt](https://en.wikipedia.org/wiki/BagIt) 

IETF Draft: [https://tools.ietf.org/html/draft-kunze-bagit-17](https://tools.ietf.org/html/draft-kunze-bagit-17)

# Usage examples


Create a bag:

    gobagit -create testinputdir


Create a bag with all possible commandline options

    gobagit -create testinputdir -output outbagdir -tar -hash sha1 -H headerfile.json


Commandline options:

    -H PATH_TO_FILE
        Additional headers for bag-info.txt. Expects path to json file
    -create PATH_TO_DIR
        Create bag. Expects path to source directory
    -hash ALGORITHM
        Hash algorithm used for manifest file when creating a bag [sha1, sha256, md5] (default "sha256")
    -output PATH_TO_DIR
        Output directory for bag. Used with create flag (default "bag_2019-06-11T225839")
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

1. Only one manifest file is supported in version 0.1.0
2. No additional tag directories are supported in version 0.1.0
3. Issues page [https://github.com/steffenfritz/bagit/issues](https://github.com/steffenfritz/bagit/issues)
