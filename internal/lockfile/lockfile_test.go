package lockfile

import (
	"os"
	"testing"
)

func TestAddAndRetrieve(t *testing.T) {
	lf := New()

	lf.Add(Package{
		Name:    "cobra",
		Version: "1.10.2",
		Source:  "github.com/spf13/cobra",
	})

	pkg, exists := lf.Packages["cobra"]
	if !exists {
		t.Fatal("expected cobra to be installed, but it wasn't found")
	}

	if pkg.Version != "1.10.2" {
		t.Errorf("expected version 1.10.2, got %s", pkg.Version)
	}
}

func TestRemovePackage(t *testing.T) {
	lf := New()
	lf.Add(Package{Name: "gin", Version: "1.9.1", Source: "github.com/gin-gonic/gin"})

	removed := lf.Remove("gin")
	if !removed {
		t.Fatal("expected Remove to return true, got false")
	}

	if _, exists := lf.Packages["gin"]; exists {
		t.Fatal("expected gin to be removed, but it still exists")
	}
}

func TestRemoveNonExistent(t *testing.T) {
	lf := New()
	removed := lf.Remove("nonexistent")
	if removed {
		t.Fatal("expected Remove to return false for nonexistent package")
	}
}

func TestSaveAndLoad(t *testing.T) {
	path := "test_pkgr.lock"
	defer os.Remove(path)

	lf := New()
	lf.Add(Package{Name: "viper", Version: "1.18.0", Source: "github.com/spf13/viper"})

	if err := lf.Save(path); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	pkg, exists := loaded.Packages["viper"]
	if !exists {
		t.Fatal("expected viper in loaded lockfile")
	}

	if pkg.Version != "1.18.0" {
		t.Errorf("expected version 1.18.0, got %s", pkg.Version)
	}
}
