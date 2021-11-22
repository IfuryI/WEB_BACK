package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/IfuryI/WEB_BACK/internal/actors"
	actorsHttp "github.com/IfuryI/WEB_BACK/internal/actors/delivery/http"
	actorsDBStorage "github.com/IfuryI/WEB_BACK/internal/actors/repository/dbstorage"
	actorsUseCase "github.com/IfuryI/WEB_BACK/internal/actors/usecase"
	"github.com/IfuryI/WEB_BACK/internal/logger"
	"github.com/IfuryI/WEB_BACK/internal/middleware"
	"github.com/IfuryI/WEB_BACK/internal/movies"
	moviesHttp "github.com/IfuryI/WEB_BACK/internal/movies/delivery/http"
	moviesDBStorage "github.com/IfuryI/WEB_BACK/internal/movies/repository/dbstorage"
	moviesUseCase "github.com/IfuryI/WEB_BACK/internal/movies/usecase"
	"github.com/IfuryI/WEB_BACK/internal/playlists"
	playlistsHttp "github.com/IfuryI/WEB_BACK/internal/playlists/delivery"
	playlistsRepository "github.com/IfuryI/WEB_BACK/internal/playlists/repository"
	playlistsUseCase "github.com/IfuryI/WEB_BACK/internal/playlists/usecase"
	"github.com/IfuryI/WEB_BACK/internal/proto"
	"github.com/IfuryI/WEB_BACK/internal/ratings"
	ratingsHttp "github.com/IfuryI/WEB_BACK/internal/ratings/delivery"
	ratingsDBStorage "github.com/IfuryI/WEB_BACK/internal/ratings/repository/dbstorage"
	ratingsUseCase "github.com/IfuryI/WEB_BACK/internal/ratings/usecase"
	"github.com/IfuryI/WEB_BACK/internal/reviews"
	reviewsHttp "github.com/IfuryI/WEB_BACK/internal/reviews/delivery/http"
	reviewsDBStorage "github.com/IfuryI/WEB_BACK/internal/reviews/repository/dbstorage"
	reviewsUseCase "github.com/IfuryI/WEB_BACK/internal/reviews/usecase"
	"github.com/IfuryI/WEB_BACK/internal/search"
	searchHttp "github.com/IfuryI/WEB_BACK/internal/search/delivery/http"
	searchUseCase "github.com/IfuryI/WEB_BACK/internal/search/usecase"
	sessionsDelivery "github.com/IfuryI/WEB_BACK/internal/sessions/delivery"
	"github.com/IfuryI/WEB_BACK/internal/users"
	usersHttp "github.com/IfuryI/WEB_BACK/internal/users/delivery/http"
	usersDBStorage "github.com/IfuryI/WEB_BACK/internal/users/repository/dbstorage"
	usersUseCase "github.com/IfuryI/WEB_BACK/internal/users/usecase"
	constants "github.com/IfuryI/WEB_BACK/pkg/const"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

