package main

import (
	"layanan-kependudukan-api/article"
	"layanan-kependudukan-api/auth"
	detailPengajuan "layanan-kependudukan-api/detail_pengajuan"
	"layanan-kependudukan-api/district"
	"layanan-kependudukan-api/education"
	"layanan-kependudukan-api/handler"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/job"
	"layanan-kependudukan-api/keluarga"
	"layanan-kependudukan-api/kelurahan"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/penduduk"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/position"
	"layanan-kependudukan-api/province"
	"layanan-kependudukan-api/religion"
	"layanan-kependudukan-api/rt"
	"layanan-kependudukan-api/rw"
	"layanan-kependudukan-api/subdistrict"
	"layanan-kependudukan-api/user"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=asdf1234 dbname=layanan-kependudukan port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	authService := auth.NewService()

	userRepository := user.NewRepsitory(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	religionRepository := religion.NewRepsitory(db)
	religionService := religion.NewService(religionRepository)
	religionHandler := handler.NewReligionHandler(religionService, authService)

	educationRepository := education.NewRepsitory(db)
	educationService := education.NewService(educationRepository)
	educationHandler := handler.NewEducationHandler(educationService, authService)

	jobRepository := job.NewRepsitory(db)
	jobService := job.NewService(jobRepository)
	jobHandler := handler.NewJobHandler(jobService, authService)

	kelurahanRepository := kelurahan.NewRepsitory(db)
	kelurahanService := kelurahan.NewService(kelurahanRepository)
	kelurahanHandler := handler.NewKelurahanHandler(kelurahanService, authService)

	positionRepository := position.NewRepsitory(db)
	positionService := position.NewService(positionRepository)
	positionHandler := handler.NewPositionHandler(positionService, authService)

	articleRepository := article.NewRepsitory(db)
	articleService := article.NewService(articleRepository)
	articleHandler := handler.NewArticleHandler(articleService, authService)

	layananRepository := layanan.NewRepsitory(db)
	layananService := layanan.NewService(layananRepository)
	layananHandler := handler.NewLayananHandler(layananService, authService)

	provinceRepository := province.NewRepsitory(db)
	provinceService := province.NewService(provinceRepository)
	provinceHandler := handler.NewProvinceHandler(provinceService, authService)

	rwRepository := rw.NewRepsitory(db)
	rwService := rw.NewService(rwRepository)
	rwHandler := handler.NewRWHandler(rwService, authService)

	rtRepository := rt.NewRepsitory(db)
	rtService := rt.NewService(rtRepository)
	rtHandler := handler.NewRTHandler(rtService, authService)

	districtRepository := district.NewRepsitory(db)
	districtService := district.NewService(districtRepository)
	districtHandler := handler.NewDistrictHandler(districtService, authService)

	detailPengajuanRepository := detailPengajuan.NewRepsitory(db)
	detailPengajuanService := detailPengajuan.NewService(detailPengajuanRepository)

	pengajuanRepository := pengajuan.NewRepsitory(db)
	pengajuanService := pengajuan.NewService(pengajuanRepository)
	pengajuanHandler := handler.NewpengajuanHandler(pengajuanService, detailPengajuanService, authService)

	pendudukRepository := penduduk.NewRepsitory(db)
	pendudukService := penduduk.NewService(pendudukRepository)
	pendudukHandler := handler.NewPendudukHandler(pendudukService, authService)

	keluargaRepository := keluarga.NewRepsitory(db)
	keluargaService := keluarga.NewService(keluargaRepository)
	keluargaHandler := handler.NewKeluargaHandler(keluargaService, pendudukService, authService)

	subDistrictRepository := subdistrict.NewRepsitory(db)
	subDistrictService := subdistrict.NewService(subDistrictRepository)
	subDistrictHandler := handler.NewSubDistrictHandler(subDistrictService, authService)

	router := gin.Default()

	corsMiddleware := cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS", "CONNECT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-CSRF-Token", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	})
	router.Use(corsMiddleware)
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegiserUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checkers", authMiddleware(authService, userService), userHandler.CheckEmailAvailablity)
	api.POST("/users", authMiddleware(authService, userService), userHandler.CreateUser)
	api.POST("/users/:ID", authMiddleware(authService, userService), userHandler.UpdateUser)
	api.GET("/users", authMiddleware(authService, userService), userHandler.GetUsers)
	api.GET("/users/:ID", authMiddleware(authService, userService), userHandler.GetUser)

	api.GET("/religions", authMiddleware(authService, userService), religionHandler.GetReligions)
	api.POST("/religions", authMiddleware(authService, userService), religionHandler.CreateReligion)
	api.POST("/religions/:ID", authMiddleware(authService, userService), religionHandler.UpdateReligion)
	api.GET("/religions/:ID", authMiddleware(authService, userService), religionHandler.GetReligion)
	api.DELETE("/religions/:ID", authMiddleware(authService, userService), religionHandler.DeleteReligion)

	api.GET("/educations", authMiddleware(authService, userService), educationHandler.GetEducations)
	api.POST("/educations", authMiddleware(authService, userService), educationHandler.CreateEducation)
	api.POST("/educations/:ID", authMiddleware(authService, userService), educationHandler.UpdateEducation)
	api.GET("/educations/:ID", authMiddleware(authService, userService), educationHandler.GetEducation)
	api.DELETE("/educations/:ID", authMiddleware(authService, userService), educationHandler.DeleteEducation)

	api.GET("/jobs", authMiddleware(authService, userService), jobHandler.GetJobs)
	api.POST("/jobs", authMiddleware(authService, userService), jobHandler.CreateJob)
	api.POST("/jobs/:ID", authMiddleware(authService, userService), jobHandler.UpdateJob)
	api.GET("/jobs/:ID", authMiddleware(authService, userService), jobHandler.GetJob)
	api.DELETE("/jobs/:ID", authMiddleware(authService, userService), jobHandler.DeleteJob)

	api.GET("/kelurahans", authMiddleware(authService, userService), kelurahanHandler.GetKelurahans)
	api.POST("/kelurahans", authMiddleware(authService, userService), kelurahanHandler.CreateKelurahan)
	api.POST("/kelurahans/:ID", authMiddleware(authService, userService), kelurahanHandler.UpdateKelurahan)
	api.GET("/kelurahans/:ID", authMiddleware(authService, userService), kelurahanHandler.GetKelurahan)
	api.DELETE("/kelurahans/:ID", authMiddleware(authService, userService), kelurahanHandler.DeleteKelurahan)

	api.GET("/provinces", authMiddleware(authService, userService), provinceHandler.GetProvinces)
	api.POST("/provinces", authMiddleware(authService, userService), provinceHandler.CreateProvince)
	api.POST("/provinces/:ID", authMiddleware(authService, userService), provinceHandler.UpdateProvince)
	api.GET("/provinces/:ID", authMiddleware(authService, userService), provinceHandler.GetProvince)
	api.DELETE("/provinces/:ID", authMiddleware(authService, userService), provinceHandler.DeleteProvince)

	api.GET("/districts", authMiddleware(authService, userService), districtHandler.GetDistricts)
	api.POST("/districts", authMiddleware(authService, userService), districtHandler.CreateDistrict)
	api.POST("/districts/:ID", authMiddleware(authService, userService), districtHandler.UpdateDistrict)
	api.GET("/districts/:ID", authMiddleware(authService, userService), districtHandler.GetDistrict)
	api.DELETE("/districts/:ID", authMiddleware(authService, userService), districtHandler.DeleteDistrict)

	api.GET("/subdistricts", authMiddleware(authService, userService), subDistrictHandler.GetSubDistricts)
	api.POST("/subdistricts", authMiddleware(authService, userService), subDistrictHandler.CreateSubDistrict)
	api.POST("/subdistricts/:ID", authMiddleware(authService, userService), subDistrictHandler.UpdateSubDistrict)
	api.GET("/subdistricts/:ID", authMiddleware(authService, userService), subDistrictHandler.GetSubDistrict)
	api.DELETE("/subdistricts/:ID", authMiddleware(authService, userService), subDistrictHandler.DeleteSubDistrict)

	api.GET("/rws", authMiddleware(authService, userService), rwHandler.GetRWs)
	api.POST("/rws", authMiddleware(authService, userService), rwHandler.CreateRW)
	api.POST("/rws/:ID", authMiddleware(authService, userService), rwHandler.UpdateRW)
	api.GET("/rws/:ID", authMiddleware(authService, userService), rwHandler.GetRW)
	api.DELETE("/rws/:ID", authMiddleware(authService, userService), rwHandler.DeleteRW)

	api.GET("/rts", authMiddleware(authService, userService), rtHandler.GetRTs)
	api.POST("/rts", authMiddleware(authService, userService), rtHandler.CreateRT)
	api.POST("/rts/:ID", authMiddleware(authService, userService), rtHandler.UpdateRT)
	api.GET("/rts/:ID", authMiddleware(authService, userService), rtHandler.GetRT)
	api.DELETE("/rts/:ID", authMiddleware(authService, userService), rtHandler.DeleteRT)

	api.GET("/positions", authMiddleware(authService, userService), positionHandler.GetPositions)
	api.POST("/positions", authMiddleware(authService, userService), positionHandler.CreatePosition)
	api.POST("/positions/:ID", authMiddleware(authService, userService), positionHandler.UpdatePosition)
	api.GET("/positions/:ID", authMiddleware(authService, userService), positionHandler.GetPositions)
	api.DELETE("/positions/:ID", authMiddleware(authService, userService), positionHandler.DeletePosition)

	api.GET("/articles", authMiddleware(authService, userService), articleHandler.GetArticles)
	api.POST("/articles", authMiddleware(authService, userService), articleHandler.CreateArticle)
	api.POST("/articles/:ID", authMiddleware(authService, userService), articleHandler.UpdateArticle)
	api.GET("/articles/:ID", authMiddleware(authService, userService), articleHandler.GetArticle)
	api.DELETE("/articles/:ID", authMiddleware(authService, userService), articleHandler.DeleteArticle)

	api.GET("/layanans", authMiddleware(authService, userService), layananHandler.GetLayanansGrouped)
	api.GET("/layanans/paging", authMiddleware(authService, userService), layananHandler.GetLayanans)
	api.GET("/layanans/rekom", authMiddleware(authService, userService), layananHandler.GetRekomLayanans)
	api.POST("/layanans", authMiddleware(authService, userService), layananHandler.CreateLayanan)
	api.POST("/layanans/:ID", authMiddleware(authService, userService), layananHandler.UpdateLayanan)
	api.GET("/layanans/:ID", authMiddleware(authService, userService), layananHandler.GetLayanan)
	api.DELETE("/layanans/:ID", authMiddleware(authService, userService), layananHandler.DeleteLayanan)

	api.GET("/pengajuans", authMiddleware(authService, userService), pengajuanHandler.GetPengajuanUser)
	api.GET("/pengajuans/admin", authMiddleware(authService, userService), pengajuanHandler.GetPengajuanAdmin)
	api.POST("/pengajuans", authMiddleware(authService, userService), pengajuanHandler.CreatePengajuan)
	api.POST("/pengajuans/:ID", authMiddleware(authService, userService), pengajuanHandler.UpdatePengajuan)
	api.GET("/pengajuans/:ID", authMiddleware(authService, userService), pengajuanHandler.GetPengajuan)
	api.DELETE("/pengajuans/:ID", authMiddleware(authService, userService), pengajuanHandler.DeletePengajuan)

	api.GET("/penduduks", authMiddleware(authService, userService), pendudukHandler.GetPenduduks)
	api.POST("/penduduks", authMiddleware(authService, userService), pendudukHandler.CreatePenduduk)
	api.POST("/penduduks/:ID", authMiddleware(authService, userService), pendudukHandler.UpdatePenduduk)
	api.GET("/penduduks/:ID", authMiddleware(authService, userService), pendudukHandler.GetPenduduk)
	api.DELETE("/penduduks/:ID", authMiddleware(authService, userService), pendudukHandler.DeletePenduduk)

	api.GET("/keluargas", authMiddleware(authService, userService), keluargaHandler.GetKeluargas)
	api.POST("/keluargas", authMiddleware(authService, userService), keluargaHandler.CreateKeluarga)
	api.POST("/keluargas/:ID", authMiddleware(authService, userService), keluargaHandler.UpdateKeluarga)
	api.GET("/keluargas/:ID", authMiddleware(authService, userService), keluargaHandler.GetKeluarga)
	api.DELETE("/keluargas/:ID", authMiddleware(authService, userService), keluargaHandler.DeleteKeluarga)

	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		tokenArray := strings.Split(authHeader, " ")
		if len(tokenArray) == 2 {
			tokenString = tokenArray[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		users, err := userService.GetUserById(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", users)
	}
}
