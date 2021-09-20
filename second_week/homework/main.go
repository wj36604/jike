package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type host struct {
	ctx context.Context
	server http.Server
}

func newHost(ctx context.Context) (*host) {
	return &host{
		ctx:ctx,
		server:http.Server{
			Addr: ":9202",
			Handler: http.DefaultServeMux,
		},
	}
}

func main() {
	group, ctx := errgroup.WithContext(context.Background())
	h := newHost(ctx)
	group.Go(h.initHttpServer)
	group.Go(h.listenLinuxSignal)

	err := group.Wait()
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (h *host) initHttpServer() (err error) {
	http.HandleFunc("/help", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("message: *********"))
	})
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("test"))
	})

	err = h.server.ListenAndServe()
	return
}

func (h *host)listenLinuxSignal() (err error) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Interrupt)
	for {
		select {
		case <-h.ctx.Done():
			fmt.Println("exit listen signal")
			return
		case <-sig:
			ctx, _ := context.WithTimeout(context.Background(), time.Millisecond * 20)
			err = h.server.Shutdown(ctx)
			if err != nil {
				fmt.Println(err)
			}
			return
		}
	}
}
