package indexer

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

// Track holds metadata for a single audio file
type Track struct {
    ID          string `json:"id"`
    Path        string `json:"path"`
    Title       string `json:"title"`
    Album       string `json:"album"`
    Artist      string `json:"artist"`
    Composer    string `json:"composer"`
    Genre       string `json:"genre"`
    TrackNumber int    `json:"track_number"`
    Cover       string `json:"cover"`
    Year        int    `json:"year"`
}

// Index stores tracks keyed by path for quick lookups
type Index struct {
    mtx    sync.RWMutex
    Tracks map[string]*Track `json:"tracks"`
    path   string
    cfgDir string
}

// NewIndex creates a new index manager at path with cfgDir for covers
func NewIndex(path string, cfgDir string) *Index {
    return &Index{Tracks: make(map[string]*Track), path: path, cfgDir: cfgDir}
}

// NewIndexAtBase constructs an index file path under baseDir/index.json
func NewIndexAtBase(baseDir string) *Index {
    return NewIndex(filepath.Join(baseDir, "index.json"), baseDir)
}

// LoadFromFile loads index from disk if exists
func (idx *Index) LoadFromFile() error {
    idx.mtx.Lock()
    defer idx.mtx.Unlock()
    if _, err := os.Stat(idx.path); err != nil {
        // no index file
        return nil
    }
    b, err := os.ReadFile(idx.path)
    if err != nil {
        return fmt.Errorf("read index: %w", err)
    }
    var wrapper struct{ Tracks map[string]*Track `json:"tracks"` }
    if err := json.Unmarshal(b, &wrapper); err != nil {
        return fmt.Errorf("unmarshal index: %w", err)
    }
    idx.Tracks = wrapper.Tracks
    if idx.Tracks == nil {
        idx.Tracks = make(map[string]*Track)
    }
    return nil
}

// SaveToFile saves index atomically to disk
func (idx *Index) SaveToFile() error {
    idx.mtx.RLock()
    defer idx.mtx.RUnlock()
    wrapper := struct{ Tracks map[string]*Track `json:"tracks"` }{Tracks: idx.Tracks}
    b, err := json.MarshalIndent(wrapper, "", "  ")
    if err != nil {
        return fmt.Errorf("marshal index: %w", err)
    }
    tmp := idx.path + ".tmp"
    if err := os.WriteFile(tmp, b, 0o644); err != nil {
        return fmt.Errorf("write tmp index: %w", err)
    }
    if err := os.Rename(tmp, idx.path); err != nil {
        return fmt.Errorf("rename tmp index: %w", err)
    }
    return nil
}

// AddOrUpdateTrack adds or updates a track in the index
func (idx *Index) AddOrUpdateTrack(t *Track) {
    idx.mtx.Lock()
    defer idx.mtx.Unlock()
    idx.Tracks[t.Path] = t
}

// RemoveTrack removes a track from index
func (idx *Index) RemoveTrack(path string) {
    idx.mtx.Lock()
    defer idx.mtx.Unlock()
    delete(idx.Tracks, path)
}

// GetAll returns a copy of all tracks
func (idx *Index) GetAll() []*Track {
    idx.mtx.RLock()
    defer idx.mtx.RUnlock()
    tracks := make([]*Track, 0, len(idx.Tracks))
    for _, v := range idx.Tracks {
        tracks = append(tracks, v)
    }
    return tracks
}

// SaveCover writes given image bytes to cover path and returns the relative cover path
func (idx *Index) SaveCover(id string, mime string, r io.Reader) (string, error) {
    covers := filepath.Join(idx.cfgDir, "covers")
    if err := os.MkdirAll(covers, 0o755); err != nil {
        return "", err
    }
    ext := ".jpg"
    if mime == "image/png" {
        ext = ".png"
    }
    outPath := filepath.Join(covers, id+ext)
    f, err := os.Create(outPath)
    if err != nil {
        return "", err
    }
    defer f.Close()
    if _, err := io.Copy(f, r); err != nil {
        return "", err
    }
    return outPath, nil
}
