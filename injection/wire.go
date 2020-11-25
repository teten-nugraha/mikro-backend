package injection

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/teten-nugraha/mikro-backend/api"
	"github.com/teten-nugraha/mikro-backend/respository"
	"github.com/teten-nugraha/mikro-backend/service"
)

func initAuthAPI(db *gorm.DB) api.AuthAPI {
	wire.Build(respository.ProviderUserRepository, service.ProviderUserService, api.ProvideAuthAPI)

	return api.AuthAPI{}
}

func initUserAPI(db *gorm.DB) api.UserAPI {
	wire.Build(respository.ProviderUserRepository, service.ProviderUserService, api.ProviderUserAPI)

	return api.UserAPI{}
}