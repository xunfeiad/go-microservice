package registry

type Registration struct {
	ServiceName       ServiceName
	ServiceURL        string
	RequiredSerivices []ServiceName
	ServiceUpdateURL  string // 通过这个url告诉哪些服务可以用
	HeartbeatURL      string // 心跳检查地址
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