// App структура главного приложения
type App struct {
	server         *http.Server
	usersUC        users.UseCase
	actorsUC       actors.UseCase
	moviesUC       movies.UseCase
	ratingsUC      ratings.UseCase
	reviewsUC      reviews.UseCase
	playlistsUC    playlists.UseCase
	searchUC       search.UseCase
	authMiddleware middleware.Auth
	csrfMiddleware middleware.Csrf
	logger         *logger.Logger
	sessionsDL     *sessionsDelivery.AuthClient
	sessionsConn   *grpc.ClientConn
	fileServer     proto.FileServerHandlerClient
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// NewApp инициализация приложения
func NewApp() *App {
	accessLogger := logger.NewAccessLogger()

	connStr, connected := os.LookupEnv("DB_CONNECT")
	if !connected {
		fmt.Println(os.Getwd())
		log.Fatal("Failed to read DB connection data")
	}
	dbpool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	sessionsGrpcConn, err := grpc.Dial(fmt.Sprintf("localhost:%s", constants.AuthPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to grpc auth server: %v\n", err)
	}
	sessionsDL := sessionsDelivery.NewAuthClient(sessionsGrpcConn)

	fileServerGrpcConn, err := grpc.Dial(fmt.Sprintf("localhost:%s", constants.FileServerPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to grpc file server: %v\n", err)
	}
	fileServerService := proto.NewFileServerHandlerClient(fileServerGrpcConn)

	usersRepo := usersDBStorage.NewUserRepository(dbpool)
	reviewsRepo := reviewsDBStorage.NewReviewRepository(dbpool)
	moviesRepo := moviesDBStorage.NewMovieRepository(dbpool)
	actorsRepo := actorsDBStorage.NewActorRepository(dbpool)
	ratingsRepo := ratingsDBStorage.NewRatingsRepository(dbpool)

	usersUC := usersUseCase.NewUsersUseCase(usersRepo, reviewsRepo, ratingsRepo, actorsRepo)
	moviesUC := moviesUseCase.NewMoviesUseCase(moviesRepo, usersRepo)
	actorsUC := actorsUseCase.NewActorsUseCase(actorsRepo)
	reviewsUC := reviewsUseCase.NewReviewsUseCase(reviewsRepo, usersRepo)
	ratingsUC := ratingsUseCase.NewRatingsUseCase(ratingsRepo)
	searchUC := searchUseCase.NewSearchUseCase(usersRepo, moviesRepo, actorsRepo)

	playlistsRepo := playlistsRepository.NewPlaylistsRepository(dbpool)
	playlistsUC := playlistsUseCase.NewPlaylistsUseCase(playlistsRepo)

	authMiddleware := middleware.NewAuthMiddleware(usersUC, sessionsDL)
	csrfMiddleware := middleware.NewCsrfMiddleware(accessLogger)

	return &App{
		usersUC:        usersUC,
		actorsUC:       actorsUC,
		moviesUC:       moviesUC,
		ratingsUC:      ratingsUC,
		reviewsUC:      reviewsUC,
		playlistsUC:    playlistsUC,
		searchUC:       searchUC,
		authMiddleware: authMiddleware,
		csrfMiddleware: csrfMiddleware,
		logger:         accessLogger,
		sessionsDL:     sessionsDL,
		sessionsConn:   sessionsGrpcConn,
		fileServer:     fileServerService,
	}
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Run запуск приложения
func (app *App) Run(port string) error {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4000", "https://cinemedia.ru"}
	config.AllowCredentials = true
	router.Use(cors.New(config))
	router.Use(middleware.AccessLogMiddleware(app.logger))

	router.Static("/avatars", constants.AvatarsFileDir)
	router.Static("/posters", constants.PostersFileDir)
	router.Static("/banners", constants.BannersFileDir)
	router.Static("/actors", constants.ActorsFileDir)

	router.Use(gin.Recovery())
	router.GET("/metrics", prometheusHandler())

	api := router.Group("/api")
	v1 := api.Group("/v1")

	usersHttp.RegisterHTTPEndpoints(v1, app.usersUC, app.sessionsDL, app.authMiddleware, app.fileServer, app.logger)
	moviesHttp.RegisterHTTPEndpoints(v1, app.moviesUC, app.authMiddleware, app.logger)
	ratingsHttp.RegisterHTTPEndpoints(v1, app.ratingsUC, app.authMiddleware, app.logger)
	reviewsHttp.RegisterHTTPEndpoints(v1, app.reviewsUC, app.usersUC, app.authMiddleware, app.logger)
	actorsHttp.RegisterHTTPEndpoints(v1, app.actorsUC, app.authMiddleware, app.logger)
	playlistsHttp.RegisterHTTPEndpoints(v1, app.playlistsUC, app.usersUC, app.authMiddleware, app.logger)
	searchHttp.RegisterHTTPEndpoints(v1, app.searchUC, app.logger)

	app.server = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := app.server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to listen and serve: ", err)
		}
	}()

	// using graceful shutdown

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	_ = app.sessionsConn.Close()
	return app.server.Shutdown(ctx)
}
