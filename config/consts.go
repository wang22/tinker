package config

const (
	// HTTP Config
	ConstHttpPort        = "http.port"
	ConstHttpAdminPrefix = "http.admin.prefix"

	// DB Config
	ConstDBType        = "db.type"
	ConstDBName        = "db.dbname"
	ConstDBPort        = "db.port"
	ConstDBHost        = "db.host"
	ConstDBUsername    = "db.username"
	ConstDBPassword    = "db.password"
	ConstDBTablePrefix = "db.tablePrefix"
	ConstDBLogEnable   = "db.logEnable"

	// Cache Config
	ConstCacheType = "cache.type"
	ConstCacheSize = "cache.size"

	// Cookie Config
	ConstAuthCookieKey = "http.authCookieKey"
)
