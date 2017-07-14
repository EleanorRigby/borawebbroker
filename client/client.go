package client

import (
	"errors"

	yaml "gopkg.in/yaml.v2"

	"github.com/dchest/uniuri"
	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/helm/pkg/helm"
)

// DBCreds provide credentials for datbase
type DBCreds struct {
	Host     string
	Uri      string
	Database string
	Username string
	Password string
	Port     string
}

const (
	tillerHost = "tiller-deploy.kube-system.svc.cluster.local:44134"

	chartPath      = "/chart"
	wordpressChart = chartPath + "/wordpress"
	drupalChart    = chartPath + "/drupal"
	mariaDBChart   = chartPath + "/mariadb"
	mySQLChart     = chartPath + "/mysql"

	wordpress_id      = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2468"
	wordpress_plan_id = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2469"
	drupal_id         = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2460"
	drupal_plan_id    = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2461"
	mariaDB_id        = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2462"
	mariaDB_plan_id   = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2463"
	mySQL_id          = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2464"
	mySQL_plan_id     = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2465"
)

var mapRelNames = map[string]string{
	//release names
	"4f6e6cf6-ffdd-425f-a2c7-3c9258ad2468": "whackywordpress",
	"4f6e6cf6-ffdd-425f-a2c7-3c9258ad2460": "DrogoDrupal",
	"4f6e6cf6-ffdd-425f-a2c7-3c9258ad2462": "MyriadMaria",
	"4f6e6cf6-ffdd-425f-a2c7-3c9258ad2464": "SuperSQL",
}

var mapAppInstallFunctions = map[string]interface{}{
	"wordpress": installWordpress,
	"drupal":    installDrupal,
	"mariadb":   installMariaDB,
	"mysql":     installMySQL,
}

var mapAppBindFunctions = map[string]interface{}{
	"wordpress": bindWordpress,
	"drupal":    bindDrupal,
	"mariadb":   bindMariaDB,
	"mysql":     bindMySQL,
}

func bindWordpress(id string) (DBCreds, error) {

	var cred DBCreds

	return cred, nil

}

func bindDrupal(id string) (DBCreds, error) {

	var cred DBCreds

	return cred, nil

}

func bindMariaDB(id string) (DBCreds, error) {

	var creds DBCreds

	creds.Host = ReleaseName(id) + "-mariadb." + ReleaseName(id) + ".svc.cluster.local"
	creds.Port = "3306"
	creds.Database = "dbname"
	creds.Username = "root"

	config, err := rest.InClusterConfig()
	if err != nil {
		return creds, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return creds, err
	}
	secret, err := clientset.Core().Secrets(ReleaseName(id)).Get(ReleaseName(id) + "-mariadb")
	if err != nil {
		return creds, err
	}

	creds.Password = string(secret.Data["mariadb-root-password"])
	creds.Uri = "mysql://" + creds.Username + ":" + creds.Password + "@" + creds.Host + ":" + creds.Port + "/" + creds.Database

	glog.Info("Debug data for binding credentials for %s are  %v", ReleaseName(id), creds)

	return creds, nil

}

func bindMySQL(id string) (DBCreds, error) {

	var creds DBCreds

	creds.Host = ReleaseName(id) + "-mysql." + ReleaseName(id) + ".svc.cluster.local"
	creds.Port = "3306"
	creds.Database = "dbname"
	creds.Username = "root"

	config, err := rest.InClusterConfig()
	if err != nil {
		return creds, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return creds, err
	}
	secret, err := clientset.Core().Secrets(ReleaseName(id)).Get(ReleaseName(id) + "-mariadb")
	if err != nil {
		return creds, err
	}

	creds.Password = string(secret.Data["mysql-root-password"])
	creds.Uri = "mysql://" + creds.Username + ":" + creds.Password + "@" + creds.Host + ":" + creds.Port + "/" + creds.Database

	glog.Info("Debug data for binding credentials for %s are  %v", ReleaseName(id), creds)

	return creds, nil

}

