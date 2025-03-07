package applicationmanager

import (
	"os"
	"reflect"
	"time"

	"github.com/libopenstorage/stork/drivers/volume"
	stork_crd "github.com/libopenstorage/stork/pkg/apis/stork"
	stork_api "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	"github.com/libopenstorage/stork/pkg/applicationmanager/controllers"
	"github.com/libopenstorage/stork/pkg/k8sutils"
	"github.com/libopenstorage/stork/pkg/resourcecollector"
	"github.com/libopenstorage/stork/pkg/version"
	"github.com/portworx/sched-ops/k8s/apiextensions"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

const (
	validateCRDInterval time.Duration = 5 * time.Second
	validateCRDTimeout  time.Duration = 1 * time.Minute
)

// ApplicationManager maintains all controllers for application level operations
type ApplicationManager struct {
	Driver            volume.Driver
	Recorder          record.EventRecorder
	ResourceCollector resourcecollector.ResourceCollector
	RsyncTime         int64
}

// Init Initializes the ApplicationManager and any children controller
func (a *ApplicationManager) Init(mgr manager.Manager, adminNamespace string, stopChannel chan os.Signal) error {
	backupController := controllers.NewApplicationBackup(mgr, a.Recorder, a.ResourceCollector)
	if err := backupController.Init(mgr, adminNamespace, a.RsyncTime); err != nil {
		return err
	}

	restoreController := controllers.NewApplicationRestore(mgr, a.Recorder, a.ResourceCollector)
	if err := restoreController.Init(mgr, adminNamespace); err != nil {
		return err
	}

	cloneController := controllers.NewApplicationClone(mgr, a.Driver, a.Recorder, a.ResourceCollector)
	if err := cloneController.Init(mgr, adminNamespace); err != nil {
		return err
	}

	scheduleController := controllers.NewApplicationBackupSchedule(mgr, a.Recorder)
	if err := scheduleController.Init(mgr); err != nil {
		return err
	}
	syncController := &controllers.BackupSyncController{
		Recorder:     a.Recorder,
		SyncInterval: 1 * time.Minute,
	}
	if err := syncController.Init(stopChannel); err != nil {
		return err
	}

	if err := controllers.RegisterDefaultCRDs(); err != nil {
		return err
	}
	return nil
}

func CreateCRD() error {
	resource := apiextensions.CustomResource{
		Name:    stork_api.BackupLocationResourceName,
		Plural:  stork_api.BackupLocationResourcePlural,
		Group:   stork_api.SchemeGroupVersion.Group,
		Version: stork_api.SchemeGroupVersion.Version,
		Scope:   apiextensionsv1beta1.NamespaceScoped,
		Kind:    reflect.TypeOf(stork_api.BackupLocation{}).Name(),
	}
	appReg := apiextensions.CustomResource{
		Name:    stork_api.ApplicationRegistrationResourceName,
		Plural:  stork_api.ApplicationRegistrationResourcePlural,
		Group:   stork_crd.GroupName,
		Version: stork_api.SchemeGroupVersion.Version,
		Scope:   apiextensionsv1beta1.ClusterScoped,
		Kind:    reflect.TypeOf(stork_api.ApplicationRegistration{}).Name(),
	}
	platformCred := apiextensions.CustomResource{
		Name:    stork_api.PlatformCredentialResourceName,
		Plural:  stork_api.PlatformCredentialResourcePlural,
		Group:   stork_crd.GroupName,
		Version: stork_api.SchemeGroupVersion.Version,
		Scope:   apiextensionsv1beta1.NamespaceScoped,
		Kind:    reflect.TypeOf(stork_api.PlatformCredential{}).Name(),
	}

	ok, err := version.RequiresV1Registration()
	if err != nil {
		return err
	}
	if ok {
		err := k8sutils.CreateCRDV1(resource)
		if err != nil && !errors.IsAlreadyExists(err) {
			return err
		}
		if err := apiextensions.Instance().ValidateCRD(resource.Plural+"."+resource.Group, validateCRDTimeout, validateCRDInterval); err != nil {
			return err
		}
		err = k8sutils.CreateCRDV1(platformCred)
		if err != nil && !errors.IsAlreadyExists(err) {
			return err
		}
		if err := apiextensions.Instance().ValidateCRD(platformCred.Plural+"."+platformCred.Group, validateCRDTimeout, validateCRDInterval); err != nil {
			return err
		}
		err = k8sutils.CreateCRDV1(appReg)
		if err != nil && !errors.IsAlreadyExists(err) {
			return err
		}
		return apiextensions.Instance().ValidateCRD(appReg.Plural+"."+appReg.Group, validateCRDTimeout, validateCRDInterval)
	}
	err = apiextensions.Instance().CreateCRDV1beta1(resource)
	if err != nil && !errors.IsAlreadyExists(err) {
		return err
	}
	if err := apiextensions.Instance().ValidateCRDV1beta1(resource, validateCRDTimeout, validateCRDInterval); err != nil {
		return err
	}
	err = apiextensions.Instance().CreateCRDV1beta1(platformCred)
	if err != nil && !errors.IsAlreadyExists(err) {
		return err
	}
	if err := apiextensions.Instance().ValidateCRDV1beta1(platformCred, validateCRDTimeout, validateCRDInterval); err != nil {
		return err
	}
	err = apiextensions.Instance().CreateCRDV1beta1(appReg)
	if err != nil && !errors.IsAlreadyExists(err) {
		return err
	}
	return apiextensions.Instance().ValidateCRDV1beta1(appReg, validateCRDTimeout, validateCRDInterval)
}
