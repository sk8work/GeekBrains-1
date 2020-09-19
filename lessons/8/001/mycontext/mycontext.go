package mycontext

import (
	"context"
	"log"
	"net/http"
	"time"
)

func doStuff(ctx context.Context, w http.ResponseWriter) error {
	time.Sleep(time.Second)
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	_, err := w.Write([]byte(`Privet`))
	return err
}

func Handler(w http.ResponseWriter, r *http.Request) {
	err := doStuff(r.Context(), w)
	if err != nil {
		log.Println(err)
	}
}
