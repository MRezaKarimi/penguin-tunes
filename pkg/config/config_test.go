package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestManagerAtCreatesConfig(t *testing.T) {
    base := t.TempDir()
    m, err := NewManagerAt(base)
    if err != nil {
        t.Fatalf("NewManagerAt: %v", err)
    }
    cfg := m.GetConfig()
    if cfg.SrcDirs == nil {
        t.Fatalf("expected srcDirs to be non-nil")
    }
    // Save and reload
    cfg.SrcDirs = []string{"/tmp/music"}
    if err := m.SaveConfig(cfg); err != nil {
        t.Fatalf("SaveConfig: %v", err)
    }
    // ensure file exists
    fn := filepath.Join(base, "config.json")
    if _, err := os.Stat(fn); err != nil {
        t.Fatalf("config file missing: %v", err)
    }
}
