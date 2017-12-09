# Agenda

A microservice powered by Apiary, containing a local cli program and a server. APIs designed on Apiary, see [Apiary Documentation] for details.

# Getting Started

Install the local program designed by Go programming language or simply

    $ cd cli
    $ go run main.go

# Usage

Agenda is a local cli program. You can use the program to access,
edit and manage your meetings. All the information is stored on our cloud
servers safely. To get started, register your own accounts or just login.

Usage:

    agenda [command]

Available Commands:

    clearmeeting  Remove all the meetings you host
    createmeeting Create a meeting
    destroy       Destroy your account and logout
    help          Help about any command
    login         Login to the agenda server
    logout        Logout from your account
    meetings      Show all meetings
    quitmeeting   TODO
    register      Register a new account
    removemeeting TODO
    users         Show all users

Flags:

        --config string   config file (default is $HOME/.cli.yaml)
    -h, --help            help for agenda
    -t, --toggle          Help message for toggle

Use "agenda [command] --help" for more information about a command.

# Test

This is a practice program. Testing its performance is one of the key points.

## Working with docker to serve

//TODO

## Go Test

By using mock server, we can easily test our local program. By executing the `cmd_test.go`, we have the result below

    =====> In TEST of Register
    Register successfully. You are logged in as
    map[phone:12233344445 id:1 key:1e3576bt username:zhang3 password:zhang email:zhang3@mail2.sysu.edu.cn]
    =====> In TEST of Login
    Logout
    Login successfully. Your api key is
    1e3576bt
    =====> In TEST of Showing all users
    [
        {
            "id":1,
            "username":"zhang3",
            "password":"zhang",
            "email":"zhang3@mail2.sysu.edu.cn",
            "phone":"12233334444"
        }, {
            "id":2,
            "username":"li4",
            "password":"li",
            "email":"li4@mail2.sysu.edu.cn",
            "phone":"12233334445"
        }
    ]
    =====> In TEST of Creating a new meeting
    Meeting 'testMeeting' created
    =====> In TEST of Showing all meetings
    [
        {
            "id":1,
            "title":"zhang3",
            "host":"zhang3",
            "members":["li4"],
            "starttime":"2006/01/02/15:04",
            "endtime":"2006/01/02/15:05"
        },{
            "id":2,
            "title":"li4",
            "host":"li4",
            "members":["zhang3"],
            "starttime":"2006/01/03/15:04",
            "endtime":"2006/01/03/15:05"
        }
    ]
    =====> In TEST of clearing all meetings
    Your meetings have been removed.
    =====> In TEST of destroying account
    Your account has been removed
    Please login first.
    PASS
    ok  	github.com/HinanawiTenshi/agenda/cli/cmd	33.747s

[Apiary Documentation]: https://agenda15.docs.apiary.io/
