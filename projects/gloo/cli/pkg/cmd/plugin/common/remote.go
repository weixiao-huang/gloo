package common

import (
	"context"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// RegistryPlugin is a plugin made available via plugin registry.
type RegistryPlugin struct {
	// Name is the name of the plugin, e.g. "glooctl-fed"
	Name string
	// Display is the user-friendly name of the plugin, e.g. "fed"
	DisplayName string
	// AvailableVersions describes the download URL for available versions of the RegistryPlugin
	AvailableVersions VersionedPlugins
}

// Registry represents a GcsRegistry of glooctl plugins.
type Registry interface {
	// Search returns all RegistryPlugins in the Registry with names containing the provided query string.
	Search(ctx context.Context, query string) ([]RegistryPlugin, error)
	// Get returns the RegistryPlugin with a name that exactly matches the provided plugin name.
	Get(ctx context.Context, name string) (RegistryPlugin, error)
}

// GcsRegistry is a Registry implementation backed by a Google Cloud Storage bucket.
type GcsRegistry struct {
	Client       *storage.Client
	Bucket       string
	PluginPrefix string
}

// NewGcsRegistry returns a Registry backed by a GCS bucket.
func NewGcsRegistry(ctx context.Context, bucket, pluginPrefix string) (*GcsRegistry, error) {
	client, err := storage.NewClient(ctx, option.WithScopes(storage.ScopeReadOnly), option.WithoutAuthentication())
	if err != nil {
		return nil, err
	}

	return &GcsRegistry{
		Client:       client,
		Bucket:       bucket,
		PluginPrefix: pluginPrefix,
	}, nil
}

func (r *GcsRegistry) Search(ctx context.Context, query string) ([]RegistryPlugin, error) {
	return r.search(ctx, query)
}

func (r *GcsRegistry) Get(ctx context.Context, name string) (RegistryPlugin, error) {
	list, err := r.search(ctx, name)
	if err != nil {
		return RegistryPlugin{}, err
	}

	for _, plugin := range list {
		if plugin.DisplayName == name {
			return plugin, nil
		}
	}
	return RegistryPlugin{}, eris.Errorf("Plugin %s not found.", name)
}

func (r *GcsRegistry) search(ctx context.Context, query string) ([]RegistryPlugin, error) {
	objects, err := r.listAllObjects(ctx)
	if err != nil {
		return nil, err
	}

	foundPlugins := make(map[string]VersionedPlugins)

	var plugins []RegistryPlugin
	for _, object := range objects {
		// Valid plugin objects are structured as "glooctl-{name}/{semver version}/glooctl-{name}-os-arch{optional .exe}"
		parts := strings.Split(strings.TrimSuffix(object.Name, "/"), "/")
		if len(parts) != 3 {
			continue
		}

		pluginName, version, binaryName := parts[0], parts[1], parts[2]

		if !strings.Contains(pluginName, query) {
			continue
		}

		pluginVersion := PluginVersion(version)

		if _, ok := foundPlugins[pluginName]; !ok {
			foundPlugins[pluginName] = make(VersionedPlugins)
		}
		if _, ok := foundPlugins[pluginName][pluginVersion]; !ok {
			foundPlugins[pluginName][pluginVersion] = make(map[string]string)
		}
		foundPlugins[pluginName][pluginVersion][binaryName] = object.MediaLink
	}

	for pluginName, versionedPlugins := range foundPlugins {
		plugins = append(plugins, RegistryPlugin{
			Name:              pluginName,
			DisplayName:       strings.TrimPrefix(pluginName, r.PluginPrefix),
			AvailableVersions: versionedPlugins,
		})
	}

	return plugins, nil
}

func (r *GcsRegistry) listAllObjects(ctx context.Context) ([]*storage.ObjectAttrs, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	var objects []*storage.ObjectAttrs
	it := r.Client.Bucket(constants.GlooctlPluginBucket).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, eris.Wrap(err, "Error listing available plugins")
		}

		objects = append(objects, attrs)
	}
	return objects, nil
}
