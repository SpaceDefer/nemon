package bigBrother

type GetAppsArgs struct {
}

type GetAppsReply struct {
	Applications []ApplicationInfo
}

type DeleteAppArgs struct {
	Application ApplicationInfo
}

type DeleteAppReply struct {
	Success  bool
	ErrorMsg string
}
