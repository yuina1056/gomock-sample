package getaccess

import "gomock-sample/httpaccess"

type message struct {
	message string
}

type GetAccess interface {
	GetAccessFunc() (statuscode int, err error)
}

type getaccess struct {
	Ta httpaccess.TestAPI
}

func NewGetAccess(ta httpaccess.TestAPI) GetAccess {
	return &getaccess{ta}
}

func (ga getaccess) GetAccessFunc() (statuscode int, err error) {
	var responceJSON message
	statuscode, err = ga.Ta.GET("/auth/test", "authorizationkey", &responceJSON)
	if err != nil {
		return statuscode, err
	}
	return statuscode, nil
}
