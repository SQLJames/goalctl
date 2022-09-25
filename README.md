
# goalctl
[![Gosec](https://github.com/SQLJames/goalctl/actions/workflows/gosec.yml/badge.svg?branch=main)](https://github.com/SQLJames/goalctl/actions/workflows/gosec.yml)
[![govulncheck](https://github.com/SQLJames/goalctl/actions/workflows/govulncheck.yml/badge.svg?branch=main)](https://github.com/SQLJames/goalctl/actions/workflows/govulncheck.yml)

Is a basic journaling program written in golang.
# Getting started
To make it easier to use the tool, we have built help commands onto every command. Running just the binary by itself will show the help information for the whole project.
You can also run the export command to export all the information to a format of your choice, incase you want to move to a different engine.

# Backend
## storage
The storage engine is running on an embedded sqlite https://github.com/glebarez/go-sqlite
This data is stored in your homedirectory under the goalctl folder

## Build 
To build the application we are using mage in a zero install model. You can see some of the common commands in the makefile.
SQLC is handling our database code generation, so if there are any updates to schema or queries you will need to run the `sqlc generate`
command on the folder `pkg\storage\sqlite\sqlc`

# Example
```
.\bin\goalctl.exe create notebook -name "work"
.\bin\goalctl.exe create entry -entry "Entry data for the log" -name "work" -tag "1" -tag "2"  
.\bin\goalctl.exe create goal --duedate "2023-08-13" --entry "Celebrate" --name "birthday2" --priority 1
.\bin\goalctl.exe link -g 1 -le 2
.\bin\goalctl.exe list goal
```
