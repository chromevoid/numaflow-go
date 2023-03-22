package main

import (
	"context"
	"strconv"
	"sync"

	functionsdk "github.com/numaproj/numaflow-go/pkg/function"
	"github.com/numaproj/numaflow-go/pkg/function/server"
)

type threadSafeInt struct {
	val int
	sync.RWMutex
}

func (t *threadSafeInt) incrementAndGet() int {
	t.Lock()
	defer t.Unlock()
	if t.val == 10 {
		t.val = 0
	}
	t.val += 1
	return t.val
}

var genKey = threadSafeInt{}

func handle(_ context.Context, key string, d functionsdk.Datum) functionsdk.Messages {
	msg := d.Value()
	outputKey := genKey.incrementAndGet()
	return functionsdk.MessagesBuilder().Append(functionsdk.MessageTo(strconv.Itoa(outputKey), msg))
}

func main() {
	server.New().RegisterMapper(functionsdk.MapFunc(handle)).Start(context.Background())
}
