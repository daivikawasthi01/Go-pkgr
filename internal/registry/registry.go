package registry

import (
	"fmt"
	"strings"
)

type PackageInfo struct {
	Name    string
	Version string
	Source  string
}

var fakeRegistry = map[string]PackageInfo{
	"cobra": {Name: "cobra", Version: "1.10.2", Source: "github.com/spf13/cobra"},
	"gin":   {Name: "gin", Version: "1.9.1", Source: "github.com/gin-gonic/gin"},
	"viper": {Name: "viper", Version: "1.18.0", Source: "github.com/spf13/viper"},
	"zap":   {Name: "zap", Version: "1.27.0", Source: "go.uber.org/zap"},
	"chi":   {Name: "chi", Version: "5.0.12", Source: "github.com/go-chi/chi"},
}

func Resolve(name string) (PackageInfo, error) {
	name = strings.ToLower(name)
	pkg, exists := fakeRegistry[name]
	if !exists {
		return PackageInfo{}, fmt.Errorf("package %q not found in registry", name)
	}
	return pkg, nil
}

func List() []PackageInfo {
	all := make([]PackageInfo, 0, len(fakeRegistry))
	for _, pkg := range fakeRegistry {
		all = append(all, pkg)
	}
	return all
}
