package controller

var ()

func GetHeader(token string) map[string]string {
	return map[string]string{
		"Authorization": "token " + token,
	}
}

func GetRelease(path string) string {
	return "https://api.github.com/repos" + path + "/releases"
}
