/*
 Copyright Istio Authors

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

package status

import (
	"reflect"
	"sync"

	"google.golang.org/protobuf/proto"

	"istio.io/api/meta/v1alpha1"
	"istio.io/istio/pilot/pkg/features"
	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pkg/config"
	"istio.io/istio/pkg/config/schema/gvk"
)

// Manager allows multiple controllers to provide input into configuration
// status without needlessly doubling the number of writes, or overwriting
// one another.  Each status controller calls newController, passing in
// an arbitrary status modification function, and then calls EnqueueStatusUpdate
// when an individual resource is ready to be updated with the relevant data.
type Manager struct {
	// TODO: is Resource the right abstraction?
	store        model.ConfigStore
	workers      WorkerQueue
	lastStatuses sync.Map
}

func NewManager(store model.ConfigStore) *Manager {
	manager := &Manager{
		store:        store,
		lastStatuses: sync.Map{},
	}
	writeFunc := func(m *config.Config, istatus any) {
		status := istatus.(GenerationProvider)
		m.Status = status.Unwrap()
		lastStatus, ok := manager.lastStatuses.Load(m.Key())
		if !ok {
			cfg := store.Get(m.GroupVersionKind, m.Name, m.Namespace)
			lastStatus = cfg.Status
			manager.lastStatuses.Store(m.Key(), lastStatus)
		}
		switch t := m.Status.(type) {
		case proto.Message:
			if lastStatus != nil && proto.Equal(lastStatus.(proto.Message), t) {
				return
			}
		default:
			if reflect.DeepEqual(lastStatus, t) {
				return
			}
		}

		scope.Debugf("writing status for resource %s/%s", m.Namespace, m.Name)
		_, err := store.UpdateStatus(*m)
		if err != nil {
			// TODO: need better error handling
			scope.Errorf("Encountered unexpected error updating status for %v, will try again later: %s", m, err)
			return
		}
		// store the last status for this resource if write was successful
		manager.lastStatuses.Store(m.Key(), m.Status)
	}
	retrieveFunc := func(resource Resource) *config.Config {
		scope.Debugf("retrieving config for status update: %s/%s", resource.Namespace, resource.Name)
		k, ok := gvk.FromGVR(resource.GroupVersionResource)
		if !ok {
			scope.Warnf("GVR %v could not be identified", resource.GroupVersionResource)
			return nil
		}

		current := store.Get(k, resource.Name, resource.Namespace)
		return current
	}
	manager.workers = NewWorkerPool(writeFunc, retrieveFunc, uint(features.StatusMaxWorkers))
	return manager
}

func (m *Manager) Start(stop <-chan struct{}) {
	scope.Info("Starting status manager")

	ctx := NewIstioContext(stop)
	m.workers.Run(ctx)
}

// CreateGenericController provides an interface for a status update function to be
// called in series with other controllers, minimizing the number of actual
// api server writes sent from various status controllers.  The UpdateFunc
// must take the target resrouce status and arbitrary context information as
// parameters, and return the updated status value.  Multiple controllers
// will be called in series, so the input status may not have been written
// to the API server yet, and the output status may be modified by other
// controllers before it is written to the server.
func (m *Manager) CreateGenericController(fn UpdateFunc) *Controller {
	result := &Controller{
		fn:      fn,
		workers: m.workers,
	}
	return result
}

func (m *Manager) CreateIstioStatusController(fn func(status *v1alpha1.IstioStatus, context any) *v1alpha1.IstioStatus) *Controller {
	wrapper := func(status any, context any) GenerationProvider {
		var input *v1alpha1.IstioStatus
		if status != nil {
			converted := status.(*IstioGenerationProvider)
			input = converted.IstioStatus
		}
		result := fn(input, context)
		return &IstioGenerationProvider{result}
	}
	result := &Controller{
		fn:      wrapper,
		workers: m.workers,
	}
	return result
}

type UpdateFunc func(status any, context any) GenerationProvider

type Controller struct {
	fn      UpdateFunc
	workers WorkerQueue
}

// EnqueueStatusUpdateResource informs the manager that this controller would like to
// update the status of target, using the information in context.  Once the status
// workers are ready to perform this update, the controller's UpdateFunc
// will be called with target and context as input.
func (c *Controller) EnqueueStatusUpdateResource(context any, target Resource) {
	// TODO: buffer this with channel
	c.workers.Push(target, c, context)
}

func (c *Controller) Delete(r Resource) {
	c.workers.Delete(r)
}
