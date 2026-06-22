package infra

import (
	"ewallet-ums/external/wallet"
	"ewallet-ums/internal/domain/auth"
	"ewallet-ums/internal/domain/user"
	"ewallet-ums/internal/handler"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/service"
)

type Dependency struct {
	UserRepo         user.IRepository
	RegisterAPI      auth.IRegisterHandler
	LoginAPI         auth.ILoginHandler
	LogoutAPI        auth.ILogoutHandler
	RefreshTokenAPI  auth.IRefreshTokenHandler
	TokenValidateAPI auth.ITokenValidationHandler
}

func DependencyInject(appDeps *AppDependencies) *Dependency {

	extWallet := &wallet.ExtWallet{}

	userRepo := &repository.UserRepository{
		DB: appDeps.PostgresDB,
	}

	registerSvc := &service.RegisterService{
		UserRepo:       userRepo,
		ExternalWallet: extWallet,
	}

	loginSvc := &service.LoginService{
		UserRepo:   userRepo,
		JwtManager: appDeps.JWTManager,
		RedisRepo:  appDeps.RedisRepo,
	}

	logoutSvc := &service.LogoutService{
		UserRepo:  userRepo,
		RedisRepo: appDeps.RedisRepo,
	}

	refreshSvc := &service.RefrshTokenService{
		UserRepo:   userRepo,
		JwtManager: appDeps.JWTManager,
		RedisRepo:  appDeps.RedisRepo,
	}

	tokenValidateSvc := &service.TokenValidationService{
		UserRepo:   userRepo,
		JwtManager: appDeps.JWTManager,
		RedisRepo:  appDeps.RedisRepo,
	}

	registerAPI := &handler.RegisterHandler{
		RegisterSvc: registerSvc,
	}
	loginAPI := &handler.LoginHandler{
		LoginSvc: loginSvc,
	}
	logoutAPI := &handler.LogoutHandler{
		LogoutSvc: logoutSvc,
	}
	refreshTokenAPI := &handler.RefreshTokenHandler{
		RefreshTokenSvc: refreshSvc,
	}
	tokenValidateAPI := &handler.TokenValidationHandler{
		TokenValidationService: tokenValidateSvc,
	}

	return &Dependency{
		UserRepo:         userRepo,
		RegisterAPI:      registerAPI,
		LoginAPI:         loginAPI,
		LogoutAPI:        logoutAPI,
		RefreshTokenAPI:  refreshTokenAPI,
		TokenValidateAPI: tokenValidateAPI,
	}
}
