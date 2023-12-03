package main

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/storage/sqlite3"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sourcehaven/mypass-api/internal/db"
	"github.com/sourcehaven/mypass-api/internal/glob"
	"github.com/sourcehaven/mypass-api/internal/model"
	"github.com/sourcehaven/mypass-api/internal/routes"
	"github.com/sourcehaven/mypass-api/internal/util"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
)

func init() {
	initLogger := func() {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	initLogger()
}

func getEnvironment() string {
	env := os.Getenv(glob.MypassEnv)

	if env == "" {
		log.Panic().Msg("Environment variable '" + glob.MypassEnv + "' must be defined!")
	}

	return util.ToEnv(env)
}

func getConfig(env string) *model.Cfg {

	loadEnvFile := func(env string) {
		err := godotenv.Load(".env." + env)
		if err != nil {
			log.Panic().Err(err).Msg("You might forgot to create the '.env." + env + "' file or it is incorrectly configured!")
		}
	}

	makeConfig := func(env string) *model.Cfg {
		getEnv := func(key string) string {
			val := os.Getenv(key)
			if val == "" {
				log.Panic().Msg("Environment variable: " + key + " must be set!")
			}
			return val
		}

		lvl, _ := strconv.ParseInt(getEnv(glob.MypassLoglevel), 10, 8)

		cfg := &model.Cfg{
			Host:               getEnv(glob.MypassHost),
			Port:               getEnv(glob.MypassPort),
			SecretKey:          getEnv(glob.MypassSecretKey),
			Environment:        model.Env(env),
			LogLevel:           zerolog.Level(lvl),
			DbConnectionUri:    getEnv(glob.MypassDbConnectionUri),
			PasswordLength:     util.ToUint64(getEnv(glob.MypassPasswordLength)),
			PasswordMinNumber:  util.ToUint64(getEnv(glob.MypassPasswordMinNumber)),
			PasswordMinCapital: util.ToUint64(getEnv(glob.MypassPasswordMinCapital)),
			PasswordMinSpecial: util.ToUint64(getEnv(glob.MypassPasswordMinSpecial)),
		}
		return cfg
	}

	loadEnvFile(env)
	return makeConfig(env)
}

func runServer(addr string, app *fiber.App) {
	if err := app.Listen(addr); err != nil {
		log.Error().Err(err).Msg("Failed to start server")
	}
}

func setupRoutes(ds db.TransactDatastore, app *fiber.App, cfg *model.Cfg) {
	// Setup API routes
	api := app.Group("/api")
	api.Get("/teapot", routes.Teapot)

	handler := routes.Handler{Store: ds, Cfg: cfg}

	user := api.Group("/user")
	user.Post("register", handler.RegisterUser)
	user.Post("login", handler.LoginUser)
}

func createTransactDb(connStr string) (db.TransactDatastore, error) {

	// dbLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Cfg{
	// 		SlowThreshold: time.Second, // Slow SQL threshold
	// 		LogLevel:      logger.Info, // Log level
	// 		Colorful:      true,        // Disable color
	// 	},
	// )

	gormDb, err := gorm.Open(sqlite.Open(connStr)) //, &gorm.Cfg{Logger: dbLogger}
	if err != nil {
		return nil, err
	}

	tranDb := &db.TransactDb{DB: gormDb}
	err = tranDb.Init()
	if err != nil {
		return nil, err
	}

	return tranDb, nil
}

func createFiberApp(jwtSigningKey string) *fiber.App {
	fiberApp := fiber.New()

	// JWT Middleware
	fiberApp.Use("/api", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSigningKey)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if c.Path() == "/api/user/register" || c.Path() == "/api/teapot" || c.Path() == "/api/user/login" {
				return nil
			}
			// If the token is missing or not valid, send an unauthorized status and message
			return c.Status(http.StatusUnauthorized).JSON(
				routes.NewResponse(
					http.StatusText(http.StatusUnauthorized),
					"Missing or invalid JWT",
					nil,
				))
		},
	}))

	storage := sqlite3.New()
	fiberApp.Use(limiter.New(limiter.Config{
		Max:               10,
		Storage:           storage,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	return fiberApp
}

func setupSwagger(app *fiber.App) {
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}
	app.Use(swagger.New(cfg))

	// Redirect from the root URL to /swagger
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger")
	})
}

func main() {
	env := getEnvironment()
	cfg := getConfig(env)

	tranDb, _ := createTransactDb(cfg.DbConnectionUri)
	app := createFiberApp(cfg.SecretKey)

	setupRoutes(tranDb, app, cfg)
	setupSwagger(app)

	runServer(cfg.Host+":"+cfg.Port, app)
}
