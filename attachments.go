// generic attachements

package libgomessage

import(

)




type ApiKey struct {
	ApiKey string
}

func (api *ApiKey) GetApiKey() string {
	return api.ApiKey
}
