package bootstrap

import (
	"PawInHand/config"
	"PawInHand/generated/dao"
	"PawInHand/repositories"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"strconv"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/fx"
	gorm_postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var FXModule_Persistence = fx.Module(
	"persistence",
	fx.Provide(
		createEmbeddedPostgres,
		createEntityManager,
		createDatabaseMigrator,
		createDatabaseDriver,
		createDatabaseConnection,
		createEntityManagerConnection,
		createAdRepo,
		createPostRepo,
		createShelterRepo,
		createUserRepo,
		createVolunteerWorkRepo,
		createRatingRepo,
		createEventRepo,
	),

	fx.Invoke(
		registerEmbeddedPostgresStopHook,
		performDatabaseSchemaMigration,
	),
)

// CreateAdRepo initializes and returns the AdRepo.
func createAdRepo(q *dao.Query) repositories.AdRepo {
	return repositories.NewAdRepo(q)
}

// CreatePostRepo initializes and returns the PostRepo.
func createPostRepo(q *dao.Query) repositories.PostRepo {
	return repositories.NewPostRepo(q)
}

// CreateShelterRepo initializes and returns the ShelterRepo.
func createShelterRepo(q *dao.Query) repositories.ShelterRepo {
	return repositories.NewShelterRepo(q)
}

// CreateVolunteerWorkRepo initializes and returns the VolunteerWorkRepo.
func createVolunteerWorkRepo(q *dao.Query) repositories.VolunteerworkRepo {
	return repositories.NewVolunteerworkRepo(q)
}

// CreateRatingRepo initializes and returns the RatingRepo.
func createRatingRepo(q *dao.Query) repositories.RatingRepo {
	return repositories.NewRatingRepo(q)
}

// CreateEventRepo initializes and returns the EventRepo.
func createEventRepo(q *dao.Query) repositories.EventRepo {
	return repositories.NewEventRepo(q)
}

// CreateUserRepo initializes and returns the UserRepo.
func createUserRepo(q *dao.Query) repositories.UserRepo {
	return repositories.NewUserRepo(q)
}

func registerEmbeddedPostgresStopHook(lc fx.Lifecycle, embeddedDB *embeddedpostgres.EmbeddedPostgres) {
	lc.Append(fx.StopHook(func() error {

		return embeddedDB.Stop()
	}))
}

func createEntityManagerConnection(db *sql.DB) (*gorm.DB, error) {
	return gorm.Open(
		gorm_postgres.New(gorm_postgres.Config{Conn: db}),
		&gorm.Config{
			TranslateError:         true,
			SkipDefaultTransaction: false,
		},
	)
}

func createEmbeddedPostgres(configuration *config.ApplicationConfiguration) (*embeddedpostgres.EmbeddedPostgres, error) {
	port, err := strconv.Atoi(configuration.DBPort)
	if err != nil {
		return nil, err
	}

	db := embeddedpostgres.NewDatabase(
		embeddedpostgres.DefaultConfig().
			Database(configuration.DBName).
			Port(uint32(port)).
			Username(configuration.DBUsername).
			Username(configuration.DBPassword),
	)

	return db, db.Start()
}

func createEntityManager(db *gorm.DB) *dao.Query {
	return dao.Use(db)
}

func createDatabaseMigrator(config *config.ApplicationConfiguration, driver database.Driver) (*migrate.Migrate, error) {
	return migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%v", config.DatabaseMigrationsPath), "postgres", driver)
}

func performDatabaseSchemaMigration(migrator *migrate.Migrate) error {
	_, dirty, _ := migrator.Version()
	if dirty {
		_ = migrator.Drop()

		return fmt.Errorf("failed performing migrations, dirty DB state detected")
	}

	err := migrator.Up()

	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}

	return err
}

func createDatabaseDriver(db *sql.DB) (database.Driver, error) {
	return postgres.WithInstance(db, &postgres.Config{})
}

func createDatabaseConnection(config *config.ApplicationConfiguration) (*sql.DB, error) {
	db, err := sql.Open("postgres", buildDatabaseURL(config))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func buildDatabaseURL(config *config.ApplicationConfiguration) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&binary_parameters=%s",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
		config.SSLMode,
		config.BinaryParameter,
	)
}
