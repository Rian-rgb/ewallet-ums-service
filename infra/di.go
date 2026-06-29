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
	RegisterHdl      auth.IRegisterHandler
	LoginHdl         auth.ILoginHandler
	LogoutHdl        auth.ILogoutHandler
	RefreshTokenHdl  auth.IRefreshTokenHandler
	TokenValidateHdl auth.ITokenValidationHandler
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
		UserRepo:   userRepo,
		JwtManager: appDeps.JWTManager,
		RedisRepo:  appDeps.RedisRepo,
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

	registerHdl := &handler.RegisterHandler{
		RegisterSvc: registerSvc,
	}
	loginHdl := &handler.LoginHandler{
		LoginSvc: loginSvc,
	}
	logoutHdl := &handler.LogoutHandler{
		LogoutSvc: logoutSvc,
	}
	refreshTokenHdl := &handler.RefreshTokenHandler{
		RefreshTokenSvc: refreshSvc,
	}
	tokenValidateHdl := &handler.TokenValidationHandler{
		TokenValidationService: tokenValidateSvc,
	}

	return &Dependency{
		UserRepo:         userRepo,
		RegisterHdl:      registerHdl,
		LoginHdl:         loginHdl,
		LogoutHdl:        logoutHdl,
		RefreshTokenHdl:  refreshTokenHdl,
		TokenValidateHdl: tokenValidateHdl,
	}
}
