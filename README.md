# Godaytime

This is intended as an example for Computer Networks 2 class at Poznań University of Technology. Mostly usable by my students. Godaytime lousy implements Daytime protocol from [RFC867](https://tools.ietf.org/html/rfc867).


## Install

```
go get -u github.com/tdi/godaytime
```

## Usage 
```
godaytime [-h] [-H HOSTNAME] [-p PORT]
```

By default godaytime listens on port localhost:3333. 

## Example run 

```
➜  nc localhost 3333
Sat, 15 Oct 2016 15:58:59 CEST
```

## Author and licence

Dariusz Dwornikowski, MIT 