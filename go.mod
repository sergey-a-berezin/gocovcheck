module github.com/sergey-a-berezin/gocovcheck

replace (
    github.com/sergey-a-berezin/gocovcheck/coverage => ./coverage
    github.com/sergey-a-berezin/gocovcheck/jsonread => ./jsonread
)
go 1.12

require (
	github.com/sergey-a-berezin/gocovcheck/coverage v0.0.0-20191027213213-3e618f8d8f3b
	github.com/sergey-a-berezin/gocovcheck/jsonread v0.0.0-20191027213213-3e618f8d8f3b // indirect
	github.com/smartystreets/goconvey v1.6.4
	golang.org/x/lint v0.0.0-20190930215403-16217165b5de // indirect
)
