# kick

Install gopherjs.
```
go get -u github.com/gopherjs/gopherjs
```

Install kick.
```
go get -u go.isomorphicgo.org/go/kick
```

Automate building gopherjs and starting a server.
```
kick --appPath=$MEDIUM_APP_ROOT --gopherjsAppPath=$MEDIUM_APP_ROOT/client --mainSourceFile=medium.go
```
