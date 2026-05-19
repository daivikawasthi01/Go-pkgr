package lockfile

import (
	"encoding/json"
	"os"
)

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Source  string `json:"source"`
}

type Lockfile struct {
	Packages map[string]Package `json:"packages"`
}

func New() *Lockfile {
	return &Lockfile{
		Packages: make(map[string]Package),
	}
}

func Load(path string) (*Lockfile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return New(), nil
		}
		return nil, err
	}

	var lf Lockfile
	if err := json.Unmarshal(data, &lf); err != nil {
		return nil, err
	}

	return &lf, nil
}

func (lf *Lockfile) Save(path string) error {
	data, err := json.MarshalIndent(lf, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func (lf *Lockfile) Add(pkg Package) {
	lf.Packages[pkg.Name] = pkg
}

func (lf *Lockfile) Remove(name string) bool {
	if _, exists := lf.Packages[name]; !exists {
		return false
	}
	delete(lf.Packages, name)
	return true
}
