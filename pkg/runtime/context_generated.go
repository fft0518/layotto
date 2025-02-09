// Code generated by github.com/seeflood/protoc-gen-p6 .

// Copyright 2021 Layotto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package runtime

import (
	grpc "mosn.io/layotto/pkg/grpc"
)

func newApplicationContext(m *MosnRuntime) *grpc.ApplicationContext {
	return &grpc.ApplicationContext{
		AppId:                 m.runtimeConfig.AppManagement.AppId,
		Hellos:                m.hellos,
		ConfigStores:          m.configStores,
		Rpcs:                  m.rpcs,
		PubSubs:               m.pubSubs,
		StateStores:           m.states,
		Files:                 m.files,
		Oss:                   m.oss,
		LockStores:            m.locks,
		Sequencers:            m.sequencers,
		SendToOutputBindingFn: m.sendToOutputBinding,
		SecretStores:          m.secretStores,
		DynamicComponents:     m.dynamicComponents,
		CustomComponent:       m.customComponent,
		EmailService:          m.emailService,

		PhoneCallService: m.phoneCallService,
	}
}
