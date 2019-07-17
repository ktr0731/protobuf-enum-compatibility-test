To check backward-compatibility: 
``` bash
$ go run v2/cmd/cmdv2/main.go | go run v1/cmd/cmdv1/main.go
```

`cmdv2` writes out serialized apiv2.Message to stdout and `cmdv1` reads binary from stdin, then encodes it to apiv1.Message.
