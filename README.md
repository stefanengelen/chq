# chq
stream graph based MPEG TS analyser

this project aims to create a ts analyser that is capable of processing large amounts of data quickly, both in real time and offline.

it achieves this by separating the analysis into small "Nodes", each of which does a simple task (e.g. remap pids, count packets, look for CC errors). The user then combies these in a .chq file so the processing is focused on what is actually required.

each node runs in its own goroutine, so the pipelines created are concurrent with each other. this means that several different types of analysis can be done in parallel on one stream.

To get started, run go get:

    go get github.com/aktungmak/chq

this should give you a `chq` executable, if not run `go build` to get it. running chq with no options, it will read the file `basic.chq`, and start processing based on the configuration specified.

    #node inp ::= NetInput 127.0.0.1 5501
    node inp ::= FileInput mpts.ts
    node ctr ::= PidCounter
    node ccc ::= CcCommenter
    node cwr ::= CommentWriter comments.txt
    node out ::= FileWriter o.ts
    node pat ::= PatParser 0
    node rmp ::= PidRemapper 310 433
    node brc ::= PcrBrCommenter 310


    conn inp -> brc -> cwr -> ctr -> rmp -> out

    conn inp -> pat

it will also start up an http server on http://localhost:10101, if you point your browser there you can see some JSON of the analysis in real time.

Check out `Server.go` for more endpoints.
