package main

import (
	"ewallet-ums/infra"
	"ewallet-ums/internal/grpc"
	"ewallet-ums/internal/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	// load config
	infra.InitConfig()

	// load log
	infra.InitLogger()

	// load db
	postgresDB := infra.InitPostgresql()

	// load redis
	redisRepo := infra.InitRedis()

	// load jwt
	jwtManager := infra.InitJWT()

	appDeps := &infra.AppDependencies{
		PostgresDB: postgresDB,
		RedisRepo:  redisRepo,
		JWTManager: jwtManager,
	}

	// Inject dependency
	dependencies := infra.DependencyInject(appDeps)

	// run grpc
	go grpc.ServeGRPC(dependencies)

	// run http
	http.Serve(dependencies, appDeps)
}
