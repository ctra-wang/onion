package model

// RepositoryParams 结构体定义
type RepositoryParams struct {
	OrgName        string `json:"orgName"`        // 组织名称
	RepoName       string `json:"repoName"`       // 仓库名称
	ApiToken       string `json:"apiToken"`       // API 访问令牌
	RepoPermission bool   `json:"repoPermission"` // 仓库权限 true：private 私有，false：public 公开
	Platform       int    `json:"platform"`       // 平台 1、Github，2、Gitee，3、Gitea，4、Gitlab
}
