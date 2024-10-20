package configs

import (
	"errors"
	"fmt"
	"project-skbackend/internal/models"
	"project-skbackend/packages/utils/utlogger"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (db DB) GetDbConnectionUrl() string {
	url := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		db.Host, db.User, db.Password, db.Name, db.Port, db.SslMode,
	)
	return url
}

func (db DB) DBSetup(gdb *gorm.DB) error {
	if err := db.SetupExtension(gdb); err != nil {
		utlogger.Error(err)
		return err
	}

	if err := db.AutoSeedEnum(gdb); err != nil {
		utlogger.Error(err)
		return err
	}

	if err := db.AutoMigrate(gdb); err != nil {
		utlogger.Error(err)
		return err
	}

	if err := db.AutoSeedData(gdb); err != nil {
		utlogger.Error(err)
		return err
	}

	return nil
}

func (db DB) SetupExtension(gdb *gorm.DB) error {
	extfuncs := []func(*gorm.DB) error{
		InstallUUIDExtension,
	}

	var (
		errs []error
	)

	for _, extfunc := range extfuncs {
		if err := extfunc(gdb); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		utlogger.Error(errs...)
		return errors.New("error setting up extension")
	}

	return nil
}

func (db DB) AutoSeedEnum(gdb *gorm.DB) error {
	seedfuncs := []func(*gorm.DB) error{
		// ! ---------------------------------- enum ---------------------------------- ! //
		SeedGenderEnum,
		SeedImageTypeEnum,
		SeedUserRoleEnum,
	}

	var (
		errs []error
	)

	for _, seedfunc := range seedfuncs {
		if err := seedfunc(gdb); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		utlogger.Error(errs...)
		return errors.New("error seeding enum")
	}

	return nil
}

func (db DB) AutoSeedData(gdb *gorm.DB) error {
	seedfuncs := []func(*gorm.DB) error{
		// * dependent

		// ! ------------------------------- credentials ------------------------------ ! //
		SeedAdminCredentials,

		// ! ---------------------------------- data ---------------------------------- ! //
		// ! put the seeder that is independent on this part of the seeder
		// * independent

		// ! put the seeder that is dependent on other data on this part of the seeder
		// * dependent
	}

	var (
		errs []error
	)

	for _, seedfunc := range seedfuncs {
		if err := seedfunc(gdb); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		utlogger.Error(errs...)
		return errors.New("error seeding data")
	}

	return nil
}

func (db DB) GetLogLevel() logger.LogLevel {
	loglevel := logger.Warn
	if db.LogMode {
		loglevel = logger.Info
	}

	return loglevel
}

func (db DB) AutoMigrate(gdb *gorm.DB) error {
	return gdb.AutoMigrate(
		&models.User{},
		&models.UserImage{},
		&models.Address{},
		&models.AddressDetail{},

		&models.Admin{},
		&models.Image{},
	)
}
