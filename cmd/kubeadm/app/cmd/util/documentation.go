/*
Copyright 2017 The Kubernetes Authors.

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

package util

import (
	"k8s.io/kubernetes/pkg/util/normalizer"
)

var (
	// AlphaDisclaimer to be places at the end of description of commands in alpha release
	AlphaDisclaimer = `
		Alpha Disclaimer: this command is currently alpha but, please try it out and give us feedback!
	`

	// MacroCommandLongDescription provide a standard description for "macro" commands
	MacroCommandLongDescription = normalizer.LongDesc(`
		This command is not meant to be run on its own. See list of available subcommands.
	`)
)
