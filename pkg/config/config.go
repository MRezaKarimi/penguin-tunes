package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Config stores application-level configuration; more groups can be added later
type Config struct {
    SrcDirs []string `json:"srcDirs"`
}

// Manager handles reading/writing config file placed inside given baseDir
type Manager struct {
    path string
    mtx  sync.RWMutex
    cfg  *Config
}

// NewManagerAt creates a new manager that stores config at baseDir/config.json
func NewManagerAt(baseDir string) (*Manager, error) {
    if err := os.MkdirAll(baseDir, 0o755); err != nil {
        return nil, fmt.Errorf("mkdir base dir: %w", err)
    }
    file := filepath.Join(baseDir, "config.json")
    m := &Manager{path: file}
    if err := m.loadOrCreate(); err != nil {
        return nil, err
    }
    return m, nil
}

// NewManager uses user config dir and PenguinTunes subdir
func NewManager() (*Manager, error) {
    dir, err := os.UserConfigDir()
    if err != nil {
        return nil, fmt.Errorf("UserConfigDir: %w", err)
    }
    return NewManagerAt(filepath.Join(dir, "PenguinTunes"))
}

func (m *Manager) loadOrCreate() error {
    m.mtx.Lock()
    defer m.mtx.Unlock()
    if _, err := os.Stat(m.path); err == nil {
        b, err := os.ReadFile(m.path)
        if err != nil {
            return fmt.Errorf("read config: %w", err)
        }
        var cfg Config
        if err := json.Unmarshal(b, &cfg); err != nil {
            return fmt.Errorf("unmarshal config: %w", err)
        }
        m.cfg = &cfg
        return nil
    }
    m.cfg = &Config{SrcDirs: []string{}}
    return m.saveLocked()
}

func (m *Manager) GetConfig() Config {
    m.mtx.RLock()
    defer m.mtx.RUnlock()
    cfg := *m.cfg
    cfg.SrcDirs = append([]string{}, m.cfg.SrcDirs...)
    return cfg
}

func (m *Manager) SaveConfig(cfg Config) error {
    m.mtx.Lock()
    defer m.mtx.Unlock()
    m.cfg = &cfg
    return m.saveLocked()
}

func (m *Manager) saveLocked() error {
    b, err := json.MarshalIndent(m.cfg, "", "  ")
    if err != nil {
        return fmt.Errorf("marshal config: %w", err)
    }
    tmp := m.path + ".tmp"
    if err := os.WriteFile(tmp, b, 0o644); err != nil {
        return fmt.Errorf("write tmp config: %w", err)
    }
    if err := os.Rename(tmp, m.path); err != nil {
        return fmt.Errorf("rename tmp: %w", err)
    }
    return nil
}
