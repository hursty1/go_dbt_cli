# DBT Helper

Helper application to automate some DBT Cloud Functions

## Installing

go install github.com/hursty1/go_dbt_cli/cmd/dbthelper@latest

or 
go build /cmd/dbthelper
and place the artifact in a folder in your system path


## Features

Basic CLI tool to help perform common DBT Cloud CLI functions

Usage: dbthelper.exe [options] <action> <projectName>

Actions:
    activate        Activates the projectName

Options:
    - l, --list     Display all configured projects
    - a, --active   Display the current active project