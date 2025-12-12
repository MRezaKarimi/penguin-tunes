package indexer

import (
	"context"
	"os"
	"path/filepath"
	cfg "penguin-tunes/pkg/config"
	"testing"
	"time"
)

type fakeEmitter struct{
    ch chan any
}

func (f *fakeEmitter) Emit(ctx context.Context, event string, data any) {
    // send sentinel
    select {
    case f.ch <- data:
    default:
    }
}

func TestWatcherEmitsOnCreate(t *testing.T) {
    base := t.TempDir()
    m, err := cfg.NewManagerAt(base)
    if err != nil {
        t.Fatalf("NewManagerAt: %v", err)
    }
    // set srcDirs
    cfg := m.GetConfig()
    music := filepath.Join(base, "music")
    if err := os.MkdirAll(music, 0o755); err != nil {
        t.Fatalf("mkdir: %v", err)
    }
    cfg.SrcDirs = []string{music}
    if err := m.SaveConfig(cfg); err != nil {
        t.Fatalf("SaveConfig: %v", err)
    }
    idx := NewIndexAtBase(base)
    fe := &fakeEmitter{ch: make(chan any, 1)}
    wa, err := NewWatcher(context.Background(), idx, m, fe)
    if err != nil {
        t.Fatalf("NewWatcher: %v", err)
    }
    defer wa.Close()
    _, err = wa.Start()
    if err != nil {
        t.Fatalf("Start: %v", err)
    }
    // create a new audio file
    f1 := filepath.Join(music, "new.mp3")
    if err := os.WriteFile(f1, []byte("dummy"), 0o644); err != nil {
        t.Fatalf("write: %v", err)
    }
    // Wait up to 3s for event
    select {
    case <-fe.ch:
        // ok
    case <-time.After(3 * time.Second):
        t.Fatalf("expected event, timed out")
    }
}
