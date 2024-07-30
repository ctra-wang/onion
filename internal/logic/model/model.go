package model

// RepositoryParams 结构体定义
type RepositoryParams struct {
	OrgName        string `json:"orgName"`        // 组织名称
	RepoName       string `json:"repoName"`       // 仓库名称
	ApiToken       string `json:"apiToken"`       // API 访问令牌
	RepoPermission bool   `json:"repoPermission"` // 仓库权限 true：private 私有，false：public 公开
	Platform       int    `json:"platform"`       // 平台 1、Github，2、Gitee，3、Gitea，4、Gitlab
}

// RpcParams Rpc 结构体定义
type RpcParams struct {
	RpcName    string `json:"rpcName"`    // 组织名称
	ModuleName string `json:"moduleName"` // 仓库名称 如：gitee.com/ajiot_vpp/vpp-go-ark-data-asset-rpc.git
	Port       int32  `json:"port"`       // 端口
}

// DatabaseConf 数据库 结构体定义
type DatabaseConf struct {
	Host         string `json:",env=DATABASE_HOST"`
	Port         int    `json:",env=DATABASE_PORT"`
	TableName    string `json:",optional,env=DATABASE_TABLENAME"`
	Username     string `json:",default=root,env=DATABASE_USERNAME"`
	Password     string `json:",optional,env=DATABASE_PASSWORD"`
	DBName       string `json:",default=simple_admin,env=DATABASE_DBNAME"`
	SSLMode      string `json:",optional,env=DATABASE_SSL_MODE"`
	Type         string `json:",default=mysql,options=[mysql,postgres,sqlite3],env=DATABASE_TYPE"`
	MaxOpenConn  int    `json:",optional,default=100,env=DATABASE_MAX_OPEN_CONN"`
	CacheTime    int    `json:",optional,default=10,env=DATABASE_CACHE_TIME"`
	DBPath       string `json:",optional,env=DATABASE_DBPATH"`
	MysqlConfig  string `json:",optional,env=DATABASE_MYSQL_CONFIG"`
	PGConfig     string `json:",optional,env=DATABASE_PG_CONFIG"`
	SqliteConfig string `json:",optional,env=DATABASE_SQLITE_CONFIG"`
}
