### newrepo

Create new Github repositories from the command line

### steps
1. 
```bash 
$ go get -u github.com/waymobetta/newrepo
```
2. generate _personal access token_ with **repo** scoped access [here](https://github.com/settings/tokens)
3. set value for this token with `GITHUB_TOKEN` as the env var name
4. run CLI help screen with `newrepo -h`

**Note**: description flag can be left blank, but name and private flags are required

### usage

```bash
$ newrepo --name=testrepo --private=true`
```

### credits

Basically, I just leveraged the code from the go-github examples [here](https://github.com/google/go-github/blob/master/example/newrepo/main.go), kept the same program name and customized it to make it a little easier to work with and slightly more extensible

MIT
