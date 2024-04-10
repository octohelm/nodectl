package logger

import (
	"context"
	"sync"
	"time"

	"github.com/dagger/dagger/dagql/idtui"
	"github.com/dagger/dagger/telemetry/sdklog"
	"github.com/muesli/termenv"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"golang.org/x/sync/errgroup"
)

func New() *TUI {
	tui := idtui.New()

	return &TUI{
		Frontend:          tui,
		streamingExporter: newStreamingExporter(),
	}
}

type TUI struct {
	*idtui.Frontend
	streamingExporter *streamingExporter
}

func (t *TUI) Shutdown(c context.Context) error {
	eg, ctx := errgroup.WithContext(c)

	if t.Plain {
		eg.Go(func() error {
			return t.streamingExporter.Shutdown(ctx)
		})
	}

	eg.Go(func() error {
		return t.Frontend.Shutdown(ctx)
	})

	return eg.Wait()
}

func (t *TUI) ExportSpans(ctx context.Context, spans []sdktrace.ReadOnlySpan) error {
	if t.Plain {
		// just collect spanNames
		if err := t.streamingExporter.ExportSpans(ctx, spans); err != nil {
			return err
		}
	}
	return t.Frontend.ExportSpans(ctx, spans)
}

func (t *TUI) ExportLogs(ctx context.Context, logs []*sdklog.LogData) error {
	if t.Plain {
		// streaming output only
		return t.streamingExporter.ExportLogs(ctx, logs)
	}
	return t.Frontend.ExportLogs(ctx, logs)
}

type streamingExporter struct {
	spanNames sync.Map

	// incr idx for spanID
	idx                 int64
	batchPrinterGetters sync.Map
	batchPrinterWg      sync.WaitGroup

	output   *termenv.Output
	outputMu sync.RWMutex

	frameTicker *time.Ticker

	done     chan struct{}
	doneOnce sync.Once
}
