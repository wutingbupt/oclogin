This project is used to learn the go language and 
also create a simple cli to help me switch the different openshift test environment

How to use it?

Build the project
go build

Run the generated binary oclogin

This script is used to switch the openshift environment easier. For example:

oclogin env-1 or oclogin env-2.

Usage:
oclogin [command]

Available Commands:
completion  generate the autocompletion script for the specified shell
context     Switch openshift config context by id
help        Help about any command
init        Init the script environment, and create folder with an sample config file
list        list available openshift environments in your config
login       login one openshift environments in your config

Flags:
--config string   config file (default is $HOME/.oclogin.yaml)
-h, --help            help for oclogin
-t, --toggle          Help message for toggle

Use "oclogin [command] --help" for more information about a command.
