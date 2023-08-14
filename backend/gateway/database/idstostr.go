package database

type ID interface {
	~string
}

func IDsToString[I ID](ids []I) []string {
	ret := make([]string, 0, len(ids))
	for _, v := range ids {
		ret = append(ret, string(v))
	}
	return ret
}
