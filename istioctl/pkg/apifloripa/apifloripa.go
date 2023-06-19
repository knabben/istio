// Copyright Amim Knabben
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apifloripa

import (
	"fmt"
	"github.com/spf13/cobra"
	"istio.io/istio/istioctl/pkg/cli"
)

var labelPairs string

func Cmd(ctx cli.Context) *cobra.Command {
	apiFloripa := &cobra.Command{
		Use:   "apifloripa",
		Short: "Starts the demo setup for the 2 API Floripa",
		Long: `
Demonstration setup for 2 API Floripa, this demo contains Kiali, Jaeger and Prometheus`,
		Example: `  # Install Istio Ambient Mesh with on Kind.
  istioctl x apifloripa install
	
  # Check components for installation in parallel, exit when ready.
  istioctl x apifloripa check-components
`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				return fmt.Errorf("unknown subcommand %q", args[0])
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.HelpFunc()(cmd, args)
			return nil
		},
	}

	installApiFloripaCmd := &cobra.Command{
		Use:   "install",
		Short: "install the required resources",
		Long:  "Install all resources for DEMO",
		Example: `  # Install DEMO on existent Ambient environment
  istioctl x apifloripa install
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := ctx.CLIClient()
			if err != nil {
				return fmt.Errorf("failed to create Kubernetes client: %v", err)
			}
			return nil
		},
	}
	checkApiFloripaCmd := &cobra.Command{
		Use:   "check-components",
		Short: "Check the installed resources",
		Long:  "Check installed resources for the DEMO",
		Example: `  # Check status of the resources
  istioctl x apifloripa check
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := ctx.CLIClient()
			if err != nil {
				return fmt.Errorf("failed to create Kubernetes client: %v", err)
			}
			return nil
		},
	}
	apiFloripa.AddCommand(installApiFloripaCmd)
	apiFloripa.AddCommand(checkApiFloripaCmd)

	return apiFloripa
}
