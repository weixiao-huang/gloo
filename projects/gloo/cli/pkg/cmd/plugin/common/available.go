package common

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/spf13/cobra"
)

type InstalledPlugin struct {
	Name     string
	FullPath string
	Warnings []error
}

type InstalledPluginRegistry interface {
	ListAll() (plugins []InstalledPlugin, warnings []error, err error)
	Get(name string) (InstalledPlugin, error)
	Uninstall(name string) error
}

type installedPluginRegistry struct {
	verifier    PathVerifier
	pluginPaths []string
}

func NewInstalledPluginRegistry(cmd *cobra.Command) InstalledPluginRegistry {
	r := &installedPluginRegistry{
		verifier: &CommandOverrideVerifier{
			root:        cmd.Root(),
			seenPlugins: make(map[string]string),
		},
		pluginPaths: filepath.SplitList(os.Getenv("PATH")),
	}

	return r
}

func (r *installedPluginRegistry) ListAll() (plugins []InstalledPlugin, warnings []error, err error) {
	for _, dir := range uniquePathsList(r.pluginPaths) {
		if len(strings.TrimSpace(dir)) == 0 {
			continue
		}

		files, err := ioutil.ReadDir(dir)
		if err != nil {
			warnings = append(warnings, eris.Errorf("unable to read directory %q in your PATH: %v", dir, err))
			continue
		}

		for _, f := range files {
			if f.IsDir() {
				continue
			}
			if !hasValidPrefix(f.Name(), constants.ValidExtensionPrefixes) {
				continue
			}

			// TODO joekelley cleanup
			name := strings.TrimPrefix(f.Name(), "glooctl-")

			plugin := InstalledPlugin{
				Name:     name,
				FullPath: filepath.Join(dir, f.Name()),
			}

			if errs := r.verifier.Verify(filepath.Join(dir, f.Name()), f.Name()); len(errs) != 0 {
				for _, err := range errs {
					plugin.Warnings = append(plugin.Warnings, err)
				}
			}

			plugins = append(plugins, plugin)
		}
	}

	return plugins, warnings, nil
}

func (r *installedPluginRegistry) Get(name string) (InstalledPlugin, error) {
	plugins, _, err := r.ListAll()
	if err != nil {
		return InstalledPlugin{}, err
	}

	for _, plugin := range plugins {
		if plugin.Name == name {
			return plugin, nil
		}
	}

	return InstalledPlugin{}, eris.Errorf("%s was not found in your PATH", name)
}

func (r *installedPluginRegistry) Uninstall(name string) error {
	plugin, err := r.Get(name)
	if err != nil {
		return err
	}

	return os.Remove(plugin.FullPath)
}

// pathVerifier receives a path and determines if it is valid or not
type PathVerifier interface {
	// Verify determines if a given path is valid
	Verify(path, binName string) []error
}

type CommandOverrideVerifier struct {
	root        *cobra.Command
	seenPlugins map[string]string
}

// Verify implements PathVerifier and determines if a given path
// is valid depending on whether or not it overwrites an existing
// glooctl command path, or a previously seen plugin.
func (v *CommandOverrideVerifier) Verify(path, binName string) []error {
	if v.root == nil {
		return []error{eris.Errorf("unable to verify path with nil root")}
	}

	errors := []error{}

	cmdPath := strings.Split(binName, "-")
	if len(cmdPath) > 1 {
		if cmdPath[0] != "glooctl" {
			errors = append(errors, eris.Errorf("warning: the prefix should always be 'glooctl' for a plugin binary, found: %s", cmdPath[0]))
			return errors
		}
		cmdPath = cmdPath[1:]
	}

	if isExec, err := isExecutable(path); err == nil && !isExec {
		errors = append(errors, eris.Errorf("warning: %s identified as a glooctl plugin, but it is not executable", path))
	} else if err != nil {
		errors = append(errors, eris.Errorf("unable to identify %s as an executable file: %v", path, err))
	}

	if existingPath, ok := v.seenPlugins[binName]; ok {
		errors = append(errors, eris.Errorf("warning: %s is overshadowed by a similarly named plugin: %s", path, existingPath))
	} else {
		v.seenPlugins[binName] = path
	}

	if cmd, _, err := v.root.Find(cmdPath); err == nil {
		errors = append(errors, eris.Errorf("warning: %s overwrites existing command: %q", binName, cmd.CommandPath()))
	}

	return errors
}

func isExecutable(fullPath string) (bool, error) {
	info, err := os.Stat(fullPath)
	if err != nil {
		return false, err
	}

	if runtime.GOOS == "windows" {
		fileExt := strings.ToLower(filepath.Ext(fullPath))

		switch fileExt {
		case ".bat", ".cmd", ".com", ".exe", ".ps1":
			return true, nil
		}
		return false, nil
	}

	if m := info.Mode(); !m.IsDir() && m&0111 != 0 {
		return true, nil
	}

	return false, nil
}

// uniquePathsList deduplicates a given slice of strings without
// sorting or otherwise altering its order in any way.
func uniquePathsList(paths []string) []string {
	seen := map[string]bool{}
	newPaths := []string{}
	for _, p := range paths {
		if seen[p] {
			continue
		}
		seen[p] = true
		newPaths = append(newPaths, p)
	}
	return newPaths
}

func hasValidPrefix(filepath string, validPrefixes []string) bool {
	for _, prefix := range validPrefixes {
		if !strings.HasPrefix(filepath, prefix+"-") {
			continue
		}
		return true
	}
	return false
}
