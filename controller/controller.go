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
			/*{
				Name:        "Cockroach DB",
				ID:          "A",
				Description: "my cockroachdb-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "105",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "Gerrit",
				ID:          "B",
				Description: "my gerrit-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "103",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "Gogs",
				ID:          "C",
				Description: "my gogs-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "102",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "Jenkins",
				ID:          "D",
				Description: "my jenkins-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "101",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "Spinnaker",
				ID:          "E",
				Description: "my spinnaker-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "100",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "Redis",
				ID:          "F",
				Description: "my redis-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "1",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "Postgres",
				ID:          "G",
				Description: "my postgres-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "2",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "Cassandra",
				ID:          "H",
				Description: "my cassandra-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "3",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "Spark",
				ID:          "I",
				Description: "my spark-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "4",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "MemCached",
				ID:          "J",
				Description: "my memcached-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "5",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "Linkerd",
				ID:          "K",
				Description: "my linkerd-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "6",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "Kafka",
				ID:          "L",
				Description: "my kafka-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "7",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "Joomla",
				ID:          "M",
				Description: "my joomla-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "8",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "DataDog",
				ID:          "N",
				Description: "my data dog-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "9",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "artifactory",
				ID:          "O",
				Description: "my artifactory-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "10",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "gitlab",
				ID:          "P",
				Description: "my gitlab-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "11",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
			{
				Name:        "consul",
				ID:          "Q",
				Description: "my consul-service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "12",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},*/
			{
				Name:        "mysql",
				ID:          client.MySQL_id,
				Description: "my sql-service",
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

	if creds, err = client.GetBinding(req.ServiceID, instanceID); err == nil {

		glog.Infof("DEBUG  :  Creds ", creds)

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

	if err != nil {
		return nil, err
	}

	return nil, err

}

func (c *borawebController) UnBind(instanceID string, bindingID string) error {
	// Since we don't persist the binding, there's nothing to do here.
	return nil
}
