package response

import "time"

type Repository struct {
	ArtifactCount int       `json:"artifact_count"`
	CreationTime  time.Time `json:"creation_time"`
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	ProjectID     int       `json:"project_id"`
	PullCount     int       `json:"pull_count"`
	UpdateTime    time.Time `json:"update_time"`
}

type Project []struct {
	CreationTime       time.Time    `json:"creation_time"`
	CurrentUserRoleIds interface{}  `json:"current_user_role_ids"`
	CveAllowlist       CveAllowlist `json:"cve_allowlist"`
	Metadata           Metadata     `json:"metadata"`
	Name               string       `json:"name"`
	OwnerID            int          `json:"owner_id"`
	OwnerName          string       `json:"owner_name"`
	ProjectID          int          `json:"project_id"`
	RepoCount          int          `json:"repo_count"`
	UpdateTime         time.Time    `json:"update_time"`
}
type CveAllowlist struct {
	CreationTime time.Time     `json:"creation_time"`
	ID           int           `json:"id"`
	Items        []interface{} `json:"items"`
	ProjectID    int           `json:"project_id"`
	UpdateTime   time.Time     `json:"update_time"`
}
type Metadata struct {
	Public string `json:"public"`
}
