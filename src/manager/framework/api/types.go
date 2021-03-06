package api

import (
	"time"

	"github.com/Dataman-Cloud/swan/src/types"
)

type App struct {
	ID               string    `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Instances        int       `json:"instances,omitempty"`
	UpdatedInstances int       `json:"updatedInstances,omitempty"`
	RunningInstances int       `json:"runningInstances"`
	RunAs            string    `json:"runAs,omitempty"`
	Priority         int       `json:"priority"`
	ClusterID        string    `json:"clusterId,omitempty"`
	Status           string    `json:"status,omitempty"`
	Created          time.Time `json:"created,omitempty"`
	Updated          time.Time `json:"updated,omitempty"`
	Mode             string    `json:"mode,omitempty"`
	State            string    `json:"state"`

	// use task for compatability now, should be slot here
	Tasks          []*Task        `json:"tasks,omitempty"`
	CurrentVersion *types.Version `json:"currentVersion"`
	Versions       []string       `json:"versions,omitempty"`
	IP             []string       `json:"ip,omitempty"`

	// current version related info
	Labels      map[string]string `json:"labels,omitempty"`
	Env         map[string]string `json:"env,omitempty"`
	Constraints []string          `json:"constraints,omitempty"`
	Uris        []string          `json:"uris,omitempty"`
	//HealthChecks      []*types.HealthCheck
	//KillPolicy        *types.KillPolicy
	//UpdatePolicy      *types.UpdatePolicy
}

// use task for compatability now, should be slot here
// and together with task history
type Task struct {
	ID        string `json:"id,omitempty"`
	AppID     string `json:"appId,omitempty"`
	VersionID string `json:"versionId,omitempty"`

	Status string `json:"status"`

	OfferID       string `json:"offerId,omitempty"`
	AgentID       string `json:"agentId,omitempty"`
	AgentHostname string `json:"agentHostname,omitempty"`

	Cpu  float64 `json:"cpu,omitempty"`
	Mem  float64 `json:"mem,omitempty"`
	Disk float64 `json:"disk,omitempty"`

	History []*TaskHistory `json:"history,omitempty"`

	IP string `json:"ip,omitempty"`

	Created time.Time `json:"created,omitempty"`

	Image   string `json:"image,omitempty"`
	Healthy bool   `json:"healthy"`
}

type TaskHistory struct {
	ID        string `json:"id,omitempty"`
	AppID     string `json:"appId,omitempty"`
	VersionID string `json:"versionId,omitempty"`

	OfferID       string `json:"offerId,omitempty"`
	AgentID       string `json:"agentId,omitempty"`
	AgentHostname string `json:"agentHostname,omitempty"`

	Cpu  float64 `json:"cpu,omitempty"`
	Mem  float64 `json:"mem,omitempty"`
	Disk float64 `json:"disk,omitempty"`

	State  string `json:"state,omitempty"`
	Reason string `json:"Reason,omitempty"`
	Stdout string `json:"stdout,omitempty"`
	Stderr string `json:"stderr,omitempty"`
}

type Stats struct {
	AppCount  int `json:"appCount,omitempty"`
	TaskCount int `json:"taskCount,omitempty"`

	CpuTotalOffered  float64 `json:"cpuTotalOffered,omitempty"`
	MemTotalOffered  float64 `json:"memTotalOffered,omitempty"`
	DiskTotalOffered float64 `json:"diskTotalOffered,omitempty"`

	CpuTotalUsed  float64 `json:"cpuTotalUsed,omitempty"`
	MemTotalUsed  float64 `json:"memTotalUsed,omitempty"`
	DiskTotalUsed float64 `json:"diskTotalUsed,omitempty"`

	AppStats map[string]int `json:"appStats,omitempty"`
}

type ProceedUpdateParam struct {
	Instances int `json:"instances"`
}

type ScaleUpParam struct {
	Instances int      `json:"instances"`
	IPs       []string `json:"ips"`
}

type ScaleDownParam struct {
	Instances int `json:"instances"`
}
