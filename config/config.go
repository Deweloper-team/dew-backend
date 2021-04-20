package config

import (
	"database/sql"
	"dew-backend/pkg/aes"
	"dew-backend/pkg/aesfront"
	"dew-backend/pkg/jwe"
	"dew-backend/pkg/jwt"

	// miniopkg "dew-backend/pkg/minio"

	postgresqlPkg "dew-backend/pkg/postgresql"
	redisPkg "dew-backend/pkg/redis"
	"dew-backend/pkg/str"
	"log"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
)

// Configs ...
type Configs struct {
	EnvConfig   map[string]string
	DB          *sql.DB
	RedisClient redisPkg.RedisClient
	JweCred     jwe.Credential
	JwtCred     jwt.Credential
	Aes         aes.Credential
	AesFront    aesfront.Credential
}

// LoadConfigs ...
func LoadConfigs() (res Configs, err error) {
	res.EnvConfig, err = godotenv.Read("../.env")
	if err != nil {
		log.Fatal("Error loading ..env file")
	}

	//postgresql conn
	dbConn := postgresqlPkg.Connection{
		Host:                    res.EnvConfig["DATABASE_HOST"],
		DbName:                  res.EnvConfig["DATABASE_DB"],
		User:                    res.EnvConfig["DATABASE_USER"],
		Password:                res.EnvConfig["DATABASE_PASSWORD"],
		Port:                    str.StringToInt(res.EnvConfig["DATABASE_PORT"]),
		SslMode:                 res.EnvConfig["DATABASE_SSL_MODE"],
		DBMaxConnection:         str.StringToInt(res.EnvConfig["DATABASE_MAX_CONNECTION"]),
		DBMAxIdleConnection:     str.StringToInt(res.EnvConfig["DATABASE_MAX_IDLE_CONNECTION"]),
		DBMaxLifeTimeConnection: str.StringToInt(res.EnvConfig["DATABASE_MAX_LIFETIME_CONNECTION"]),
	}
	res.DB, err = dbConn.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	res.DB.SetMaxOpenConns(dbConn.DBMaxConnection)
	res.DB.SetMaxIdleConns(dbConn.DBMAxIdleConnection)
	res.DB.SetConnMaxLifetime(time.Duration(dbConn.DBMaxLifeTimeConnection) * time.Second)

	// redis conn
	redisOption := &redis.Options{
		Addr:     res.EnvConfig["REDIS_HOST"],
		Password: res.EnvConfig["REDIS_PASSWORD"],
		DB:       0,
	}
	res.RedisClient = redisPkg.RedisClient{Client: redis.NewClient(redisOption)}

	// jwe
	res.JweCred = jwe.Credential{
		KeyLocation: res.EnvConfig["APP_PRIVATE_KEY_LOCATION"],
		Passphrase:  res.EnvConfig["APP_PRIVATE_KEY_PASSPHRASE"],
	}

	// jwt
	res.JwtCred = jwt.Credential{
		Secret:           res.EnvConfig["TOKEN_SECRET"],
		ExpSecret:        str.StringToInt(res.EnvConfig["TOKEN_EXP_SECRET"]),
		RefreshSecret:    res.EnvConfig["TOKEN_REFRESH_SECRET"],
		RefreshExpSecret: str.StringToInt(res.EnvConfig["TOKEN_EXP_REFRESH_SECRET"]),
	}

	// aes
	res.Aes = aes.Credential{
		Key: res.EnvConfig["AES_KEY"],
	}

	// aes front
	res.AesFront = aesfront.Credential{
		Key: res.EnvConfig["AES_FRONT_KEY"],
		Iv:  res.EnvConfig["AES_FRONT_IV"],
	}

	// Minio connection
	// minioInfo := miniopkg.Connection{
	// 	EndSuperhub: res.EnvConfig["MINIO_ENDPOINT"],
	// 	AccessKey:   res.EnvConfig["MINIO_ACCESS_KEY_ID"],
	// 	SecretKey:   res.EnvConfig["MINIO_SECRET_ACCESS_KEY"],
	// 	UseSSL:      str.StringToBool(res.EnvConfig["MINIO_USE_SSL"]),
	// }
	// res.Minio, err = minioInfo.InitClient()
	// if err != nil {
	// 	return res, err
	// }

	// res.Mandrill = mandrill.Credential{
	// 	Key:      res.EnvConfig["MANDRILL_KEY"],
	// 	FromMail: res.EnvConfig["MANDRILL_FROM_MAIL"],
	// 	FromName: res.EnvConfig["MANDRILL_FROM_NAME"],
	// }

	return res, err
}
