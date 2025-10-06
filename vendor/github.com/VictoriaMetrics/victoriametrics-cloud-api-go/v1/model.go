package v1

import (
	"time"
)

// DeploymentType - type of the deployment (single_node / cluster)
type DeploymentType string

const (
	// DeploymentTypeSingleNode - single node deployment
	DeploymentTypeSingleNode DeploymentType = "single_node"
	// DeploymentTypeCluster - cluster deployment
	DeploymentTypeCluster DeploymentType = "cluster"
)

func (t DeploymentType) String() string {
	return string(t)
}

// DeploymentCloudProvider - cloud provider type for deployment
type DeploymentCloudProvider string

const (
	// DeploymentCloudProviderAWS - AWS cloud provider
	DeploymentCloudProviderAWS DeploymentCloudProvider = "aws"
)

func (p DeploymentCloudProvider) String() string {
	return string(p)
}

// DeploymentStatus - current status of the deployment
type DeploymentStatus string

const (
	DeploymentStatusProvisioning DeploymentStatus = "PROVISIONING"
	DeploymentStatusRunning      DeploymentStatus = "RUNNING"
	DeploymentStatusError        DeploymentStatus = "ERROR"
	DeploymentStatusStopped      DeploymentStatus = "STOPPED"
)

func (s DeploymentStatus) String() string {
	return string(s)
}

// DurationUnit represents a unit of time
type DurationUnit string

const (
	// DurationUnitMillisecond - milliseconds
	DurationUnitMillisecond DurationUnit = "ms"
	// DurationUnitSecond - seconds
	DurationUnitSecond DurationUnit = "s"
	// DurationUnitDay - days
	DurationUnitDay DurationUnit = "d"
	// DurationUnitMonth - months
	DurationUnitMonth DurationUnit = "m"
)

// AccessMode defines token access mode
type AccessMode string

const (
	// AccessModeRead - read-only access mode
	AccessModeRead AccessMode = "r"
	// AccessModeWrite - write-only access mode
	AccessModeWrite AccessMode = "w"
	// AccessModeReadWrite - read+write access mode
	AccessModeReadWrite AccessMode = "rw"
)

func (a AccessMode) String() string {
	return string(a)
}

// StorageUnit - storage unit type (GB / TB)
type StorageUnit string

const (
	// StorageUnitGB - gigabyte storage unit
	StorageUnitGB StorageUnit = "GB"
	// StorageUnitTB - terabyte storage unit
	StorageUnitTB StorageUnit = "TB"
)

// MaintenanceWindow - maintenance window for the deployment
type MaintenanceWindow string

const (
	// MaintenanceWindowWeekendDays - maintenance window on weekdays
	MaintenanceWindowWeekendDays MaintenanceWindow = "Sat-Sun 3-4am"
	// MaintenanceWindowBusinessDays - maintenance window on business days
	MaintenanceWindowBusinessDays MaintenanceWindow = "Mon-Fri 4-5am"
)

// FlagList - list of command-line flags
type FlagList []string

// DeploymentFlags - Customized command-line flags for the deployment
type DeploymentFlags struct {
	// CommonFlags - Common command-line flags for vmsingle component
	SingleFlags FlagList `json:"single_flags"`
	// SelectFlags - Customized command-line flags for the vmselect component
	SelectFlags FlagList `json:"select_flags"`
	// StorageFlags - Customized command-line flags for the vmstorage component
	StorageFlags FlagList `json:"storage_flags"`
	// InsertFlags - Customized command-line flags for the vminsert component
	InsertFlags FlagList `json:"insert_flags"`
}

func (u StorageUnit) String() string {
	return string(u)
}

// TierInfo represents the information about the tier in public VMCloud API
type TierInfo struct {
	// ID is the unique identifier of the tier of given type
	ID uint32 `json:"id"`
	// Type of the deployment (single_node / cluster)
	Type DeploymentType `json:"type"`
	// CloudProvider is the name of Cloud provider of the deployment (aws)
	CloudProvider DeploymentCloudProvider `json:"cloud_provider"`
	// Name is the name of the tier
	Name string `json:"name"`
	// ComputeCostPerHour is the cost of the deployment per hour
	ComputeCostPerHour float64 `json:"compute_cost_per_hour"`
	// IngestionRate is the maximum ingestion rate of the tier
	IngestionRate int `json:"ingestion_rate"`
	// ActiveTimeSeries is the maximum number of active time series of the tier
	ActiveTimeSeries int `json:"active_time_series"`
	// NewSeriesOver24h is the maximum number of new series over 24 hours of the tier
	NewSeriesOver24h int `json:"new_series_over_24h"`
	// DatapointsReadRate is the maximum read rate of the tier
	DatapointsReadRate int `json:"datapoints_read_rate"`
	// SeriesReadPerQuery is the maximum number of series read per query of the tier
	SeriesReadPerQuery int `json:"series_read_per_query"`
	// AccessTokenLimit is the maximum number of concurrent requests for each access token
	AccessTokenConcurrentRequests int `json:"access_token_concurrent_requests"`
}

