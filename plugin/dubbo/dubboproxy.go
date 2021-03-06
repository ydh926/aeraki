// Copyright Aeraki Authors
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

package dubbo

import (
	"github.com/aeraki-framework/aeraki/pkg/model"
	dubbo "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/dubbo_proxy/v3"
)

func buildOutboundProxy(context *model.EnvoyFilterContext) *dubbo.DubboProxy {
	route, err := buildOutboundRouteConfig(context)
	if err != nil {
		generatorLog.Errorf("Failed to generate Dubbo EnvoyFilter: %v, %v", context.ServiceEntry, err)
		return nil
	}

	return &dubbo.DubboProxy{
		StatPrefix: model.BuildClusterName(model.TrafficDirectionOutbound, "",
			context.ServiceEntry.Spec.Hosts[0], int(context.ServiceEntry.Spec.Ports[0].Number)),
		ProtocolType:      dubbo.ProtocolType_Dubbo,
		SerializationType: dubbo.SerializationType_Hessian2,
		// we only support one to one mapping of interface and service. If there're multiple interfaces in one process,
		// these interfaces can be defined in separated services, one service for one interface.
		RouteConfig: []*dubbo.RouteConfiguration{
			route,
		},
	}
}

func buildInboundProxy(context *model.EnvoyFilterContext) *dubbo.DubboProxy {
	route, err := buildInboundRouteConfig(context)
	if err != nil {
		generatorLog.Errorf("Failed to generate Dubbo EnvoyFilter: %v, %v", context.ServiceEntry, err)
		return nil
	}

	return &dubbo.DubboProxy{
		StatPrefix: model.BuildClusterName(model.TrafficDirectionInbound, "",
			context.ServiceEntry.Spec.Hosts[0], int(context.ServiceEntry.Spec.Ports[0].Number)),
		ProtocolType:      dubbo.ProtocolType_Dubbo,
		SerializationType: dubbo.SerializationType_Hessian2,
		// we only support one to one mapping of interface and service. If there're multiple interfaces in one process,
		// these interfaces can be defined in separated services, one service for one interface.
		RouteConfig: []*dubbo.RouteConfiguration{
			route,
		},
	}
}
