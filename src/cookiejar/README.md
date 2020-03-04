# golang exportable cookiejar

The package clones [`net/http/cookiejar`](https://github.com/golang/go/tree/3d5703babe9c5344252db3fb8e96f20cd036535a/src/net/http/cookiejar)
and add several method in a new `exportable.go` to make `cookiejar.Jar` serializable.

It's really simple and easy to use.
