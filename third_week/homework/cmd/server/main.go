package server

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"golang.org/x/sync/errgroup"
	"jike/third_week/homework/api"
	"jike/third_week/homework/configs"
	"jike/third_week/homework/internal/biz"
	"log"
	"net/http"
	"os"
	"os/signal"
)

//
func main()  {
	var filePath string
	flag.StringVar(&filePath, "c", "config.yml", "")
	flag.Parse()

	cfg, err := configs.LoadConfig(filePath)
	if err != nil {
		log.Fatal("load config fail")
	}
	fmt.Println(cfg)

	group, ctx := errgroup.WithContext(context.Background())
	group.Go(func() error {
		biz.CheckTask(ctx)
		return errors.New("normal exit")
	})

	group.Go(func() error {
		return serverApp(ctx)
	})
	group.Go(func() error {
		return serverDebug(ctx)
	})

	//todo: wire构建依赖

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Kill, os.Interrupt)
	<- sig

	group.Wait()
}

func serverApp(ctx context.Context) error {
	handler := api.NewServerApiInit()
	return server(ctx, ":9202", handler)
}

func serverDebug(ctx context.Context) error {
	return server(ctx, ":9201", http.DefaultServeMux)
}

func server(ctx context.Context, addr string, handler http.Handler) error {
	h:= http.Server{
		Addr:addr,
		Handler: handler,
	}
	go func() {
		<- ctx.Done()
		h.Shutdown(context.Background())
	}()

	err := h.ListenAndServe()
	return err
}
