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
				Name:        "mysql",
				ID:          client.MySQL_id,
				Description: "my sql-service",
				Tags:        []string{"https://www.google.com/url?sa=i&rct=j&q=&esrc=s&source=images&cd=&cad=rja&uact=8&ved=0ahUKEwjGhfKHvJbVAhVUVWMKHXWyDfcQjRwIBw&url=https%3A%2F%2Fwww.mysql.com%2Fabout%2Flegal%2Flogos.html&psig=AFQjCNEMKsHpZjGFfG41ZSdtwaIrYS3Vxw&ust=1500592292658463"},
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          client.MySQL_plan_id,
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "wordpress",
				ID:          client.Wordpress_id,
				Description: "wordpress service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          client.Wordpress_id,
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "drupal",
				ID:          client.Drupal_id,
				Description: "Drupal service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          client.Drupal_plan_id,
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "mariadb",
				ID:          client.MariaDB_id,
				Description: "Maria DB service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          client.MariaDB_plan_id,
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

	glog.Info("DEBUG  :  Printing createrequest : ", *req)

	glog.Info("DEBUG  :  Printing parameters received : ", req.Parameters)

	//Based on Service ID we can invoke specific chart installation.
	if err := client.Install(req.ServiceID, (req.Parameters["instance"]).(string), id, req.Parameters["namespace"].(string), req.Parameters); err != nil {
		return nil, err
	}

	glog.Infof("Created %s Service Instance:\n\n", id)

	return &brokerapi.CreateServiceInstanceResponse{}, nil
}

func (c *borawebController) GetServiceInstance(id string) (string, error) {
	return client.NameDetails[id].Name, nil
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

	glog.Infof("DEBUG  :  Printing parameters of bindin request  %v ", req.Parameters)

	if creds, err = client.GetBinding(req.ServiceID, instanceID); err != nil {

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
