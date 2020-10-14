/*
Copyright 2019 The Skaffold Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gcb

import (
	"testing"

	"google.golang.org/api/cloudbuild/v1"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/build/kaniko"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/util"
	"github.com/GoogleContainerTools/skaffold/testutil"
)

func TestKanikoBuildSpec(t *testing.T) {
	tests := []struct {
		description  string
		artifact     *latest.KanikoArtifact
		expectedArgs []string
	}{
		{
			description: "simple build",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
			},
			expectedArgs: []string{},
		},
		{
			description: "with BuildArgs",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				BuildArgs: map[string]*string{
					"arg1": util.StringPtr("value1"),
					"arg2": nil,
				},
			},
			expectedArgs: []string{
				kaniko.BuildArgsFlag, "arg1=value1",
				kaniko.BuildArgsFlag, "arg2",
			},
		},
		{
			description: "with Cache",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				Cache:          &latest.KanikoCache{},
			},
			expectedArgs: []string{
				kaniko.CacheFlag,
			},
		},
		{
			description: "with Cleanup",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				Cleanup:        true,
			},
			expectedArgs: []string{
				kaniko.CleanupFlag,
			},
		},
		{
			description: "with DigestFile",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				DigestFile:     "/tmp/digest",
			},
			expectedArgs: []string{
				kaniko.DigestFileFlag, "/tmp/digest",
			},
		},
		{
			description: "with Force",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				Force:          true,
			},
			expectedArgs: []string{
				kaniko.ForceFlag,
			},
		},
		{
			description: "with ImageNameWithDigestFile",
			artifact: &latest.KanikoArtifact{
				DockerfilePath:          "Dockerfile",
				ImageNameWithDigestFile: "/tmp/imageName",
			},
			expectedArgs: []string{
				kaniko.ImageNameWithDigestFileFlag, "/tmp/imageName",
			},
		},
		{
			description: "with Insecure",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				Insecure:       true,
			},
			expectedArgs: []string{
				kaniko.InsecureFlag,
			},
		},
		{
			description: "with InsecurePull",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				InsecurePull:   true,
			},
			expectedArgs: []string{
				kaniko.InsecurePullFlag,
			},
		},
		{
			description: "with InsecureRegistry",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				InsecureRegistry: []string{
					"s1.registry.url:5000",
					"s2.registry.url:5000",
				},
			},
			expectedArgs: []string{
				kaniko.InsecureRegistryFlag, "s1.registry.url:5000",
				kaniko.InsecureRegistryFlag, "s2.registry.url:5000",
			},
		},
		{
			description: "with LogFormat",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				LogFormat:      "json",
			},
			expectedArgs: []string{
				kaniko.LogFormatFlag, "json",
			},
		},
		{
			description: "with LogTimestamp",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				LogTimestamp:   true,
			},
			expectedArgs: []string{
				kaniko.LogTimestampFlag,
			},
		},
		{
			description: "with NoPush",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				NoPush:         true,
			},
			expectedArgs: []string{
				kaniko.NoPushFlag,
			},
		},
		{
			description: "with OCILayoutPath",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				OCILayoutPath:  "/tmp/builtImage",
			},
			expectedArgs: []string{
				kaniko.OCILayoutFlag, "/tmp/builtImage",
			},
		},
		{
			description: "with RegistryCertificate",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				RegistryCertificate: map[string]*string{
					"s1.registry.url": util.StringPtr("/etc/certs/certificate1.cert"),
					"s2.registry.url": util.StringPtr("/etc/certs/certificate2.cert"),
				},
			},
			expectedArgs: []string{
				kaniko.RegistryCertificateFlag, "s1.registry.url=/etc/certs/certificate1.cert",
				kaniko.RegistryCertificateFlag, "s2.registry.url=/etc/certs/certificate2.cert",
			},
		},
		{
			description: "with RegistryMirror",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				RegistryMirror: "mirror.gcr.io",
			},
			expectedArgs: []string{
				kaniko.RegistryMirrorFlag, "mirror.gcr.io",
			},
		},
		{
			description: "with Reproducible",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				Reproducible:   true,
			},
			expectedArgs: []string{
				kaniko.ReproducibleFlag,
			},
		},
		{
			description: "with SingleSnapshot",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				SingleSnapshot: true,
			},
			expectedArgs: []string{
				kaniko.SingleSnapshotFlag,
			},
		},
		{
			description: "with SkipTLS",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				SkipTLS:        true,
			},
			expectedArgs: []string{
				kaniko.SkipTLSFlag,
				kaniko.SkipTLSVerifyRegistryFlag, "gcr.io",
			},
		},
		{
			description: "with SkipTLSVerifyPull",
			artifact: &latest.KanikoArtifact{
				DockerfilePath:    "Dockerfile",
				SkipTLSVerifyPull: true,
			},
			expectedArgs: []string{
				kaniko.SkipTLSVerifyPullFlag,
			},
		},
		{
			description: "with SkipTLSVerifyRegistry",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				SkipTLSVerifyRegistry: []string{
					"s1.registry.url:5000",
					"s2.registry.url:5000",
				},
			},
			expectedArgs: []string{
				kaniko.SkipTLSVerifyRegistryFlag, "s1.registry.url:5000",
				kaniko.SkipTLSVerifyRegistryFlag, "s2.registry.url:5000",
			},
		},
		{
			description: "with SkipUnusedStages",
			artifact: &latest.KanikoArtifact{
				DockerfilePath:   "Dockerfile",
				SkipUnusedStages: true,
			},
			expectedArgs: []string{
				kaniko.SkipUnusedStagesFlag,
			},
		},
		{
			description: "with Target",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				Target:         "builder",
			},
			expectedArgs: []string{
				kaniko.TargetFlag, "builder",
			},
		},
		{
			description: "with SnapshotMode",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				SnapshotMode:   "redo",
			},
			expectedArgs: []string{
				"--snapshotMode", "redo",
			},
		},
		{
			description: "with TarPath",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				TarPath:        "/workspace/tars",
			},
			expectedArgs: []string{
				kaniko.TarPathFlag, "/workspace/tars",
			},
		},
		{
			description: "with UseNewRun",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				UseNewRun:      true,
			},
			expectedArgs: []string{
				kaniko.UseNewRunFlag,
			},
		},
		{
			description: "with Verbosity",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				Verbosity:      "trace",
			},
			expectedArgs: []string{
				kaniko.VerbosityFlag, "trace",
			},
		},
		{
			description: "with WhitelistVarRun",
			artifact: &latest.KanikoArtifact{
				DockerfilePath:  "Dockerfile",
				WhitelistVarRun: true,
			},
			expectedArgs: []string{
				kaniko.WhitelistVarRunFlag,
			},
		},
		{
			description: "with WhitelistVarRun",
			artifact: &latest.KanikoArtifact{
				DockerfilePath:  "Dockerfile",
				WhitelistVarRun: true,
			},
			expectedArgs: []string{
				kaniko.WhitelistVarRunFlag,
			},
		},
		{
			description: "with Labels",
			artifact: &latest.KanikoArtifact{
				DockerfilePath: "Dockerfile",
				Label: map[string]*string{
					"label1": util.StringPtr("value1"),
					"label2": nil,
				},
			},
			expectedArgs: []string{
				kaniko.LabelFlag, "label1=value1",
				kaniko.LabelFlag, "label2",
			},
		},
	}

	builder := NewBuilder(&mockConfig{
		gcb: latest.GoogleCloudBuild{
			KanikoImage: "gcr.io/kaniko-project/executor",
			DiskSizeGb:  100,
			MachineType: "n1-standard-1",
			Timeout:     "10m",
		},
	})

	defaultExpectedArgs := []string{
		"--destination", "gcr.io/nginx",
		"--dockerfile", "Dockerfile",
	}

	for _, test := range tests {
		testutil.Run(t, test.description, func(t *testutil.T) {
			artifact := &latest.Artifact{
				ArtifactType: latest.ArtifactType{
					KanikoArtifact: test.artifact,
				},
			}

			desc, err := builder.buildSpec(artifact, "gcr.io/nginx", "bucket", "object")

			expected := cloudbuild.Build{
				LogsBucket: "bucket",
				Source: &cloudbuild.Source{
					StorageSource: &cloudbuild.StorageSource{
						Bucket: "bucket",
						Object: "object",
					},
				},
				Steps: []*cloudbuild.BuildStep{{
					Name: "gcr.io/kaniko-project/executor",
					Args: append(defaultExpectedArgs, test.expectedArgs...),
				}},
				Options: &cloudbuild.BuildOptions{
					DiskSizeGb:  100,
					MachineType: "n1-standard-1",
				},
				Timeout: "10m",
			}

			t.CheckNoError(err)
			t.CheckDeepEqual(expected, desc)
		})
	}
}
