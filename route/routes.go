package route

import (
	"phone_review/constants"
	"phone_review/controller"
	"phone_review/middleware"
	m "phone_review/middleware"

	"github.com/labstack/echo/v4"
	midd "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	//CRUD ADMIN
	e.GET("/admins", controller.GetAdminsController)
	e.GET("/admins/:id", controller.GetAdminsByID)
	m.LogMiddleware(e)
	e.POST("/admins", controller.CreateAdminController)
	e.DELETE("/admins/:id", controller.DeleteAdminByID)
	e.PUT("/admins/:id", controller.UpdateAdminByID)
	e.POST("/login", controller.LoginAdminController)

	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(midd.BasicAuth(middleware.BasicAuthDB))
	eAuthBasic.GET("/admins", controller.GetAdminsController)

	eJwt := e.Group("/jwt")
	eJwt.Use(midd.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/admins", controller.GetAdminsController)

	//CRUD MERK
	e.GET("/merks", controller.GetMerks)
	e.GET("/merks/:id", controller.GetMerksByID)
	e.POST("/merks", controller.CreateMerk)
	e.DELETE("/merks/:id", controller.DeleteMerkByID)
	e.PUT("/merks/:id", controller.UpdateMerkByID)

	//CRUD Type Smartphone
	e.GET("/types", controller.GetTypes)
	e.GET("/types/:id", controller.GetTypesByID)
	e.POST("/types", controller.CreateType)
	e.DELETE("/types/:id", controller.DeleteTypeByID)
	e.PUT("/types/:id", controller.UpdateTypeByID)

	//CRUD Advances Smartphone
	e.GET("/advances", controller.GetAdvances)
	e.GET("/advances/:id", controller.GetAdvancesByID)
	e.POST("/advances", controller.CreateAdvances)
	e.DELETE("/advances/:id", controller.DeleteAdvanceByID)
	e.PUT("/advances/:id", controller.UpdateAdvanceByID)

	//CRUD Audios Smartphone
	e.GET("/audios", controller.GetAudios)
	e.GET("/audios/:id", controller.GetAudiosByID)
	e.POST("/audios", controller.CreateAudio)
	e.DELETE("/audios/:id", controller.DeleteAudioByID)
	e.PUT("/audios/:id", controller.UpdateAudioByID)

	//CRUD Batteries Smartphone
	e.GET("/batteries", controller.GetBatteries)
	e.GET("/batteries/:id", controller.GetBatteriesByID)
	e.POST("/batteries", controller.CreateBattery)
	e.DELETE("/batteries/:id", controller.DeleteBatteryByID)
	e.PUT("/batteries/:id", controller.UpdateBatteryByID)

	//CRUD Desains Smartphone
	e.GET("/desains", controller.GetDesains)
	e.GET("/desains/:id", controller.GetDesainsByID)
	e.POST("/desains", controller.CreateDesains)
	e.DELETE("/desains/:id", controller.DeleteDesainByID)
	e.PUT("/desains/:id", controller.UpdateDesainByID)

	//CRUD Deskripsi Smartphone
	e.GET("/descs", controller.GetDescs)
	e.GET("/descs/:id", controller.GetDescsByID)
	e.POST("/descs", controller.CreateDescs)
	e.DELETE("/descs/:id", controller.DeleteDescsByID)
	e.PUT("/descs/:id", controller.UpdateDescsByID)

	//CRUD Kamera Smartphone
	e.GET("/cameras", controller.GetCameras)
	e.GET("/cameras/:id", controller.GetCamerasByID)
	e.POST("/cameras", controller.CreateCameras)
	e.DELETE("/cameras/:id", controller.DeleteCamerasByID)
	e.PUT("/cameras/:id", controller.UpdateCamerasByID)

	//CRUD Tampilan Smartphone
	e.GET("/displays", controller.GetDisplays)
	e.GET("/displays/:id", controller.GetDisplaysByID)
	e.POST("/displays", controller.Createdisplays)
	e.DELETE("/displays/:id", controller.DeleteDisplayByID)
	e.PUT("/displays/:id", controller.UpdateDisplaysByID)

	//CRUD Kinerja Smartphone
	e.GET("/performs", controller.GetPerfoms)
	e.GET("/performs/:id", controller.GetPerfomsByID)
	e.POST("/performs", controller.CreatePerfoms)
	e.DELETE("/performs/:id", controller.DeletePerfomsByID)
	e.PUT("/performs/:id", controller.UpdatePerfomsByID)

	//CRUD Features Smartphone
	e.GET("/features", controller.GetFeatures)
	e.GET("/features/:id", controller.GetFeaturesByID)
	e.POST("/features", controller.CreateFeature)
	e.DELETE("/features/:id", controller.DeleteFeatureByID)
	e.PUT("/features/:id", controller.UpdateFeatureByID)

	return e
}
