package config

import (
	"database/sql"
	"tradesignal-backend/pkg/aes"
	"tradesignal-backend/pkg/aesfront"
	"tradesignal-backend/pkg/aesmansek"
	"tradesignal-backend/pkg/jwe"
	"tradesignal-backend/pkg/jwt"
	"tradesignal-backend/pkg/mail"
	"tradesignal-backend/pkg/mailing"
	"tradesignal-backend/pkg/mandrill"
	"tradesignal-backend/pkg/mansekum"

	// miniopkg "tradesignal-backend/pkg/minio"
	"log"
	"time"
	"tradesignal-backend/pkg/mssqldb"
	postgresqlPkg "tradesignal-backend/pkg/postgresql"
	"tradesignal-backend/pkg/recaptcha"
	redisPkg "tradesignal-backend/pkg/redis"
	"tradesignal-backend/pkg/str"

	"cloud.google.com/go/firestore"
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
)

// Configs ...
type Configs struct {
	EnvConfig   map[string]string
	DB          *sql.DB
	DBMS        *sql.DB
	RedisClient redisPkg.RedisClient
	JweCred     jwe.Credential
	JwtCred     jwt.Credential
	Aes         aes.Credential
	AesMansek   aesmansek.Credential
	AesFront    aesfront.Credential
	Minio       *minio.Client
	Firestore   *firestore.Client
	Mandrill    mandrill.Credential
	Recaptcha   recaptcha.Credential
	Mail        mail.Connection
	Mailing     mailing.GoMailConfig
	Mansekum    mansekum.Credential
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

	// SQL server OAO2 DB connection
	dbmsInfo := mssqldb.Connection{
		Server:   res.EnvConfig["MSSQL_HOST"],
		Port:     str.StringToInt(res.EnvConfig["MSSQL_PORT"]),
		User:     res.EnvConfig["MSSQL_USERNAME"],
		Password: res.EnvConfig["MSSQL_PASSWORD"],
		DB:       res.EnvConfig["MSSQL_DB"],
		Debug:    str.StringToBool(res.EnvConfig["MSSQL_DEBUG"]),
	}
	res.DBMS, err = dbmsInfo.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

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

	// AES Mansek credential
	res.AesMansek = aesmansek.Credential{
		Key: res.EnvConfig["AES_MANSEK_KEY"],
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

	res.Mandrill = mandrill.Credential{
		Key:      res.EnvConfig["MANDRILL_KEY"],
		FromMail: res.EnvConfig["MANDRILL_FROM_MAIL"],
		FromName: res.EnvConfig["MANDRILL_FROM_NAME"],
	}

	res.Recaptcha = recaptcha.Credential{
		Secret: res.EnvConfig["RECAPTCHA_SECRET"],
	}

	res.Mail = mail.Connection{
		Host:     res.EnvConfig["SMTP_HOST"],
		Port:     str.StringToInt(res.EnvConfig["SMTP_PORT"]),
		Username: res.EnvConfig["SMTP_USERNAME"],
		Password: res.EnvConfig["SMTP_PASSWORD"],
	}

	res.Mailing = mailing.GoMailConfig{
		SMTPHost: res.EnvConfig["SMTP_HOST"],
		SMTPPort: str.StringToInt(res.EnvConfig["SMTP_PORT"]),
		Sender:   res.EnvConfig["SMTP_FROM"],
		Username: res.EnvConfig["SMTP_USERNAME"],
		Password: res.EnvConfig["SMTP_PASSWORD"],
	}

	res.Mansekum = mansekum.Credential{
		BaseURL:  res.EnvConfig["MANSEK_UM_BASE_URL"],
		Username: res.EnvConfig["MANSEK_UM_USERNAME"],
		Password: res.EnvConfig["MANSEK_UM_PASSWORD"],
		AppCode:  res.EnvConfig["MANSEK_UM_APP_CODE"],
		AdminID:  res.EnvConfig["MANSEK_UM_ADMIN_ID"],
		AdminKey: res.EnvConfig["MANSEK_UM_ADMIN_KEY"],
		AppID:    res.EnvConfig["MANSEK_UM_APP_ID"],
	}

	return res, err
}
