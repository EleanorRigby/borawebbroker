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
	URL      string
}

const (
	tillerHost = "tiller-deploy.kube-system.svc.cluster.local:44134"

	chartPath      = "/charts"
	wordpressChart = chartPath + "/wordpress"
	drupalChart    = chartPath + "/drupal"
	mariaDBChart   = chartPath + "/mariadb"
	mySQLChart     = chartPath + "/mysql"

	Wordpress_id      = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2468"
	Wordpress_plan_id = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2469"
	Drupal_id         = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2460"
	Drupal_plan_id    = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2461"
	MariaDB_id        = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2462"
	MariaDB_plan_id   = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2463"
	MySQL_id          = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2464"
	MySQL_plan_id     = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2465"
)

//This below is a very unsafe way to store release names.
//What if broker crashes.
//We have no way of recovering id-->releasename relation.
//We can potentially return the release name in response to orchestrator
// Orchestrator can deploy an etcd to store these safely.
type instanceDetails struct {
	name      string
	namespace string
}

var count int = 1

var namedetails = make(map[string]*instanceDetails)

var mapRelNames = map[string]string{
	//release names
	"4f6e6cf6-ffdd-425f-a2c7-3c9258ad2468": "whackywordpress",
	"4f6e6cf6-ffdd-425f-a2c7-3c9258ad2460": "drogodrupal",
	"4f6e6cf6-ffdd-425f-a2c7-3c9258ad2462": "myriadmaria",
	"4f6e6cf6-ffdd-425f-a2c7-3c9258ad2464": "supersql",
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

	var creds DBCreds

	creds.Host = namedetails[id].name + "-drupal." + namedetails[id].namespace + ".svc.cluster.local"
	creds.Port = "80"
	creds.Username = "user"

	config, err := rest.InClusterConfig()
	if err != nil {
		return creds, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return creds, err
	}
	secret, err := clientset.Core().Secrets(namedetails[id].namespace).Get(namedetails[id].name + "-drupal")
	if err != nil {
		return creds, err
	}

	service, err := clientset.Core().Services(namedetails[id].namespace).Get(namedetails[id].name + "-drupal")
	hostname := service.Status.LoadBalancer.Ingress[0].Hostname
	creds.URL = "http://" + hostname
	creds.Password = string(secret.Data["drupal-password"])

	glog.Info("Debug data for binding credentials for %s are  %v", namedetails[id].name, creds)

	return creds, nil

}

func bindMariaDB(id string) (DBCreds, error) {

	var creds DBCreds

	creds.Host = namedetails[id].name + "-mariadb." + namedetails[id].namespace + ".svc.cluster.local"
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
	secret, err := clientset.Core().Secrets(namedetails[id].namespace).Get(namedetails[id].name + "-mariadb")
	if err != nil {
		return creds, err
	}

	creds.Password = string(secret.Data["mariadb-root-password"])
	creds.Uri = "mysql://" + creds.Username + ":" + creds.Password + "@" + creds.Host + ":" + creds.Port + "/" + creds.Database

	glog.Info("Debug data for binding credentials for %s are  %v", namedetails[id].name, creds)

	return creds, nil

}

func bindMySQL(id string) (DBCreds, error) {

	var creds DBCreds

	creds.Host = namedetails[id].name + "-mysql." + namedetails[id].namespace + ".svc.cluster.local"
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
	secret, err := clientset.Core().Secrets(namedetails[id].namespace).Get(namedetails[id].name + "-mariadb")
	if err != nil {
		return creds, err
	}

	creds.Password = string(secret.Data["mysql-root-password"])
	creds.Uri = "mysql://" + creds.Username + ":" + creds.Password + "@" + creds.Host + ":" + creds.Port + "/" + creds.Database

	glog.Info("Debug data for binding credentials for %s are  %v", namedetails[id].name, creds)

	return creds, nil

}

