package casbinfs

const (
	defaultPathBase      = "/policies"
	defaultPathModel     = "/model.conf"
	defaultPathPolicy    = "/policy.csv"
	defaultPathModelAbs  = defaultPathBase + defaultPathModel
	defaultPathPolicyAbs = defaultPathBase + defaultPathPolicy
)

func normalizePaths(paths []string) []string {
	if len(paths) == 0 {
		return []string{
			defaultPathModelAbs,
			defaultPathPolicyAbs,
		}
	}
	if len(paths) == 1 {
		return []string{
			paths[0] + defaultPathModel,
			paths[0] + defaultPathPolicy,
		}
	}
	return []string{paths[0], paths[1]}
}
