package getaccess

import (
	mock "gomock-sample/mock_httpaccess"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetAccessFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testMock := mock.NewMockTestAPI(ctrl)
	var responceJSON message
	testMock.EXPECT().GET("/auth/test", "authorizationkey", &responceJSON).Return(http.StatusOK, nil)

	test := NewGetAccess(testMock)
	statuscode, err := test.GetAccessFunc()
	if err != nil {
		t.Error("failed Test")
	}
	if statuscode != http.StatusOK {
		t.Error("failed Test")
	}
}
