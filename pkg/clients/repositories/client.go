package repositories

import (
	"context"
	"strings"

	"github.com/argoproj/argo-cd/v2/pkg/apiclient"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/repository"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"

	"google.golang.org/grpc"
)

const (
	errorRepositoryNotFound = "code = NotFound desc = repo"
)

// RepositoryServiceClient wraps the functions to connect to argocd repositories
type RepositoryServiceClient interface {
	// Get returns a repository or its credentials
	Get(ctx context.Context, in *repository.RepoQuery, opts ...grpc.CallOption) (*v1alpha1.Repository, error)
	// ListRepositories gets a list of all configured repositories
	ListRepositories(ctx context.Context, in *repository.RepoQuery, opts ...grpc.CallOption) (*v1alpha1.RepositoryList, error)
	// Create creates a repo or a repo credential set
	CreateRepository(ctx context.Context, in *repository.RepoCreateRequest, opts ...grpc.CallOption) (*v1alpha1.Repository, error)
	// Update updates a repo or repo credential set
	UpdateRepository(ctx context.Context, in *repository.RepoUpdateRequest, opts ...grpc.CallOption) (*v1alpha1.Repository, error)
	// Delete deletes a repository from the configuration
	DeleteRepository(ctx context.Context, in *repository.RepoQuery, opts ...grpc.CallOption) (*repository.RepoResponse, error)
}

// NewRepositoryServiceClient creates a new API client from a set of config options, or fails fatally if the new client creation fails.
func NewRepositoryServiceClient(clientOpts *apiclient.ClientOptions) repository.RepositoryServiceClient {
	_, repoIf := apiclient.NewClientOrDie(clientOpts).NewRepoClientOrDie()
	return repoIf
}

// IsErrorRepositoryNotFound helper function to test for errorRepositoryNotFound error.
func IsErrorRepositoryNotFound(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), errorRepositoryNotFound)
}
