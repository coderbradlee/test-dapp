export GOROOT=/usr/lib/golang
export GOPATH=`pwd`

rm -fr ballot
go build -o ballot testballot
