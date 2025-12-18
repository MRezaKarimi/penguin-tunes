package indexer

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"

	cfg "penguin-tunes/pkg/config"
)

// EventEmitter receives emitted events from the package (e.g. index updates)
type EventEmitter interface {
    Emit(ctx context.Context, event string, data any)
}

// Watcher watches directories and updates index on changes
type Watcher struct {
    ctx       context.Context
    w         *fsnotify.Watcher
    idx       *Index
    cm        *cfg.Manager
    closed    bool
    mtx       sync.Mutex
    saveMtx   sync.Mutex
    saveTimer *time.Timer
    emitter   EventEmitter
}

// NewWatcher creates a new Watcher; ctx may be used by emitter
func NewWatcher(ctx context.Context, idx *Index, cm *cfg.Manager, emitter EventEmitter) (*Watcher, error) {
    w, err := fsnotify.NewWatcher()
    if err != nil {
        return nil, err
    }
    return &Watcher{w: w, idx: idx, cm: cm, ctx: ctx, emitter: emitter}, nil
}

// Start starts the watching loop, returns stop func
func (wa *Watcher) Start() (func(), error) {
    cfg := wa.cm.GetConfig()
    for _, d := range cfg.SrcDirs {
        if err := wa.addWatchesRecursive(d); err != nil {
            fmt.Printf("watch add error: %v\n", err)
        }
    }
    stop := make(chan struct{})
    go func() {
        for {
            select {
            case ev, ok := <-wa.w.Events:
                if !ok {
                    return
                }
                wa.handleEvent(ev)
            case err, ok := <-wa.w.Errors:
                if !ok {
                    return
                }
                fmt.Printf("fsnotify error: %v\n", err)
            case <-stop:
                wa.w.Close()
                return
            }
        }
    }()
    closer := func() { close(stop) }
    return closer, nil
}

func (wa *Watcher) addWatchesRecursive(root string) error {
    return filepath.WalkDir(root, func(path string, de os.DirEntry, walkErr error) error {
        if walkErr != nil {
            return nil
        }
        if !de.IsDir() {
            return nil
        }
        if err := wa.w.Add(path); err != nil {
            return err
        }
        return nil
    })
}

func (wa *Watcher) handleEvent(ev fsnotify.Event) {
    p := ev.Name
    if ev.Op&fsnotify.Create == fsnotify.Create {
        if isDir(p) {
            wa.w.Add(p) // best-effort
            return
        }
        if isAudioFile(p) {
            t, err := readMetadata(p, wa.idx.cfgDir)
            if err == nil {
                wa.idx.AddOrUpdateTrack(t)
                wa.scheduleSave(1 * time.Second)
            }
        }
    }
    if ev.Op&fsnotify.Write == fsnotify.Write {
        if isAudioFile(p) {
            t, err := readMetadata(p, wa.idx.cfgDir)
            if err == nil {
                wa.idx.AddOrUpdateTrack(t)
                wa.scheduleSave(1 * time.Second)
            }
        }
    }
    if ev.Op&(fsnotify.Remove|fsnotify.Rename) != 0 {
        if !isDir(p) && isAudioFile(p) {
            wa.idx.RemoveTrack(p)
            wa.scheduleSave(1 * time.Second)
        }
    }
}

func (wa *Watcher) scheduleSave(d time.Duration) {
    wa.saveMtx.Lock()
    defer wa.saveMtx.Unlock()
    if wa.saveTimer != nil {
        wa.saveTimer.Stop()
    }
    wa.saveTimer = time.AfterFunc(d, func() {
        if err := wa.idx.SaveToFile(); err != nil {
            fmt.Printf("index save error: %v\n", err)
        }
        if wa.emitter != nil {
            wa.emitter.Emit(wa.ctx, "index-updated", nil)
        }
    })
}

func (wa *Watcher) Close() error {
    wa.mtx.Lock()
    defer wa.mtx.Unlock()
    if wa.closed {
        return nil
    }
    wa.closed = true
    if wa.saveTimer != nil {
        wa.saveTimer.Stop()
    }
    return wa.w.Close()
}

func isDir(path string) bool {
    fi, err := os.Stat(path)
    if err != nil {
        return false
    }
    return fi.IsDir()
}
