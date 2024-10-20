package cronservice

import (
	"fmt"
	"project-skbackend/configs"
	"project-skbackend/packages/utils/utlogger"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type (
	CronService struct {
		cfg *configs.Config
	}

	ICronService interface {
		Init() (gocron.Scheduler, error)
	}
)

func NewCronService(
	cfg *configs.Config,
) *CronService {
	return &CronService{
		cfg: cfg,
	}
}

func (s *CronService) Init() (gocron.Scheduler, error) {
	// * get time location from env
	tz, err := time.LoadLocation(s.cfg.API.Timezone)
	if err != nil {
		return nil, err
	}

	// * define a new scheduler instance
	gsch, err := gocron.NewScheduler(
		gocron.WithLocation(tz),
	)

	if err != nil {
		return nil, err
	}

	// * add a order job
	s.parentExampleSchedule(gsch)

	// * start the scheduler
	gsch.Start()

	return gsch, nil
}

// TODO: change this into the parent of the service that you want to schedule
func (s *CronService) parentExampleSchedule(gsch gocron.Scheduler) {
	var (
		errs []error
	)

	err := s.childExampleSchedule(gsch)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) != 0 {
		utlogger.Error(err)
	}
}

func (s *CronService) childExampleSchedule(gsch gocron.Scheduler) error {
	_, err := gsch.NewJob(
		gocron.DurationJob(
			// TODO: use the env variable
			time.Duration(1)*time.Minute,
		),
		gocron.NewTask(
			func() error {
				// * do something here regarding the about your schedule
				// * return an error if the function
				return nil
			},
		),
	)

	if err != nil {
		utlogger.Error(err)
		return err
	}

	utlogger.Info(fmt.Sprintf("Service for Cron %s Running!", "Example Schedule"))

	return nil
}
