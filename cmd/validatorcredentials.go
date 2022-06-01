// Copyright © 2022 Weald Technology Trading
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

package cmd

import (
	"github.com/spf13/cobra"
)

// validatorCredentialsCmd represents the validator credentials command
var validatorCredentialsCmd = &cobra.Command{
	Use:   "credentials",
	Short: "Manage Ethereum consensu validator credentials",
	Long:  `Manage Ethereum consensu validator credentials.`,
}

func init() {
	validatorCmd.AddCommand(validatorCredentialsCmd)
}

func validatorCredentialsFlags(cmd *cobra.Command) {
}