# Agenda

[![Build Status](https://travis-ci.org/HinanawiTenshi/agenda.svg?branch=master)](https://travis-ci.org/HinanawiTenshi/agenda)

A microservice powered by Apiary, containing a local cli program and a server. APIs designed on Apiary, see [Apiary Documentation] for details.

The main purpose of this repository is to learning and practising. It mainly focuses on the construction of a microservice. Thus, a lot of tests will be displayed here.

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

1. Pull from docker hub

        # docker pull hinanawitenshi/agenda
        Using default tag: latest
        latest: Pulling from hinanawitenshi/agenda
        85b1f47fba49: Already exists
        ba6bd283713a: Already exists
        817c8cd48a09: Already exists
        cca33291b6d1: Already exists
        aee79ff31530: Already exists
        a4877a39ec8e: Already exists
        be4367dbb932: Already exists
        b4a451904d17: Pull complete
        10fc1817c0f7: Pull complete
        3aac60ae16d3: Pull complete
        d8273def28b5: Pull complete
        e8f5d3283b0a: Pull complete
        Digest: sha256:caf20b3d5a91ddb47a3b403935cab3643264a552c38664ac6dad37a545077fd7
        Status: Downloaded newer image for hinanawitenshi/agenda:latest

        # docker images
        REPOSITORY              TAG                 IMAGE ID            CREATED             SIZE
        hinanawitenshi/agenda   latest              faf415974e8b        20 hours ago        819MB
        golang                  1.8                 a8ef0d2260ca        5 weeks ago         712MB
        mysql                   latest              5709795eeffa        5 weeks ago         408MB
        hello-world             latest              725dcfab7d63        5 weeks ago         1.84kB
        golang                  onbuild             5d82e356477f        4 months ago        699MB

2. Run the server(To keep the databases, remove the `--rm` flag)

        # docker run --rm -d --name agenda -p 8080:8080 hinanawitenshi/agenda
        fb4ef49b31cb89389a6ec6964227006c7d14452e863273f068b642e6a0b7dcf3

3. Use cli to access the server

    - Register

            # agenda register -u admin -p admin -e empty@email.com -o 12345678901
            Register successfully. You are logged in as
            {
              "id": 1,
              "key": "21232f297a57a5a743894a0e4a801fc3",
              "username": "admin",
              "password": "admin",
              "email": "empty@email.com",
              "phone": "12345678901"
            }

    - Query users

            # agenda users
            [
              {
                "id": 1,
                "key": "******",
                "username": "admin",
                "password": "******",
                "email": "empty@email.com",
                "phone": "12345678901"
              }
            ]

    - Create a new meeting

            # agenda createmeeting -t meeting1 -m a,b,c -s 2017/10/25/13:00 -e 2017/10/25/15:00
            Meeting 'meeting1' created

    - Query meetings

            # agenda meetings
            [
              {
                "id": 1,
                "title": "meeting1",
                "host": "admin",
                "members": [
                  "a",
                  "b",
                  "c"
                ],
                "starttime": "2017/10/25/13:00",
                "endtime": "2017/10/25/15:00"
              }
            ]



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

## ab Test

Also, ab test is necessary to test the performace of our server

    # ab -n 10000 -c 1000 http://localhost:8080/v1/meetings?21232f297a57a5a743894a0e4a801fc3
    This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
    Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
    Licensed to The Apache Software Foundation, http://www.apache.org/

    Benchmarking localhost (be patient)
    Completed 1000 requests
    Completed 2000 requests
    Completed 3000 requests
    Completed 4000 requests
    Completed 5000 requests
    Completed 6000 requests
    Completed 7000 requests
    Completed 8000 requests
    Completed 9000 requests
    Completed 10000 requests
    Finished 10000 requests


    Server Software:
    Server Hostname:        localhost
    Server Port:            8080

    Document Path:          /v1/meetings?21232f297a57a5a743894a0e4a801fc3
    Document Length:        0 bytes

    Concurrency Level:      1000
    Time taken for tests:   2.064 seconds
    Complete requests:      10000
    Failed requests:        0
    Non-2xx responses:      10000
    Total transferred:      1230000 bytes
    HTML transferred:       0 bytes
    Requests per second:    4845.14 [#/sec] (mean)
    Time per request:       206.392 [ms] (mean)
    Time per request:       0.206 [ms] (mean, across all concurrent requests)
    Transfer rate:          581.98 [Kbytes/sec] received

    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        0   27 162.3      1    1029
    Processing:     0   22  92.4     11    1030
    Waiting:        0   21  92.5     10    1030
    Total:          1   49 207.3     12    2057

    Percentage of the requests served within a certain time (ms)
      50%     12
      66%     15
      75%     17
      80%     19
      90%     27
      95%     46
      98%   1042
      99%   1049
     100%   2057 (longest request)

[Apiary Documentation]: https://agenda15.docs.apiary.io/
