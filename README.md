pfcd
====

[![Build Status](https://travis-ci.org/picfight/pfcd.png?branch=master)](https://travis-ci.org/picfight/pfcd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/picfight/pfcd)

pfcd is a PicFight full node implementation written in Go (golang).

This acts as a chain daemon for the [PicFight](https://picfight.org) cryptocurrency.
pfcd maintains the entire past transactional ledger of PicFight and allows
 relaying of transactions to other PicFight nodes across the world.  To read more
about PicFight please see the
[project documentation](https://docs.picfight.org/#overview).

Note: To send or receive funds and join Proof-of-Stake mining, you will also 
need [pfcwallet](https://github.com/picfight/pfcwallet).

This project is currently under active development and is in a Beta state.  It
is extremely stable and has been in production use since February 2016.

It is forked from [btcd](https://github.com/btcsuite/btcd) which is a bitcoin
full node implementation written in Go.  btcd is a ongoing project under active
development.  Because pfcd is constantly synced with btcd codebase, it will
get the benefit of btcd's ongoing upgrades to peer and connection handling,
database optimization and other blockchain related technology improvements.

## Requirements

[Go](http://golang.org) 1.9 or newer.

## Getting Started

- pfcd (and utilities) will now be installed in either ```$GOROOT/bin``` or
  ```$GOPATH/bin``` depending on your configuration.  If you did not already
  add the bin directory to your system path during Go installation, we
  recommend you do so now.

## Updating

#### Windows

Install a newer MSI

#### Linux/BSD/MacOSX/POSIX - Build from Source

- **Dep**

  Dep is used to manage project dependencies and provide reproducible builds.
  To install:

  `go get -u github.com/golang/dep/cmd/dep`

Unfortunately, the use of `dep` prevents a handy tool such as `go get` from
automatically downloading, building, and installing the source in a single
command.  Instead, the latest project and dependency sources must be first
obtained manually with `git` and `dep`, and then `go` is used to build and
install the project.

**Getting the source**:

For a first time installation, the project and dependency sources can be
obtained manually with `git` and `dep` (create directories as needed):

```
git clone https://github.com/picfight/pfcd $GOPATH/src/github.com/picfight/pfcd
cd $GOPATH/src/github.com/picfight/pfcd
dep ensure
go install . ./cmd/...
```

To update an existing source tree, pull the latest changes and install the
matching dependencies:

```
cd $GOPATH/src/github.com/picfight/pfcd
git pull
dep ensure
go install . ./cmd/...
```

For more information about PicFight and how to set up your software please go to
our docs page at
[docs.picfight.org](https://docs.picfight.org/getting-started/beginner-guide/).

## Docker

### Running pfcd

You can run a picfight node from inside a docker container.  To build the image
yourself, use the following command:

```
docker build -t picfight/pfcd .
```

Or you can create an alpine based image (requires Docker 17.05 or higher):

```
docker build -t picfight/pfcd:alpine -f Dockerfile.alpine .
```

You can then run the image using:

```
docker run picfight/pfcd
```

You may wish to use an external volume to customise your config and persist the
data in an external volume:

```
docker run --rm -v /home/user/pfcdata:/root/.pfcd/data picfight/pfcd
```

For a minimal image, you can use the picfight/pfcd:alpine tag.  This is typically
a more secure option while also being a much smaller image.

You can run pfcctl from inside the image.  For example, run an image (mounting
your data from externally) with:

```
docker run --rm -ti --name=pfcd-1 -v /home/user/.pfcd:/root/.pfcd \
  picfight/pfcd:alpine
```

And then run pfcctl commands against it.  For example:

```
docker exec -ti pfcd-1 pfcctl getbestblock
```


### Running Tests

All tests and linters may be run in a docker container using the script
`run_tests.sh`.  This script defaults to using the current supported version of
go.  You can run it with the major version of Go you would like to use as the
only arguement to test a previous on a previous version of Go (generally PicFight
supports the current version of Go and the previous one).

```
./run_tests.sh 1.9
```

To run the tests locally without docker:

```
./run_tests.sh local
```

## Contact

If you have any further questions you can find us at:

- irc.freenode.net (channel #picfight)
- [webchat](https://webchat.freenode.net/?channels=picfight)
- forum.picfight.org
- picfight.slack.com

## Issue Tracker

The [integrated github issue tracker](https://github.com/picfight/pfcd/issues)
is used for this project.

## Documentation

The documentation is a work-in-progress.  It is located in the
[docs](https://github.com/picfight/pfcd/tree/master/docs) folder.

## License

pfcd is licensed under the [copyfree](http://copyfree.org) ISC License.
