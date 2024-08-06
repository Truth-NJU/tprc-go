module local

go 1.12

replace github.com/Truth-NJU/tprc-go/trpcprotocol/local => ./stub/github.com/Truth-NJU/tprc-go/trpcprotocol/local

require (
	git.code.oa.com/trpc-go/trpc-filter/debuglog v0.1.12
	git.code.oa.com/trpc-go/trpc-filter/recovery v0.1.4
	git.code.oa.com/trpc-go/trpc-go v0.18.3
	github.com/Truth-NJU/tprc-go/trpcprotocol/local v0.0.0-00010101000000-000000000000
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.3.0
)
