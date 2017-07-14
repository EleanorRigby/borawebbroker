/*
Copyright 2016 The Kubernetes Authors.
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

package controller

import (
	"errors"
	"fmt"

	"github.com/eleanorrigby/borawebbroker/client"
	"github.com/golang/glog"
	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/broker/controller"
	"github.com/kubernetes-incubator/service-catalog/pkg/brokerapi"
)

type errNoSuchInstance struct {
	instanceID string
}

func (e errNoSuchInstance) Error() string {
	return fmt.Sprintf("no such instance with ID %s", e.instanceID)
}

type borawebServiceInstance struct {
}

type borawebController struct {
}

// CreateController creates an instance of a service broker controller.
func CreateController() controller.Controller {
	return &borawebController{}
}

func (c *borawebController) Catalog() (*brokerapi.Catalog, error) {
	return &brokerapi.Catalog{
		Services: []*brokerapi.Service{
			{
				Name:        "web-publishing-service",
				ID:          client.WebPub,
				Description: "Web Pub as a service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "86064792-7ea2-467b-af93-ac9694d96d52",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
		},
	}, nil
}

func (c *borawebController) CreateServiceInstance(
	id string,
	req *brokerapi.CreateServiceInstanceRequest,
) (*brokerapi.CreateServiceInstanceResponse, error) {

	var namespaceParam string

	if req.ContextProfile.Namespace == "" {
		namespaceParam = client.ReleaseName(id)
	} else {
		namespaceParam = client.ReleaseName(id) /****TODO TA ***/ //req.ContextProfile.Namespace
		//Commented because it is not knowwn right now that how to get name of a namespace while creating a binding.
		//we will need to create host address which will require clientset.Core().Pods().Get("fg").GetNamespace()
		//(Just figured this out) make changes later
	}
	//Based on Service ID we can invoke specific chart installation.
	if err := client.Install(id, namespaceParam, req.Parameters); err != nil {
		return nil, err
	}

	glog.Infof("Created %s Service Instance:\n\n", id)
	//glog.Info("Printing request %v", *req)
	return &brokerapi.CreateServiceInstanceResponse{}, nil
}

func (c *borawebController) GetServiceInstance(id string) (string, error) {
	return "", errors.New("Unimplemented")
}

func (c *borawebController) RemoveServiceInstance(id string) (*brokerapi.DeleteServiceInstanceResponse, error) {

	if err := client.Delete(id); err != nil {
		return nil, err
	}

	return &brokerapi.DeleteServiceInstanceResponse{}, nil
}

func (c *borawebController) Bind(
	instanceID,
	bindingID string,
	req *brokerapi.BindingRequest,
) (*brokerapi.CreateServiceBindingResponse, error) {

	var err error
	var creds client.DBCreds

	if creds, err = client.GetBinding(instanceID); err != nil {

		if err != nil {
			return nil, err
		}
		return &brokerapi.CreateServiceBindingResponse{
			Credentials: brokerapi.Credential{
				"username": creds.Username,
				"password": creds.Password,
				"port":     creds.Port,
				"host":     creds.Host,
				"uri":      creds.Uri,
				"database": creds.Database,
			},
		}, nil

	}

	return nil, err

}

func (c *borawebController) UnBind(instanceID string, bindingID string) error {
	// Since we don't persist the binding, there's nothing to do here.
	return nil
}
