package thirdparty

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/iqbaldp78/form3-api/apps/models"
	"github.com/spf13/viper"
)

type form3 struct {
	resty *resty.Client
	host  string
}

func NewInitForm3(resty *resty.Client) *form3 {
	return &form3{
		resty: resty,
		host:  viper.GetString("client.form3_url"),
	}
}

func (f *form3) RequestFetchData() (result models.RespGetAccount, err error) {
	pathUrl := fmt.Sprintf("%v/v1/organisation/accounts/%v", f.host, "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	resp, err := f.resty.SetDebug(true).R().
		SetHeader("Accept", "application/json").
		SetResult(&result).
		Get(pathUrl)

	if err != nil {
		return result, err
	}
	if resp.IsError() {
		return result, fmt.Errorf("unable to request form3 api ")
	}
	return result, nil
}