// TierInfoList represents the list of TierInfo
type TierInfoList []TierInfo

// DeploymentSummary - simplified representation of deployment for list view
type DeploymentSummary struct {
	// ID - unique identifier of the deployment
	ID string `json:"id"`
	// Name - human-readable name of the deployment
	Name string `json:"name"`
	// Type of the deployment (single_node / cluster)
	Type DeploymentType `json:"type"`
	// Tier - tier identifier of the deployment
	Tier uint32 `json:"tier"`
	// Version of VictoriaMetrics used in the deployment
	Version string `json:"version"`
	// CloudProvider - ID of then cloud provider of the deployment
	CloudProvider DeploymentCloudProvider `json:"cloud_provider"`
	// Region of the deployment in specified cloud provider
	Region string `json:"region"`
	// CreatedAt - timestamp of deployment creation
	CreatedAt time.Time `json:"created_at"`
	// Status - current status of the deployment
	Status DeploymentStatus `json:"status"`
}

// DeploymentSummaryList represents the list of DeploymentSummary
type DeploymentSummaryList []DeploymentSummary

// DeploymentPrice - price of the deployment (USD per month)
type DeploymentPrice struct {
	// ComputeCost - cost of the compute resources
	ComputeCost float64 `json:"compute_cost"`
	// StorageCost - cost of the storage resources
	StorageCost float64 `json:"storage_cost"`
	// TotalCost - total cost of the deployment (without network costs)
	TotalCost float64 `json:"total_cost"`
}

// DeploymentInfo - deployment details for the public API
type DeploymentInfo struct {
	// ID - unique identifier of the deployment
	ID string `json:"id"`
	// Name - human-readable name of the deployment
	Name string `json:"name"`
	// Type of the deployment (single_node / cluster)
	Type DeploymentType `json:"type"`
	// Tier - tier identifier of the deployment
	Tier uint32 `json:"tier"`
	// Version of VictoriaMetrics used in the deployment
	Version string `json:"version"`
	// CloudProvider - ID of then cloud provider of the deployment
	CloudProvider DeploymentCloudProvider `json:"cloud_provider"`
	// Region of the deployment in specified cloud provider
	Region string `json:"region"`
	// CreatedAt - timestamp of deployment creation
	CreatedAt time.Time `json:"created_at"`
	// Status - current status of the deployment
	Status DeploymentStatus `json:"status"`
	// RetentionValue - retention period of the deployment
	RetentionValue uint32 `json:"retention_value"`
	// RetentionUnit - retention period unit of the deployment
	RetentionUnit DurationUnit `json:"retention_unit"`
	// DeduplicationValue - deduplication period of the deployment
	DeduplicationValue uint32 `json:"deduplication_value"`
	// DeduplicationUnit - deduplication period unit of the deployment
	DeduplicationUnit DurationUnit `json:"deduplication_unit"`
	// StorageSizeGb - storage size of the deployment
	StorageSizeGb uint64 `json:"storage_size_gb"`
	// MaintenanceWindow - maintenance window of the deployment
	MaintenanceWindow MaintenanceWindow `json:"maintenance_window"`
	// Price - price of the deployment
	Price DeploymentPrice `json:"price"`
	// VMSingleSettings - settings for single-node deployment
	VMSingleSettings []string `json:"vmsingle_settings,omitempty"`
	// VMStorageSettings - VMStorage settings for cluster deployment
	VMStorageSettings []string `json:"vmstorage_settings,omitempty"`
	// VMSelectSettings - VMSelect settings for cluster deployment
	VMSelectSettings []string `json:"vmselect_settings,omitempty"`
	// VMInsertSettings - VMInsert settings for cluster deployment
	VMInsertSettings []string `json:"vminsert_settings,omitempty"`
	// AccessEndpoint - endpoint of the deployment (URL entrypoint to API of the deployment)
	AccessEndpoint string `json:"access_endpoint"`
}

// DeploymentInfoList represents the list of DeploymentInfo
type DeploymentInfoList []DeploymentInfo

