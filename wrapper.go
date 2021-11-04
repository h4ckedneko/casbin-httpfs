package casbinfs

import (
	"io/ioutil"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
)

const (
	defaultPathModel  = "model.conf"
	defaultPathPolicy = "policy.csv"
)

func normalizePaths(paths []string) []string {
	if len(paths) == 0 {
		return []string{
			defaultPathModel,
			defaultPathPolicy,
		}
	}
	if len(paths) == 1 {
		return []string{
			paths[0] + "/" + defaultPathModel,
			paths[0] + "/" + defaultPathPolicy,
		}
	}
	return []string{paths[0], paths[1]}
}

// NewModel returns a newly initialized Model instance.
func NewModel(fs http.FileSystem, path string) (model.Model, error) {
	f, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return model.NewModelFromString(string(b))
}

// NewEnforcer returns a newly initialized Enforcer instance.
func NewEnforcer(fs http.FileSystem, paths ...string) (*casbin.Enforcer, error) {
	paths = normalizePaths(paths)
	model, err := NewModel(fs, paths[0])
	if err != nil {
		return nil, err
	}
	return casbin.NewEnforcer(model, NewAdapter(fs, paths[1]))
}

// NewCachedEnforcer returns a newly initialized CachedEnforcer instance.
func NewCachedEnforcer(fs http.FileSystem, paths ...string) (*casbin.CachedEnforcer, error) {
	paths = normalizePaths(paths)
	model, err := NewModel(fs, paths[0])
	if err != nil {
		return nil, err
	}
	return casbin.NewCachedEnforcer(model, NewAdapter(fs, paths[1]))
}

// NewSyncedEnforcer returns a newly initialized SyncedEnforcer instance.
func NewSyncedEnforcer(fs http.FileSystem, paths ...string) (*casbin.SyncedEnforcer, error) {
	paths = normalizePaths(paths)
	model, err := NewModel(fs, paths[0])
	if err != nil {
		return nil, err
	}
	return casbin.NewSyncedEnforcer(model, NewAdapter(fs, paths[1]))
}
