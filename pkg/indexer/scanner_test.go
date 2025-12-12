package indexer

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestScanDirsFindsAudioFiles(t *testing.T) {
    base := t.TempDir()
    // create a simple audio file
    mdir := filepath.Join(base, "music")
    if err := os.MkdirAll(mdir, 0o755); err != nil {
        t.Fatalf("mkdir: %v", err)
    }
    f1 := filepath.Join(mdir, "song1.mp3")
    if err := os.WriteFile(f1, []byte("dummy"), 0o644); err != nil {
        t.Fatalf("write: %v", err)
    }
    idx := NewIndexAtBase(base)
    if err := ScanDirs([]string{mdir}, idx, 1); err != nil {
        t.Fatalf("ScanDirs: %v", err)
    }
    tracks := idx.GetAll()
    if len(tracks) != 1 {
        t.Fatalf("expected 1 track, got %d", len(tracks))
    }
    t0 := tracks[0]
    if t0.Path != f1 {
        t.Fatalf("expected path %s got %s", f1, t0.Path)
    }
    if t0.Title == "" {
        t.Fatalf("expected title to be set")
    }
    // scan on nested directory
    if err := os.MkdirAll(filepath.Join(mdir, "sub"), 0o755); err != nil {
        t.Fatalf("mkdir sub: %v", err)
    }
    f2 := filepath.Join(mdir, "sub", "song2.flac")
    if err := os.WriteFile(f2, []byte("dummy"), fs.FileMode(0o644)); err != nil {
        t.Fatalf("write: %v", err)
    }
    if err := ScanDirs([]string{mdir}, idx, 1); err != nil {
        t.Fatalf("ScanDirs: %v", err)
    }
    if len(idx.GetAll()) != 2 {
        t.Fatalf("expected 2 tracks, got %d", len(idx.GetAll()))
    }
}
