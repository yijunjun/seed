hash=`git rev-parse --short HEAD`
rc=`date "+%Y-%m-%d_%H:%M:%S"`
target=seed
go build -ldflags "-s -w -X seed/config.ConfigArg=build -X main.GitHash=${hash} -X main.CompileTime=${rc}" -o ${target} main.go
chmod a+x ${target}
