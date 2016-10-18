# Godaytime

[![Build Status](https://travis-ci.org/tdi/godaytime.svg?branch=master)](https://travis-ci.org/tdi/godaytime)

This is intended as an example for Computer Networks 2 class at Poznań University of Technology. Mostly usable by my students. Godaytime lousy implements Daytime protocol from [RFC867](https://tools.ietf.org/html/rfc867). 

For the echo server, check [here](https://github.com/tdi/goecho). 


## Install

```
go get -u github.com/tdi/godaytime
```
This will install godaytime to your `$GOPATH/bin`. 

## Usage 
```
godaytime [-h] [-H HOSTNAME] [-p PORT] [-u]
```

By default godaytime listens on localhost:3333 TCP. `-u` flag makes the server listen on UDP port instead. 

## Example run 

```
➜  nc localhost 3333
Sat, 15 Oct 2016 15:58:59 CEST
```

## Author and licence

Dariusz Dwornikowski, MIT 
