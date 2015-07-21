raft-leveldb
===========

This repository provides the `raftleveldb` package. The package exports the
`Store` which is an implementation of both a `LogStore` and `StableStore`.

It is meant to be used as a backend for the `raft` [package
here](https://github.com/hashicorp/raft).

This implementation uses [LevelDB](https://github.com/syndtr/goleveldb).
