package testutils

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"testing"

	"github.com/filecoin-project/bacalhau/pkg/docker"
	"github.com/filecoin-project/bacalhau/pkg/model"
	"github.com/filecoin-project/bacalhau/pkg/publicapi"
	"github.com/filecoin-project/bacalhau/pkg/system"
	"github.com/stretchr/testify/require"
)

func GetJobFromTestOutput(ctx context.Context, t *testing.T, c *publicapi.APIClient, out string) *model.Job {
	jobID := system.FindJobIDInTestOutput(out)
	uuidRegex := regexp.MustCompile(`[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`)
	require.Regexp(t, uuidRegex, jobID, "Job ID should be a UUID")

	j, _, err := c.Get(ctx, jobID)
	require.NoError(t, err)
	require.NotNil(t, j, "Failed to get job with ID: %s", out)
	return j
}

func FirstFatalError(t *testing.T, output string) (model.TestFatalErrorHandlerContents, error) {
	linesInOutput := system.SplitLines(output)
	fakeFatalError := &model.TestFatalErrorHandlerContents{}
	for _, line := range linesInOutput {
		err := model.JSONUnmarshalWithMax([]byte(line), fakeFatalError)
		if err != nil {
			return model.TestFatalErrorHandlerContents{}, err
		} else {
			return *fakeFatalError, nil
		}
	}
	return model.TestFatalErrorHandlerContents{}, fmt.Errorf("no fatal error found in output")
}

// If the test is running in an environment that cannot support cross-platform
// Docker images, the test is skipped.
func MustHaveDocker(t *testing.T) {
	MaybeNeedDocker(t, true)
}

// If the test is running in an environment that cannot support cross-platform
// Docker images, and the passed boolean flag is true, the test is skipped.
func MaybeNeedDocker(t *testing.T, needDocker bool) {
	_, isCI := os.LookupEnv("CI")
	if needDocker && isCI && (runtime.GOOS == "windows" || runtime.GOOS == "darwin") {
		t.Skip("Cannot run this test in a", runtime.GOOS, "runtime on a CI environment because it requires Docker")
	}

	if needDocker {
		client, err := docker.NewDockerClient()
		require.NoError(t, err)

		installed := docker.IsInstalled(context.Background(), client)
		if !installed {
			t.Fatalf("Docker is not running")
		}
	}
}

func SkipIfArm(t *testing.T, issueURL string) {
	if runtime.GOARCH == "arm64" {
		t.Skip("Test does not pass natively on arm64", issueURL)
	}
}

func MakeGenericJob() *model.Job {
	return MakeJob(model.EngineDocker, model.VerifierNoop, model.PublisherNoop, []string{
		"echo",
		"$(date +%s)",
	})
}

func MakeNoopJob() *model.Job {
	return MakeJob(model.EngineNoop, model.VerifierNoop, model.PublisherNoop, []string{
		"echo",
		"$(date +%s)",
	})
}

func MakeJob(
	engineType model.Engine,
	verifierType model.Verifier,
	publisherType model.Publisher,
	entrypointArray []string) *model.Job {
	j := model.NewJob()

	j.Spec = model.Spec{
		Engine:    engineType,
		Verifier:  verifierType,
		Publisher: publisherType,
		Docker: model.JobSpecDocker{
			Image:      "ubuntu:latest",
			Entrypoint: entrypointArray,
		},
		// Inputs:  inputStorageList,
		// Outputs: testCase.Outputs,
	}

	j.Deal = model.Deal{
		Concurrency: 1,
	}

	return j
}
