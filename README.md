# GTerm
This repo is a CLI tool that is attempting to replicate the Linux terminal user experience for Google Drive.

## Environment setup
Follow steps one and two [here](https://developers.google.com/drive/api/v3/quickstart/go)

## To build and run
```bash
go build && ./gterm
```

## Initial challenges (TODO)
1. Figure out how to find id of "My Drive" programatically. This will be the equivilent of the "Home" folder of the user. I can currently find it by making a test file in my drive, searching for that file, and then pulling the parent ID (you can see this in the `foo` function, it's currently hardcoded). This is not replicable. I've tried searching for `kind` of `drive#parentReference` but the API doesn't like this. I've also tried searching for `isRoot` as `true` but it also doesn't seem to support this parameter either. There doesn't seem to be pariety for the `q` param used to search and the properties of the returned objects.
2. After we can determine this initial context basic folder navigation should be simple, we just store the context of where the user is in regards to navigation as the current parent id. Then when we do things like `ls` we just search for the files with that parent.

## Google Documentation
Just as a note, several items in this documentation appear to be incorrect. For instance `name` is actually `Title`. This makes life difficult but you can usually explore the returned objects and figure things out.

### REST APIs
https://developers.google.com/drive/api/v3/about-sdk
### Golang SDK
https://godoc.org/google.golang.org/api/drive/v3