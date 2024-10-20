package controllers

import (
	"project-skbackend/configs"
	"project-skbackend/internal/controllers/requests"
	"project-skbackend/internal/middlewares"
	"project-skbackend/internal/services/baseroleservice"
	"project-skbackend/internal/services/fileservice"
	"project-skbackend/internal/services/userservice"
	"project-skbackend/packages/consttypes"
	"project-skbackend/packages/utils/utresponse"
	"project-skbackend/packages/utils/uttoken"

	"github.com/gin-gonic/gin"
)

type (
	profileroutes struct {
		cfg   *configs.Config
		suser userservice.IUserService
		sfile fileservice.IFileService
		sbase baseroleservice.IBaseRoleService
	}
)

func newProfileRoutes(
	rg *gin.RouterGroup,
	cfg *configs.Config,
	suser userservice.IUserService,
	sfile fileservice.IFileService,
	sbase baseroleservice.IBaseRoleService,
) {
	r := &profileroutes{
		cfg:   cfg,
		suser: suser,
		sfile: sfile,
		sbase: sbase,
	}

	gprofilepvt := rg.Group("profiles")
	gprofilepvt.Use(middlewares.JWTAuthMiddleware(
		cfg,
		consttypes.UR_ADMIN,
	))
	{
		// * global route
		gprofilepvt.GET("me", r.getOwnProfile)

		// * picture's route
		gpicture := gprofilepvt.Group("pictures")
		{
			gpicture.PATCH("", r.updateProfilePicture)
		}

		gpassword := gprofilepvt.Group("passwords")
		{
			gpassword.PATCH("own", r.updateOwnPassword)
		}
	}
}

func (r *profileroutes) getOwnProfile(ctx *gin.Context) {
	var (
		function = "get own profile"
	)

	userres, err := uttoken.GetUser(ctx)
	if err != nil {
		utresponse.GeneralUnauthorized(
			ctx,
			err,
		)
		return
	}

	roleres, err := r.suser.GetRoleDataByUserID(userres.ID)
	if err != nil {
		utresponse.GeneralUnauthorized(
			ctx,
			err,
		)
		return
	}

	if roleres == nil {
		utresponse.GeneralInternalServerError(
			function,
			ctx,
			consttypes.ErrUserInvalidRole,
		)
		return
	}

	utresponse.GeneralSuccessFetch(
		function,
		ctx,
		roleres.Data,
	)
}

func (r *profileroutes) updateProfilePicture(ctx *gin.Context) {
	var (
		function = "update profile picture"
		entity   = "profile picture"
		req      *requests.UpdateImage
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ve := utresponse.ValidationResponse(err)
		utresponse.GeneralInvalidRequest(
			function,
			ctx,
			ve,
			err,
		)
		return
	}

	// * get the current logged in user
	userres, err := uttoken.GetUser(ctx)
	if err != nil {
		utresponse.GeneralUnauthorized(
			ctx,
			err,
		)
		return
	}

	// * validate and upload the image
	if err := req.Validate(); err != nil {
		utresponse.GeneralInvalidRequest(
			function,
			ctx,
			nil,
			err,
		)
		return
	}

	multipart, err := req.GetMultipartFile()
	if err != nil {
		utresponse.GeneralInternalServerError(
			function,
			ctx,
			err,
		)
		return
	}

	err = r.sfile.UploadProfilePicture(userres.ID, multipart)
	if err != nil {
		utresponse.GeneralInternalServerError(
			function,
			ctx,
			err,
		)
		return
	}

	utresponse.GeneralSuccessUpdate(
		entity,
		ctx,
		nil,
	)
}

func (r *profileroutes) updateOwnPassword(ctx *gin.Context) {
	var (
		function = "update own password"
		entity   = "own password"
		req      *requests.UpdatePassword
		err      error
	)

	userres, err := uttoken.GetUser(ctx)
	if err != nil {
		utresponse.GeneralUnauthorized(
			ctx,
			err,
		)
		return
	}

	if err := ctx.ShouldBind(&req); err != nil {
		ve := utresponse.ValidationResponse(err)
		utresponse.GeneralInvalidRequest(
			function,
			ctx,
			ve,
			err,
		)
		return
	}

	err = r.suser.UpdateOwnPassword(userres.ID, *req)
	if err != nil {
		utresponse.GeneralInternalServerError(
			function,
			ctx,
			err,
		)
		return
	}

	utresponse.GeneralSuccessUpdate(
		entity,
		ctx,
		nil,
	)
}
