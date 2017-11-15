package providers


import (

)

var Config configModel

type configModel struct{

}

func (t *configModel) GetConfig()  {

}

func init() {
	Config = * &configModel{}
}
