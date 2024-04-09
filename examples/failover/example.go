package main

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/exp/slog"

	slogmulti "github.com/samber/slog-multi"
)

func main() {
	// ncat -l 1000 -k
	// ncat -l 1001 -k
	// ncat -l 1002 -k

	logstash1, _ := net.Dial("tcp", "localhost:1000")
	logstash2, _ := net.Dial("tcp", "localhost:1001")
	logstash3, _ := net.Dial("tcp", "localhost:1002")

	logger := slog.New(
		slogmulti.Failover()(
			slog.NewJSONHandler(logstash1, &slog.HandlerOptions{}),
			slog.NewJSONHandler(logstash2, &slog.HandlerOptions{}),
			slog.NewJSONHandler(logstash3, &slog.HandlerOptions{}),
		),
	)

	logger.
		With(
			slog.Group("user",
				slog.String("id", "user-123"),
				slog.Time("created_at", time.Now().AddDate(0, 0, -1)),
			),
		).
		With("environment", "dev").
		With("error", fmt.Errorf("an error")).
		Error("A message")
}
