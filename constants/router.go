package constants

const (
	IdParam       string = "id"
	ApiV1Route    string = "/api/v1"
	UserRoute     string = "/user"
	IdRoute       string = "/:" + IdParam
	UserByIdRoute string = UserRoute + IdRoute
)