// DeploymentCreationRequest represents the request for creating a deployment
type DeploymentCreationRequest struct {
	// Name - human-readable name of the deployment
	Name string `json:"name"`
	// Type of the deployment (single_node / cluster)
	Type DeploymentType `json:"type"`
	// Provider - cloud provider of the deployment
	Provider DeploymentCloudProvider `json:"provider"`
	// Region of the deployment in specified cloud provider
	Region string `json:"region"`
	// Tier - tier identifier of the deployment
	Tier uint32 `json:"tier"`
	// StorageSize - storage size in units specified in StorageSizeUnit
	StorageSize uint64 `json:"storage_size"`
	// StorageSizeUnit - storage size unit (GB / TB)
	StorageSizeUnit StorageUnit `json:"storage_size_unit"`
	// Deduplication window for the deployment in units specified in DeduplicationUnit
	Deduplication uint32 `json:"deduplication"`
	// DeduplicationUnit - deduplication window unit for the deployment
	DeduplicationUnit DurationUnit `json:"deduplication_unit"`
	// Retention period for the deployment in units specified in RetentionUnit
	Retention uint32 `json:"retention"`
	// RetentionUnit - retention period unit for the deployment
	RetentionUnit DurationUnit `json:"retention_unit"`
	// MaintenanceWindow - maintenance window for the deployment
	MaintenanceWindow MaintenanceWindow `json:"maintenance_window"`
}

// DeploymentUpdateRequest represents the request for updating a deployment
type DeploymentUpdateRequest = struct {
	// Name - human-readable name of the deployment
	Name string `json:"name"`
	// Tier - tier identifier of the deployment
	Tier uint32 `json:"tier"`
	// StorageSize - storage size in units specified in StorageSizeUnit
	StorageSize uint64 `json:"storage_size"`
	// StorageSizeUnit - storage size unit (GB / TB)
	StorageSizeUnit StorageUnit `json:"storage_size_unit"`
	// Deduplication window for the deployment in units specified in DeduplicationUnit
	Deduplication uint32 `json:"deduplication"`
	// DeduplicationUnit - deduplication window unit for the deployment
	DeduplicationUnit DurationUnit `json:"deduplication_unit"`
	// Retention period for the deployment in units specified in RetentionUnit
	Retention uint32 `json:"retention"`
	// RetentionUnit - retention period unit for the deployment
	RetentionUnit DurationUnit `json:"retention_unit"`
	// MaintenanceWindow - maintenance window for the deployment
	MaintenanceWindow MaintenanceWindow `json:"maintenance_window"`
	// Flags - customized command-line flags for the deployment
	Flags DeploymentFlags `json:"flags"`
}

// AccessTokensList represents the list of AccessToken
type AccessTokensList []AccessToken

// AccessToken represents the access token in public VMCloud API
type AccessToken struct {
	// ID is the unique identifier of the access token
	ID string `json:"id"`
	// Secret is the secret value of the access token (only 4 symbols are returned in access tokens list, for full secret use reveal endpoint)
	Secret string `json:"value"`
	// Type is the access mode of the token (read-only, write-only, read+write)
	Type AccessMode `json:"type"`
	// Description is the human-readable description of the access token
	Description string `json:"description"`
	// CreatedBy is the user who created the access token
	CreatedBy string `json:"created_by"`
	// CreatedAt is the timestamp of the access token creation
	CreatedAt time.Time `json:"created_at"`
	// TenantID represents the unique identifier of the tenant associated with this access token (optional)
	TenantID string `json:"tenant_id,omitempty"`
	// Timestamp of the last usage of the access token (within the last 7 days)
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
}

// AccessTokenCreateRequest represents the request for creating an access token
type AccessTokenCreateRequest struct {
	// Type is the access mode of the token (read-only, write-only, read+write)
	Type AccessMode `json:"type"`
	// Description is the human-readable description of the access token
	Description string `json:"description"`
	// TenantID represents the unique identifier of the tenant associated with this access token (optional)
	TenantID string `json:"tenant_id,omitempty"`
}

// RegionInfo represents the information about the region in public VMCloud API
type RegionInfo struct {
	// Provider is the name of the cloud provider
	CloudProvider DeploymentCloudProvider `json:"cloud_provider"`
	// Name is the name of the region
	Name string `json:"name"`
}

// RegionInfoList represents the list of RegionInfo
type RegionInfoList []RegionInfo

// CloudProviderInfo represents the information about the cloud provider in public VMCloud API
type CloudProviderInfo struct {
	// ID is the unique identifier of the cloud provider
	ID DeploymentCloudProvider `json:"id"`
	// URL is the URL of the cloud provider
	URL string `json:"url"`
}

// CloudProviderInfoList represents the list of CloudProviderInfo
type CloudProviderInfoList []CloudProviderInfo
