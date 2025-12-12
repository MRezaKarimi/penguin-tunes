package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	goruntime "runtime"

	cfg "penguin-tunes/pkg/config"
	"penguin-tunes/pkg/indexer"

	wailsruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx        context.Context
	cfgManager *cfg.Manager
	idx        *indexer.Index
	watcher    *indexer.Watcher
}

// wailsEmitter adapts Wails runtime to indexer.EventEmitter
type wailsEmitter struct{}

func (w wailsEmitter) Emit(ctx context.Context, event string, data any) {
	wailsruntime.EventsEmit(ctx, event, data)
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// Initialize config manager
	cm, err := cfg.NewManager()
	if err != nil {
		fmt.Printf("config manager error: %v\n", err)
		return
	}
	a.cfgManager = cm
	// determine index path
	cfgDir, _ := os.UserConfigDir()
	appDir := filepath.Join(cfgDir, "PenguinTunes")
	idxPath := filepath.Join(appDir, "index.json")
	a.idx = indexer.NewIndex(idxPath, appDir)
	if err := a.idx.LoadFromFile(); err != nil {
		fmt.Printf("load index error: %v\n", err)
	}
	// Watcher
	wa, err := indexer.NewWatcher(a.ctx, a.idx, a.cfgManager, wailsEmitter{})
	if err != nil {
		fmt.Printf("watcher error: %v\n", err)
	} else {
		a.watcher = wa
		stop, _ := wa.Start()
		// for now, stop is not saved; we'll keep it running until app exit
		_ = stop
	}
	// Start initial scan in background
	// Emit current index (if any) so frontend can display it instantly
	wailsruntime.EventsEmit(a.ctx, "index-updated", a.idx.GetAll())
	go func() {
		cfg := cm.GetConfig()
		if len(cfg.SrcDirs) > 0 {
			// scan
			if err := indexer.ScanDirs(cfg.SrcDirs, a.idx, goruntime.NumCPU()); err != nil {
				fmt.Printf("scan error: %v\n", err)
			}
			// emit event to frontend
			wailsruntime.EventsEmit(a.ctx, "index-updated", a.idx.GetAll())
		}
	}()
}

// GetConfig returns current configuration
func (a *App) GetConfig() (cfg.Config, error) {
	if a.cfgManager == nil {
		return cfg.Config{}, fmt.Errorf("config manager not initialized")
	}
	return a.cfgManager.GetConfig(), nil
}

// SaveConfig updates and persists config, restarting watchers/scans if changed
func (a *App) SaveConfig(cfg cfg.Config) error {
	if a.cfgManager == nil {
		return fmt.Errorf("not initialized")
	}
	if err := a.cfgManager.SaveConfig(cfg); err != nil {
		return err
	}
	// When srcDirs change, restart scan and watchers
	go func() {
		// Kick off a scan
		if a.idx != nil {
			if err := indexer.ScanDirs(cfg.SrcDirs, a.idx, goruntime.NumCPU()); err != nil {
				fmt.Printf("scan error: %v\n", err)
			}
			wailsruntime.EventsEmit(a.ctx, "index-updated", a.idx.GetAll())
		}
	}()
	// restart watchers to pick new srcDirs
	if a.watcher != nil {
		_ = a.watcher.Close()
	}
	// create new watcher and add watches
	wa, err := indexer.NewWatcher(a.ctx, a.idx, a.cfgManager, wailsEmitter{})
	if err == nil {
		a.watcher = wa
		_, _ = wa.Start()
	}
	return nil
}

// GetTracks returns all indexed track metadata
func (a *App) GetTracks() ([]*indexer.Track, error) {
	if a.idx == nil {
		return nil, fmt.Errorf("index not initialized")
	}
	return a.idx.GetAll(), nil
}

// AddSrcDir adds a directory to srcDirs and persists it
func (a *App) AddSrcDir(dir string) error {
	if a.cfgManager == nil {
		return fmt.Errorf("not initialized")
	}
	cfg := a.cfgManager.GetConfig()
	for _, d := range cfg.SrcDirs {
		if d == dir {
			return nil // already exists
		}
	}
	cfg.SrcDirs = append(cfg.SrcDirs, dir)
	return a.SaveConfig(cfg)
}

// RemoveSrcDir removes a directory from srcDirs and persists it
func (a *App) RemoveSrcDir(dir string) error {
	if a.cfgManager == nil {
		return fmt.Errorf("not initialized")
	}
	cfg := a.cfgManager.GetConfig()
	n := make([]string, 0, len(cfg.SrcDirs))
	for _, d := range cfg.SrcDirs {
		if d == dir {
			continue
		}
		n = append(n, d)
	}
	cfg.SrcDirs = n
	return a.SaveConfig(cfg)
}
