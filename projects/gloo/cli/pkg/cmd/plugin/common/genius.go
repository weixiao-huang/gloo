package common

import (
	"context"

	"github.com/spf13/cobra"
)

type PluginDescriptor struct {
	InstalledPlugin             InstalledPlugin
	RegistryPlugin              RegistryPlugin
	RemoteError, InstalledError error
}

type PluginGenius interface {
	Describe(ctx context.Context, name string) (PluginDescriptor, error)
}

type smartyPants struct {
	Remote    Registry
	Installed InstalledPluginRegistry
}

func NewPluginGenius(ctx context.Context, cmd *cobra.Command, bucket, pluginPrefix string) (PluginGenius, error) {
	remote, err := NewGcsRegistry(ctx, bucket, pluginPrefix)
	if err != nil {
		return nil, err
	}

	return &smartyPants{
		Remote:    remote,
		Installed: NewInstalledPluginRegistry(cmd),
	}, nil
}

func (s *smartyPants) Describe(ctx context.Context, name string) (PluginDescriptor, error) {
	installed, installedError := s.Installed.Get(name)
	remote, remoteError := s.Remote.Get(ctx, name)

	return PluginDescriptor{
		InstalledPlugin: installed,
		RegistryPlugin:  remote,
		InstalledError:  installedError,
		RemoteError:     remoteError,
	}, nil
}
