package configs

import (
	"errors"
	"fmt"
	"os"
	"project-skbackend/internal/models"
	"project-skbackend/internal/models/base"
	"project-skbackend/packages/consttypes"
	"project-skbackend/packages/customs/ctdatatype"

	"project-skbackend/packages/utils/utlogger"
	"project-skbackend/packages/utils/utstring"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	uuidstr = "123e4567-e89b-12d3-a456-426614174000"
	id, _   = uuid.Parse(uuidstr)
)

func getGlobalHashedPassword(password string) string {
	hash, err := utstring.HashPassword(password)
	if err != nil {
		utlogger.Error(err)
		return ""
	}

	return hash
}

func checkEnumIsExist(db *gorm.DB, key string) bool {
	var (
		count int64
	)

	if err := db.Table("pg_type").Where("typname = ?", key).Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}

func createEnum(db *gorm.DB, enumname string, enumvalues ...any) error {
	if !checkEnumIsExist(db, enumname) {
		values := make([]string, len(enumvalues))
		for i, v := range enumvalues {
			values[i] = fmt.Sprintf("%v", v)
		}

		query := fmt.Sprintf("CREATE TYPE %s AS ENUM ('%s');", enumname, strings.Join(values, "','"))
		err := db.Exec(query).Error
		if err != nil {
			utlogger.Error(err)
			return err
		}
	}
	return nil
}

func InstallUUIDExtension(db *gorm.DB) error {
	// Check if the UUID generation function already exists
	checkFuncSQL := `
	SELECT EXISTS (
		SELECT 1 
		FROM pg_proc 
		WHERE proname = 'uuid_generate_v7'
	);
	`

	var exists bool
	if err := db.Raw(checkFuncSQL).Scan(&exists).Error; err != nil {
		utlogger.Error(err)
		return err
	}

	// If the function doesn't exist, create it
	if !exists {
		uuidsql := `
		CREATE OR REPLACE FUNCTION uuid_generate_v7()
		RETURNS uuid AS $$
		BEGIN
			-- Use random v4 UUID as starting point (which has the same variant we need)
			-- Then overlay timestamp
			-- Finally set version 7 by flipping the 2 and 1 bit in the version 4 string
			RETURN encode(
				set_bit(
					set_bit(
						overlay(
							uuid_send(gen_random_uuid())
							PLACING substring(int8send(floor(extract(epoch FROM clock_timestamp()) * 1000)::bigint) FROM 3)
							FROM 1 FOR 6
						),
						52, 1
					),
					53, 1
				),
				'hex'
			)::uuid;
		END
		$$ LANGUAGE plpgsql VOLATILE;
		`

		if err := db.Exec(uuidsql).Error; err != nil {
			utlogger.Error(err)
			return err
		}
	}

	// Create the UUID-OSSP extension if it doesn't exist
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		utlogger.Error(err)
		return err
	}

	return nil
}

func SeedUserRoleEnum(db *gorm.DB) error {
	// * add another role here
	return createEnum(db,
		"user_role_enum",
		consttypes.UR_ADMIN.Uint(),
		consttypes.UR_USER.Uint(),
	)
}

func SeedGenderEnum(db *gorm.DB) error {
	return createEnum(db,
		"gender_enum",
		consttypes.G_MALE.String(),
		consttypes.G_FEMALE.String(),
		consttypes.G_OTHER.String(),
	)
}

func SeedImageTypeEnum(db *gorm.DB) error {
	return createEnum(db,
		"image_type_enum",
		consttypes.IT_PROFILE.String(),
		consttypes.IT_MEAL.String(),
		consttypes.IT_MEAL_CATEGORY.String(),
		consttypes.IT_DONATION_PROOF.String(),
	)
}

func SeedAdminCredentials(db *gorm.DB) error {
	if db.Migrator().HasTable(&models.User{}) && db.Migrator().HasTable(&models.Admin{}) {
		if err := db.First(&models.Admin{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			admins := []*models.Admin{
				{
					Model: base.Model{ID: id},
					User: models.User{
						Image: &models.UserImage{
							Image: models.Image{
								Name: "Profile Image",
								Path: "https://meals-minio-api.103.127.137.58.nip.io/meals-bucket/image/jpeg-01918055-96aa-7a39-bcea-6f7fd1440165.jpg",
								Type: consttypes.IT_PROFILE,
							},
						},
						ConfirmedAt: consttypes.TimeNow(),
						Email:       os.Getenv("ADMIN_EMAIL"),
						Password:    getGlobalHashedPassword(os.Getenv("ADMIN_PASSWORD")),
						Role:        consttypes.UR_ADMIN,
						Addresses: []*models.Address{
							{
								Name:    "Indian Ocean Address",
								Address: "Indian Ocean Address, 1st floor,",
								Note:    "Ocean with the blue water",
								AddressDetail: &models.AddressDetail{
									Geolocation: models.Geolocation{
										Longitude: "-26.10305",
										Latitude:  "56.91996",
									},
									FormattedAddress: "Indian Ocean",
								},
							},
						},
					},
					FirstName:   os.Getenv("ADMIN_FIRSTNAME"),
					LastName:    os.Getenv("ADMIN_LASTNAME"),
					Gender:      consttypes.G_MALE,
					DateOfBirth: ctdatatype.CDT_DATE{Time: consttypes.TimeNow()},
				},
			}

			err := db.Session(&gorm.Session{FullSaveAssociations: true}).Create(&admins).Error
			if err != nil {
				utlogger.Error(err)
				return err
			}
		}
	}

	return nil
}
