package main

import (
	"context"
	"dagger/renovate/internal/dagger"
	"fmt"
)

type Renovate struct{}

// Returns lines that match a pattern in the files of the provided Directory
func (m *Renovate) RenovateScan(
	ctx context.Context,
	repository string,
	// +optional
	// +default="main"
	baseBranche string,
	renovateToken *dagger.Secret,
	// +optional
	// +default="info"
	logLevel string,
) (string, error) {
	return dag.Container().From("renovate/renovate:38").
		WithEnvVariable("LOG_LEVEL", logLevel).
		WithEnvVariable("RENOVATE_REPOSITORIES", repository).
		WithEnvVariable("RENOVATE_BASE_BRANCHES", fmt.Sprintf("[\"%s\"]", baseBranche)).
		WithSecretVariable("RENOVATE_TOKEN", renovateToken).
		WithExec([]string{"renovate", "--platform", "github", "--onboarding", "false"}).Stdout(ctx)
}
