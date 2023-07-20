# What this is
This is a quick proof of concept using golang to play with the instagram API and query it.
## How to run & test project? 

Install Dependencies

<i>Go/Dep install steps will be different for pc</i>
```
#if you need homebrew
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
#else 
brew update

brew install go
brew install dep
brew install fswatch
```

Clone the project
``` 
cd $GOPATH
cd go/src
open .
```

In the opened window put the unzipped folder here. You could also mv the folder from your downloads here.

Install/Update vendor files
``` 
#back in terminal
cd dep-ex
dep ensure
```

Run project w/ file watching 
```
make dev
# go to http://localhost:8080
```

Run project w/ out file watching
```
go run cmd/main.go
# go to http://localhost:8080
```

Test project
```
make test
```

## Overview
From the form on page http://localhost:8080/ displays their info on the page POST http://localhost:8080/instagram.

Info displayed:
* The tag that was searched.
* 5 recent pictures with captions.
