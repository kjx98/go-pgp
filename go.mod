module github.com/kjx98/go-pgp

go 1.11

require (
	github.com/kjx98/openpgp v0.0.0-20191125135020-168fae501ad1
	golang.org/x/crypto v0.0.0
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190513172903-22d7a77e9e5f
	golang.org/x/net => github.com/golang/net v0.0.0-20190404232315-eb5bcb51f2a3
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190412213103-97732733099d
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
