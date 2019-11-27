# kubecuddle
kubecuddle k8s operator


### Installing the Operator SDK CLI

```bash
$ GO111MODULE=on go get -d github.com/operator-framework/operator-sdk
go: finding github.com/operator-framework/operator-sdk v0.12.0
go: downloading github.com/operator-framework/operator-sdk v0.12.0
go: extracting github.com/operator-framework/operator-sdk v0.12.0
go get: github.com/operator-framework/operator-sdk@v0.12.0 requires
	k8s.io/api@v0.0.0: reading k8s.io/api/go.mod at revision v0.0.0: unknown revision v0.0.0

$ cd $GOPATH/src/github.com/operator-framework/operator-sdk

~/go/src/github.com/operator-framework/operator-sdk$ git checkout master

~/go/src/github.com/operator-framework/operator-sdk$ make tidy

~/go/src/github.com/operator-framework/operator-sdk$ make install
```

Try `$ operator-sdk`:
```bash
$ operator-sdk 
An SDK for building operators with ease

Usage:
  operator-sdk [command]

Available Commands:
  add         Adds a controller or resource to the project
  alpha       Run an alpha subcommand
  build       Compiles code and builds artifacts
  completion  Generators for shell completions
  generate    Invokes specific generator
  help        Help about any command
  migrate     Adds source code to an operator
  new         Creates a new operator application
  olm-catalog Invokes a olm-catalog command
  print-deps  Print Golang packages and versions required to run the operator
  run         Runs a generic operator
  scorecard   Run scorecard tests
  test        Tests the operator
  up          Launches the operator
  version     Prints the version of operator-sdk

Flags:
  -h, --help      help for operator-sdk
      --verbose   Enable verbose logging

Use "operator-sdk [command] --help" for more information about a command.

```
---

1. Creating operator project by name `tgik-operator`:
```bash
$ cd ~/tanveer/k8s/kubecuddle/
~/tanveer/k8s/kubecuddle$ operator-sdk new tgik-operator --repo github.com/tanalam2411/kubecuddle
INFO[0000] Creating new Go operator 'tgik-operator'.    
INFO[0000] Created go.mod                               
INFO[0000] Created tools.go                             
INFO[0000] Created cmd/manager/main.go                  
INFO[0000] Created build/Dockerfile                     
INFO[0000] Created build/bin/entrypoint                 
INFO[0000] Created build/bin/user_setup                 
INFO[0000] Created deploy/service_account.yaml          
INFO[0000] Created deploy/role.yaml                     
INFO[0000] Created deploy/role_binding.yaml             
INFO[0000] Created deploy/operator.yaml                 
INFO[0000] Created pkg/apis/apis.go                     
INFO[0000] Created pkg/controller/controller.go         
INFO[0000] Created version/version.go                   
INFO[0000] Created .gitignore                           
INFO[0000] Validating project                           
go: finding github.com/operator-framework/operator-sdk master
go: downloading github.com/operator-framework/operator-sdk v0.12.1-0.20191127152849-96dd1d2d5de1
go: extracting github.com/operator-framework/operator-sdk v0.12.1-0.20191127152849-96dd1d2d5de1
INFO[0049] Project validation successful.               
INFO[0049] Project creation complete.   


~/tanveer/k8s/kubecuddle$ tree
.
├── LICENSE
├── README.md
└── tgik-operator
    ├── build
    │   ├── bin
    │   │   ├── entrypoint
    │   │   └── user_setup
    │   └── Dockerfile
    ├── cmd
    │   └── manager
    │       └── main.go
    ├── deploy
    │   ├── operator.yaml
    │   ├── role_binding.yaml
    │   ├── role.yaml
    │   └── service_account.yaml
    ├── go.mod
    ├── go.sum
    ├── pkg
    │   ├── apis
    │   │   └── apis.go
    │   └── controller
    │       └── controller.go
    ├── tools.go
    └── version
        └── version.go

10 directories, 16 files
       
```
