package awsngfw

// Logging constants.
const (
	LogQuiet = 1 << (iota + 1)
	LogLogin
	LogGet
	LogPost
	LogPatch
	LogPut
	LogDelete
	LogAction
	LogPath
	LogSend
	LogReceive
)
