package v1alpha1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// MigrationResourceName is name for "migration" resource
	MigrationResourceName = "migration"
	// MigrationResourcePlural is plural for "migration" resource
	MigrationResourcePlural = "migrations"
)

// MigrationSpec is the spec used to migrate apps between clusterpairs
type MigrationSpec struct {
	ClusterPair                  string            `json:"clusterPair"`
	AdminClusterPair             string            `json:"adminClusterPair"`
	Namespaces                   []string          `json:"namespaces"`
	NamespaceSelectors           map[string]string `json:"namespaceSelectors"`
	IncludeResources             *bool             `json:"includeResources"`
	IncludeVolumes               *bool             `json:"includeVolumes"`
	StartApplications            *bool             `json:"startApplications"`
	PurgeDeletedResources        *bool             `json:"purgeDeletedResources"`
	SkipServiceUpdate            *bool             `json:"skipServiceUpdate"`
	IncludeNetworkPolicyWithCIDR *bool             `json:"includeNetworkPolicyWithCIDR"`
	Selectors                    map[string]string `json:"selectors"`
	ExcludeSelectors             map[string]string `json:"excludeSelectors"`
	PreExecRule                  string            `json:"preExecRule"`
	PostExecRule                 string            `json:"postExecRule"`
	IncludeOptionalResourceTypes []string          `json:"includeOptionalResourceTypes"`
	SkipDeletedNamespaces        *bool             `json:"skipDeletedNamespaces"`
	TransformSpecs               []string          `json:"transformSpecs"`
	IgnoreOwnerReferencesCheck   *bool             `json:"ignoreOwnerReferencesCheck"`
}

// MigrationStatus is the status of a migration operation
type MigrationStatus struct {
	Stage                            MigrationStageType       `json:"stage"`
	Status                           MigrationStatusType      `json:"status"`
	Resources                        []*MigrationResourceInfo `json:"resources"`
	Volumes                          []*MigrationVolumeInfo   `json:"volumes"`
	FinishTimestamp                  meta.Time                `json:"finishTimestamp"`
	VolumeMigrationFinishTimestamp   meta.Time                `json:"volumeMigrationFinishTimestamp"`
	ResourceMigrationFinishTimestamp meta.Time                `json:"resourceMigrationFinishTimestamp"`
	// Summary provides a short summary on the migration
	Summary *MigrationSummary `json:"summary"`
}

// MigrationResourceInfo is the info for the migration of a resource
type MigrationResourceInfo struct {
	Name                  string `json:"name"`
	Namespace             string `json:"namespace"`
	meta.GroupVersionKind `json:",inline"`
	Status                MigrationStatusType `json:"status"`
	Reason                string              `json:"reason"`
	TransformedBy         string              `json:"transformedBy"`
}

// MigrationSummary provides a short summary on the migration
type MigrationSummary struct {
	// TotalNumberOfVolumes gives the total count of volumes
	TotalNumberOfVolumes uint64 `json:"totalNumberOfVolumes"`
	// NumberOfMigratedVolumes gives the total count of successfully migrated volumes
	NumberOfMigratedVolumes uint64 `json:"numOfMigratedVolumes"`
	// TotalNumberOfResources gives the total count of resourcs
	TotalNumberOfResources uint64 `json:"totalNumberOfResources"`
	// NumberOfMigratedResources gives the total count of migrated k8s resources
	NumberOfMigratedResources uint64 `json:"numOfMigratedResources"`
	// TotalBytesMigrated gives the total amount of bytes migrated across all the volumes
	TotalBytesMigrated uint64 `json:"totalBytesMigrated"`
	// ElapsedTimeForVolumeMigration provides the total time the
	// volume migration stage has been running or the total time
	// taken for the volume migration to complete if the volume migration has finished
	ElapsedTimeForVolumeMigration string `json:"elapsedTimeForVolumeMigration"`
	// ElapsedTimeForResourceMigration provides the total time the
	// resource migration stage has been running or the total time
	// taken for the resource migration to complete if the volume migration has finished
	ElapsedTimeForResourceMigration string `json:"elapsedTimeForResourceMigration"`
}

// MigrationVolumeInfo is the info for the migration of a volume
type MigrationVolumeInfo struct {
	PersistentVolumeClaim string              `json:"persistentVolumeClaim"`
	Namespace             string              `json:"namespace"`
	Volume                string              `json:"volume"`
	Status                MigrationStatusType `json:"status"`
	BytesTotal            uint64              `json:"bytesTotal"`
	Reason                string              `json:"reason"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Migration represents migration status
type Migration struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`
	Spec            MigrationSpec   `json:"spec"`
	Status          MigrationStatus `json:"status"`
}

// MigrationStatusType is the status of the migration
type MigrationStatusType string

const (
	// MigrationStatusInitial is the initial state when migration is created
	MigrationStatusInitial MigrationStatusType = ""
	// MigrationStatusPending for when migration is still pending
	MigrationStatusPending MigrationStatusType = "Pending"
	// MigrationStatusInProgress for when migration is in progress
	MigrationStatusInProgress MigrationStatusType = "InProgress"
	// MigrationStatusFailed for when migration has failed
	MigrationStatusFailed MigrationStatusType = "Failed"
	// MigrationStatusPartialSuccess for when migration was partially successful
	MigrationStatusPartialSuccess MigrationStatusType = "PartialSuccess"
	// MigrationStatusSuccessful for when migration has completed successfully
	MigrationStatusSuccessful MigrationStatusType = "Successful"
	// MigrationStatusPurged for when migration objects has been deleted
	MigrationStatusPurged MigrationStatusType = "Purged"
)

// MigrationStageType is the stage of the migration
type MigrationStageType string

const (
	// MigrationStageInitial for when migration is created
	MigrationStageInitial MigrationStageType = ""
	// MigrationStagePreExecRule for when the PreExecRule is being executed
	MigrationStagePreExecRule MigrationStageType = "PreExecRule"
	// MigrationStagePostExecRule for when the PostExecRule is being executed
	MigrationStagePostExecRule MigrationStageType = "PostExecRule"
	// MigrationStageVolumes for when volumes are being migrated
	MigrationStageVolumes MigrationStageType = "Volumes"
	// MigrationStageApplications for when applications are being migrated
	MigrationStageApplications MigrationStageType = "Applications"
	// MigrationStageFinal is the final stage for migration
	MigrationStageFinal MigrationStageType = "Final"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MigrationList is a list of Migrations
type MigrationList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []Migration `json:"items"`
}
