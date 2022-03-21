package bigBrother

type GetAppsArgs struct {
}

type GetAppsReply struct {
	Applications []string
}

type DeleteAppArgs struct {
	Application string
}

type DeleteAppReply struct {
	Success bool
}
