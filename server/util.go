package server

type Request struct {
	ApplicationName string `json:"applicationName"`
	WorkerIp        string `json:"workerIp"`
}

type Reply struct {
	Ok bool `json:"ok"`
}
