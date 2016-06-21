package munkiserver

import (
	"github.com/micromdm/squirrel/munki/datastore"
	"github.com/micromdm/squirrel/munki/munki"
	"golang.org/x/net/context"
)

// Service describes the actions of a munki server
type Service interface {
	ListManifests(ctx context.Context) (*munki.ManifestCollection, error)
	ShowManifest(ctx context.Context, path string) (*munki.Manifest, error)
}

type service struct {
	repo datastore.Datastore
}

func (svc service) ListManifests(ctx context.Context) (*munki.ManifestCollection, error) {
	return svc.repo.AllManifests()
}

func (svc service) ShowManifest(ctx context.Context, path string) (*munki.Manifest, error) {
	return svc.repo.Manifest(path)
}

// NewService creates a new munki api service
func NewService(repo datastore.Datastore) (Service, error) {
	return &service{
		repo: repo,
	}, nil
}
