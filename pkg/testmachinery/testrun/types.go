//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testrun

import (
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/gardener/test-infra/pkg/testmachinery/testflow/node"

	tmv1beta1 "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1"
	"github.com/gardener/test-infra/pkg/testmachinery/testflow"
)

// Testrun is the internal representation of a testrun crd
type Testrun struct {
	Info            *tmv1beta1.Testrun
	Testflow        *testflow.Testflow
	OnExitTestflow  *testflow.Testflow
	HelperResources []client.Object
	ProjectedTokens map[string]*node.ProjectedTokenMount
}
