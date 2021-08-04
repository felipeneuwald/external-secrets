/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
limitations under the License.
*/
package gitlab

import (
	"os"

	// nolint
	. "github.com/onsi/ginkgo"
	// nolint
	. "github.com/onsi/ginkgo/extensions/table"

	"github.com/external-secrets/external-secrets/e2e/framework"
	"github.com/external-secrets/external-secrets/e2e/suite/common"
)

var _ = Describe("[gitlab] ", func() {
	f := framework.New("esogitlab")
	credentials := os.Getenv("GITLAB_TOKEN")
	projectID := os.Getenv("GITLAB_PROJECT_ID")
	prov := newGitlabProvider(f, credentials, projectID)

	DescribeTable("sync secrets", framework.TableFunc(f, prov),
		Entry(common.SimpleDataSync(f)),
		// Entry(common.JSONDataWithProperty(f)),
		// Entry(common.JSONDataFromSync(f)),
		// Entry(common.NestedJSONWithGJSON(f)),
		// Entry(common.JSONDataWithTemplate(f)),
	)
})