// GetBinding returns credential for the passed instance ID
func GetBinding(sid, id string) (DBCreds, error) {

	var err error
	var creds DBCreds

	switch id {
	case Wordpress_id:
		creds, err = mapAppBindFunctions["wordpress"].(func(string) (DBCreds, error))(sid)
	case MariaDB_id:
		creds, err = mapAppBindFunctions["mariadb"].(func(string) (DBCreds, error))(sid)
	case Drupal_id:
		creds, err = mapAppBindFunctions["drupal"].(func(string) (DBCreds, error))(sid)
	case MySQL_id:
		creds, err = mapAppBindFunctions["mysql"].(func(string) (DBCreds, error))(sid)
	default:
		glog.Info("Something Very Very Wrong")
		err = errors.New("Wrong Service ID Passed in Bind function")
	}

	if err != nil {
		glog.Infof("Failed to create binding for %s : %v \n\n", namedetails[id].name, err)
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
	glog.Infof("%v   %v   %v   %v   ", wordpressChart, namespace, releaseName, vals)
	_, err = helmClient.InstallRelease(wordpressChart, namespace, helm.ReleaseName(releaseName), helm.ValueOverrides(vals))
	if err != nil {
		glog.Infof("Failed to create wordpress : %v \n\n", err)
		return err
	}

	return nil
}

func installDrupal(releaseName, namespace string, parameter map[string]interface{}) error {

	if _, ok := parameter["drupalPassword"]; !ok {
		parameter["drupalPassword"] = uniuri.New()
	}

	vals, err := yaml.Marshal(parameter)
	if err != nil {
		return err
	}

	helmClient := helm.NewClient(helm.Host(tillerHost))
	_, err = helmClient.InstallRelease(drupalChart, namespace, helm.ReleaseName(releaseName), helm.ValueOverrides(vals))
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
	_, err = helmClient.InstallRelease(mariaDBChart, namespace, helm.ReleaseName(releaseName), helm.ValueOverrides(vals))
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
	_, err = helmClient.InstallRelease(mySQLChart, namespace, helm.ReleaseName(releaseName), helm.ValueOverrides(vals))
	if err != nil {
		glog.Infof("Failed to create mysqlDB : %v \n\n", err)
		return err
	}

	return nil
}

// Install creates a new chart release
func Install(sid, id, namespace string, parameter map[string]interface{}) error {

	/* *******TODO*******
	 * @TA : Only one instance will be supported for now per namespace as names will collide. Later we will need to track these.
	 */

	var err error

	nameofrelease := ReleaseName(sid) + string(count)
	count = count + 1
	namedetails[id] = &instanceDetails{}
	namedetails[id].name = nameofrelease
	namedetails[id].namespace = namespace

	glog.Infof("Debug : The value of map being stored is %v", namedetails[id])

	switch sid {
	case Wordpress_id:
		err = mapAppInstallFunctions["wordpress"].(func(string, string, map[string]interface{}) error)(nameofrelease, namespace, parameter)
	case MariaDB_id:
		err = mapAppInstallFunctions["mariadb"].(func(string, string, map[string]interface{}) error)(nameofrelease, namespace, parameter)
	case Drupal_id:
		err = mapAppInstallFunctions["drupal"].(func(string, string, map[string]interface{}) error)(nameofrelease, namespace, parameter)
	case MySQL_id:
		err = mapAppInstallFunctions["mysql"].(func(string, string, map[string]interface{}) error)(nameofrelease, namespace, parameter)
	default:
		glog.Info("Something Very Very Wrong")
		err = errors.New("Wrong Service ID Passed in Install function")
	}

	if err != nil {
		glog.Infof("Failed to create %s : %v \n\n", namedetails[id].name, err)
		return err
	}

	return nil
}

// Delete deletes a particular chart release
func Delete(id string) error {
	helmClient := helm.NewClient(helm.Host(tillerHost))
	releasename := namedetails[id].name
	glog.Infof("Debug : Release name to be removed %s", releasename)
	if _, err := helmClient.DeleteRelease(releasename); err != nil {
		return err
		glog.Info("This effing error %v", err)
	}

	return nil
}

// ReleaseName provides a string release isdentifier for gived id
func ReleaseName(id string) string {
	return mapRelNames[id]
}