// GetBinding returns credential for the passed instance ID
func GetBinding(id string) (DBCreds, error) {

	var err error
	var creds DBCreds

	switch id {
	case wordpress_id:
		creds, err = mapAppBindFunctions["wordpress"].(func(string) (DBCreds, error))(id)
	case mariaDB_id:
		creds, err = mapAppBindFunctions["mariadb"].(func(string) (DBCreds, error))(id)
	case drupal_id:
		creds, err = mapAppBindFunctions["drupal"].(func(string) (DBCreds, error))(id)
	case mySQL_id:
		creds, err = mapAppBindFunctions["mysql"].(func(string) (DBCreds, error))(id)
	default:
		glog.Info("Something Very Very Wrong")
		err = errors.New("Wrong Service ID Passed in Bind function")
	}

	if err != nil {
		glog.Infof("Failed to create binding for %s : %v \n\n", ReleaseName(id), err)
		return creds, err
	}

	return creds, nil

}

func installWordpress(releaseName, namespace string, parameter map[string]interface{}) error {

	vals, err := yaml.Marshal(parameter)
	if err != nil {
		return err
	}
	helmClient := helm.NewClient(helm.Host(tillerHost))
	_, err = helmClient.InstallRelease(wordpressChart, namespace, helm.ReleaseName(releaseName), helm.ValueOverrides(vals))
	if err != nil {
		glog.Infof("Failed to create wordpress : %v \n\n", err)
		return err
	}

	return nil
}

func installDrupal(releaseName, namespace string, parameter map[string]interface{}) error {

	vals, err := yaml.Marshal(parameter)
	if err != nil {
		return err
	}

	helmClient := helm.NewClient(helm.Host(tillerHost))
	_, err = helmClient.InstallRelease(chartPath, namespace, helm.ReleaseName(releaseName), helm.ValueOverrides(vals))
	if err != nil {
		glog.Infof("Failed to create drupal : %v \n\n", err)
		return err
	}

	return nil
}

func installMariaDB(releaseName, namespace string, parameter map[string]interface{}) error {

	parameter["mariadbRootPassword"] = uniuri.New()
	parameter["mariadbDatabase"] = "dbname"

	vals, err := yaml.Marshal(parameter)

	if err != nil {
		return err
	}

	helmClient := helm.NewClient(helm.Host(tillerHost))
	_, err = helmClient.InstallRelease(chartPath, namespace, helm.ReleaseName(releaseName), helm.ValueOverrides(vals))
	if err != nil {
		glog.Infof("Failed to create MariaDB : %v \n\n", err)
		return err
	}

	return nil
}

func installMySQL(releaseName, namespace string, parameter map[string]interface{}) error {

	parameter["mysqlRootPassword"] = uniuri.New()
	parameter["mysqlDatabase"] = "dbname"

	vals, err := yaml.Marshal(parameter)

	if err != nil {
		return err
	}

	helmClient := helm.NewClient(helm.Host(tillerHost))
	_, err = helmClient.InstallRelease(chartPath, namespace, helm.ReleaseName(releaseName), helm.ValueOverrides(vals))
	if err != nil {
		glog.Infof("Failed to create mysqlDB : %v \n\n", err)
		return err
	}

	return nil
}

// Install creates a new chart release
func Install(id, namespace string, parameter map[string]interface{}) error {

	/* *******TODO*******
	 * @TA : Only one instance will be supported for now per namespace as names will collide. Later we will need to track these.
	 */

	var err error

	switch id {
	case wordpress_id:
		err = mapAppInstallFunctions["wordpress"].(func(string, string, map[string]interface{}) error)(ReleaseName(id), namespace, parameter)
	case mariaDB_id:
		err = mapAppInstallFunctions["mariadb"].(func(string, string, map[string]interface{}) error)(ReleaseName(id), namespace, parameter)
	case drupal_id:
		err = mapAppInstallFunctions["drupal"].(func(string, string, map[string]interface{}) error)(ReleaseName(id), namespace, parameter)
	case mySQL_id:
		err = mapAppInstallFunctions["mysql"].(func(string, string, map[string]interface{}) error)(ReleaseName(id), namespace, parameter)
	default:
		glog.Info("Something Very Very Wrong")
		err = errors.New("Wrong Service ID Passed in Install function")
	}

	if err != nil {
		glog.Infof("Failed to create %s : %v \n\n", ReleaseName(id), err)
		return err
	}

	return nil
}

// Delete deletes a particular chart release
func Delete(id string) error {
	helmClient := helm.NewClient(helm.Host(tillerHost))
	if _, err := helmClient.DeleteRelease(ReleaseName(id)); err != nil {
		return err
	}
	return nil
}

// ReleaseName provides a string release isdentifier for gived id
func ReleaseName(id string) string {
	return mapRelNames[id]
}
