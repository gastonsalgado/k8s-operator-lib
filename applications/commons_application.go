package applications

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymentType string

const (
	Standard    DeploymentType = "Standard"
	MultiRegion DeploymentType = "MultiRegion"
	MultiCloud  DeploymentType = "MultiCloud"
)

type Phase string

const (
	Rollout  Phase = "Rollout"
	Rollback Phase = "Rollback"
	Failover Phase = "Failover"
	Fallback Phase = "Fallback"
)

type Role string

const (
	Primary   Role = "Primary"
	Secondary Role = "Secondary"
)

type CurrentSyncStatus string

const (
	OutOfSync  CurrentSyncStatus = "OutOfSync"
	Syncing    CurrentSyncStatus = "Syncing"
	Synced     CurrentSyncStatus = "Synced"
	SyncFailed CurrentSyncStatus = "SyncFailed"
)

type LastSyncResult string

const (
	SyncOk    LastSyncResult = "SyncOk"
	SyncError LastSyncResult = "SyncError"
)

type Health string

const (
	Missing     Health = "Missing"
	Progressing Health = "Progressing"
	Healthy     Health = "Healthy"
	Degraded    Health = "Degraded"
)

type AppStatus struct {
	Phase              Phase             `json:"phase,omitempty"`
	CurrentSyncStatus  CurrentSyncStatus `json:"currentSyncStatus,omitempty"`
	LastSyncResult     LastSyncResult    `json:"lastSyncResult,omitempty"`
	Health             Health            `json:"health,omitempty"`
	LastDeploymentTime *metav1.Time      `json:"lastDeploymentTime,omitempty" protobuf:"bytes,6,opt,name=lastDeploymentTime"`
}

type Application interface {
	InitStatus()
	GetSpec() interface{}
	GetAppStatus() AppStatus
	GetConditionsStatus() *[]metav1.Condition
	GetStatus(string) string
	SetSpec(interface{})
	SetAppStatus(AppStatus)
	SetStatus(string, string)
}
