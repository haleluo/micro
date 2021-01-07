package micro

import (
	"fmgo.io/microv2/go-micro/v2/client"
	"fmgo.io/microv2/go-micro/v2/debug/trace"
	"fmgo.io/microv2/go-micro/v2/server"
	"fmgo.io/microv2/go-micro/v2/store"

	// set defaults
	gcli "fmgo.io/microv2/go-micro/v2/client/grpc"
	memTrace "fmgo.io/microv2/go-micro/v2/debug/trace/memory"
	gsrv "fmgo.io/microv2/go-micro/v2/server/grpc"
	memoryStore "fmgo.io/microv2/go-micro/v2/store/memory"
)

func init() {
	// default client
	client.DefaultClient = gcli.NewClient()
	// default server
	server.DefaultServer = gsrv.NewServer()
	// default store
	store.DefaultStore = memoryStore.NewStore()
	// set default trace
	trace.DefaultTracer = memTrace.NewTracer()
}
