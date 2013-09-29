[![Build Status](https://travis-ci.org/tristanwietsma/metastore.png?branch=master)](https://travis-ci.org/tristanwietsma/metastore)

MetaStore
=========

A Store is an abstraction over a string map that supports get, set, delete, publish, and subscribe.

A MetaStore ("a Store of Stores") is an abstraction over a Store that breaks the key space into buckets. By doing so, we get finer lock granularity when deployed in a concurrent environment, such as in JackDB (where MetaStore is used).
