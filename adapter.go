package casbinfs

import (
	"bufio"
	"errors"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

var (
	loader  = persist.LoadPolicyLine
	trimmer = strings.TrimSpace
)

// LoadPolicyFromFS loads all policy rules from http.FileSystem.
// It registers all the loaded policy rules into a specified model.
// This function is used by the Adapter but feel free to use it as is.
func LoadPolicyFromFS(fs http.FileSystem, path string, model model.Model) error {
	f, err := fs.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		loader(trimmer(line), model)
	}
	return scanner.Err()
}

// An Adapter implements the Casbin adapter for http.FileSystem.
type Adapter struct {
	fs   http.FileSystem
	path string
}

// New returns a newly initialized Adapter instance.
func New(fs http.FileSystem, path string) *Adapter {
	return &Adapter{fs, path}
}

// LoadPolicy loads all policy rules from the storage.
func (a *Adapter) LoadPolicy(model model.Model) error {
	return LoadPolicyFromFS(a.fs, a.path, model)
}

// SavePolicy saves all policy rules into the storage.
func (a *Adapter) SavePolicy(model model.Model) error {
	return errors.New("not implemented")
}

// AddPolicy adds a policy rule into the storage.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	return errors.New("not implemented")
}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return errors.New("not implemented")
}

// RemoveFilteredPolicy removes a policy rule from storage that matches the filter.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fid int, fvals ...string) error {
	return errors.New("not implemented")
}
