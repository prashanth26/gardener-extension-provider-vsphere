/*
 * Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package infra_cli

import (
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"sigs.k8s.io/yaml"

	"github.com/gardener/gardener-extension-provider-vsphere/pkg/vsphere/infrastructure"
	"github.com/gardener/gardener-extension-provider-vsphere/pkg/vsphere/infrastructure/ensurer"
)

func CreateInfrastructure(logger logr.Logger, cfg *infrastructure.NSXTConfig, specString string, fixedVersion *string) (*string, error) {
	infrastructureEnsurer, err := ensurer.NewNSXTInfrastructureEnsurer(logger, cfg, "")
	if err != nil {
		return nil, errors.Wrapf(err, "creating ensurer failed")
	}
	spec := infrastructure.NSXTInfraSpec{}
	err = yaml.Unmarshal([]byte(specString), &spec)
	if err != nil {
		return nil, errors.Wrapf(err, "unmarshalling spec failed")
	}
	state, err := infrastructureEnsurer.NewStateWithVersion(fixedVersion)
	if err != nil {
		return nil, errors.Wrapf(err, "NewStateWithVersion failed")
	}
	err = infrastructureEnsurer.EnsureInfrastructure(spec, state)
	stateString := showState(logger, state)
	if err != nil {
		return &stateString, errors.Wrapf(err, "creating infrastructure failed")
	}
	logger.Info("done")
	return &stateString, nil
}

func GetNSXTVersion(cfg *infrastructure.NSXTConfig) (*string, error) {
	return ensurer.GetNSXTVersion(cfg)
}
