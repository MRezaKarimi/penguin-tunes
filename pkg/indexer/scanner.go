package indexer

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	tag "github.com/dhowden/tag"
)

var audioExtensions = map[string]bool{
    ".mp3": true, ".flac": true, ".m4a": true, ".wav": true, ".ogg": true, ".opus": true, ".aac": true, ".wma": true,
}

func isAudioFile(path string) bool {
    ext := strings.ToLower(filepath.Ext(path))
    return audioExtensions[ext]
}

func idFromPath(path string) string {
    ab, _ := filepath.Abs(path)
    h := sha1.Sum([]byte(ab))
    return hex.EncodeToString(h[:])
}

// readMetadata reads an audio file's tags and returns a Track
func readMetadata(path string, cfgDir string) (*Track, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    m, err := tag.ReadFrom(f)
    t := &Track{ID: idFromPath(path), Path: path}
    if err != nil {
        // return basic info fallback
        t.Title = filepath.Base(path)
        t.Album = "Unknown Album"
        t.Artist = "Unknown Artist"
        return t, nil
    }
    if t.Title = m.Title(); t.Title == "" {
        t.Title = filepath.Base(path)
    }
    if t.Album = m.Album(); t.Album == "" {
        t.Album = "Unknown Album"
    }
    if t.Artist = m.Artist(); t.Artist == "" {
        t.Artist = "Unknown Artist"
    }
    if t.Composer = m.Composer(); t.Composer == "" {
        t.Composer = "Unknown Artist"
    }
    if t.Genre = m.Genre(); t.Genre == "" {
        t.Genre = ""
    }
    rn, _ := m.Track()
    if rn > 0 {
        t.TrackNumber = rn
    }
    t.Year = m.Year()
    p := m.Picture()
    if p != nil {
        ext := ".jpg"
        if strings.Contains(strings.ToLower(p.MIMEType), "png") {
            ext = ".png"
        }
        out := filepath.Join(cfgDir, "covers")
        _ = os.MkdirAll(out, 0o755)
        pth := filepath.Join(out, t.ID+ext)
        if err := os.WriteFile(pth, p.Data, 0o644); err == nil {
            t.Cover = pth
        }
    }
    return t, nil
}

// ScanDirs will scan dirs recursively and update index
func ScanDirs(dirs []string, idx *Index, concurrency int) error {
    if idx == nil {
        return fmt.Errorf("nil index")
    }
    if concurrency <= 0 {
        concurrency = runtime.NumCPU()
    }
    paths := make(chan string, 2048)
    var wg sync.WaitGroup
    for i := 0; i < concurrency; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for p := range paths {
                t, err := readMetadata(p, idx.cfgDir)
                if err != nil {
                    continue
                }
                idx.AddOrUpdateTrack(t)
            }
        }()
    }
    for _, d := range dirs {
        _ = filepath.WalkDir(d, func(path string, de os.DirEntry, walkErr error) error {
            if walkErr != nil {
                return nil
            }
            if de.IsDir() {
                return nil
            }
            if isAudioFile(path) {
                paths <- path
            }
            return nil
        })
    }
    close(paths)
    wg.Wait()
    return idx.SaveToFile()
}
