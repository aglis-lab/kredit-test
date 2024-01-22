package app

import (
	"context"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

type appContext struct {
	config           *Configuration
	gorm             *gorm.DB
	validator        *validator.Validate
	tracer           trace.Tracer
	kafkaTransaction *kafka.Conn
}

var appCtx appContext

func Init(ctx context.Context) error {
	// Init Config
	config, err := InitConfig(ctx)
	if err != nil {
		return err
	}

	// Init Kafka
	kafkaTransaction, err := kafka.DialLeader(context.Background(), "tcp", config.Kafka.Host, config.Kafka.TransactionTopic, 0)
	if err != nil {
		return err
	}

	// Init Gorm Database
	// TODO: I don't like show query inside terminal
	// logger := logger.New(
	// 	logrus.NewWriter(),
	// 	logger.Config{
	// 		SlowThreshold: time.Millisecond,
	// 		LogLevel:      logger.Warn,
	// 		Colorful:      false,
	// 	},
	// )
	// gormConfig := gorm.Config{
	// 	Logger:                                   logger,
	// 	DisableForeignKeyConstraintWhenMigrating: true,
	// }

	gormConfig := gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	dbConfig := config.MySQL
	gorm, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbConfig.Username, dbConfig.Password, dbConfig.ConnURI, dbConfig.Database)), &gormConfig)
	if err != nil {
		log.Panic(err)
	}

	// Set Gorm Tracing
	if err := gorm.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		log.Panic(err)
	}

	appCtx = appContext{
		gorm:             gorm,
		config:           config,
		tracer:           otel.Tracer(config.ServiceName, trace.WithInstrumentationVersion(config.ServiceVersion)),
		validator:        validator.New(),
		kafkaTransaction: kafkaTransaction,
	}

	return nil
}

func GormDB() *gorm.DB {
	return appCtx.gorm
}

func Config() *Configuration {
	return appCtx.config
}

func Tracer() trace.Tracer {
	return appCtx.tracer
}

func Validator() *validator.Validate {
	return appCtx.validator
}

func KafkaTransaction() *kafka.Conn {
	return appCtx.kafkaTransaction
}
