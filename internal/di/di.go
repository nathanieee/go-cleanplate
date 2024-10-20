package di

import (
	"context"
	"project-skbackend/configs"
	"project-skbackend/external/services/distancematrixservice"
	"project-skbackend/internal/repositories/adminrepo"
	"project-skbackend/internal/repositories/imagerepo"
	"project-skbackend/internal/repositories/userimagerepo"
	"project-skbackend/internal/repositories/userrepo"
	"project-skbackend/internal/services/authservice"
	"project-skbackend/internal/services/baseroleservice"
	"project-skbackend/internal/services/consumerservice"
	"project-skbackend/internal/services/cronservice"
	"project-skbackend/internal/services/fileservice"
	"project-skbackend/internal/services/mailservice"
	"project-skbackend/internal/services/producerservice"
	"project-skbackend/internal/services/userservice"
	"project-skbackend/packages/utils/utlogger"

	"github.com/minio/minio-go/v7"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type DependencyInjection struct {
	// * internal services
	UserService     *userservice.UserService
	AuthService     *authservice.AuthService
	MailService     *mailservice.MailService
	ConsumerService *consumerservice.ConsumerService
	CronService     *cronservice.CronService
	FileService     *fileservice.FileService
	BaseRoleService *baseroleservice.BaseRoleService

	// * external services
	DistanceMatrixService *distancematrixservice.DistanceMatrixService
}

func NewDependencyInjection(ctx context.Context, db *gorm.DB, ch *amqp.Channel, cfg *configs.Config, rdb *redis.Client, minio *minio.Client) *DependencyInjection {
	// ! -------------------------------- database -------------------------------- ! //
	if cfg.DB.LogMode {
		db = db.Debug()
	}

	// ! ------------------------------- repository ------------------------------- ! //
	ruser := userrepo.NewUserRepository(db)
	radmin := adminrepo.NewAdminRepository(db)
	rimg := imagerepo.NewImageRepository(db)
	ruimg := userimagerepo.NewUserImageRepository(db)

	// ! --------------------------------- service -------------------------------- ! //
	// * external services
	sdsmx := distancematrixservice.NewDistanceMatrixService(cfg)

	// * internal services
	sbsrl := baseroleservice.NewBaseRoleService(radmin)
	sprod := producerservice.NewProducerService(ch, cfg, ctx)
	suser := userservice.NewUserService(ruser, radmin)
	smail := mailservice.NewMailService(cfg, ruser, sprod)
	sauth := authservice.NewAuthService(cfg, rdb, ruser, smail, suser)
	scons := consumerservice.NewConsumerService(ch, cfg, smail)
	scron := cronservice.NewCronService(cfg)
	sfile := fileservice.NewFileService(cfg, ctx, *minio, ruser, rimg, ruimg)

	return &DependencyInjection{
		// * internal services
		UserService:     suser,
		AuthService:     sauth,
		MailService:     smail,
		ConsumerService: scons,
		CronService:     scron,
		FileService:     sfile,
		BaseRoleService: sbsrl,

		// * external services
		DistanceMatrixService: sdsmx,
	}
}

func (di *DependencyInjection) InitServices() {
	var (
		err error
	)

	// * setup consumer service
	di.ConsumerService.ConsumeTask()

	// * init cron service
	_, err = di.CronService.Init()
	utlogger.Fatal(err)
}
