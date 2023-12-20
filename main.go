package main

import (
	"context"
	"layanan-kependudukan-api/article"
	"layanan-kependudukan-api/auth"
	belumMenikah "layanan-kependudukan-api/belum_menikah"
	"layanan-kependudukan-api/berpergian"
	berpergianDetail "layanan-kependudukan-api/berpergian_detail"
	"layanan-kependudukan-api/district"
	"layanan-kependudukan-api/domisili"
	"layanan-kependudukan-api/education"
	"layanan-kependudukan-api/handler"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/janda"
	"layanan-kependudukan-api/job"
	"layanan-kependudukan-api/kelahiran"
	"layanan-kependudukan-api/keluarga"
	"layanan-kependudukan-api/kelurahan"
	"layanan-kependudukan-api/kematian"
	"layanan-kependudukan-api/kepolisian"
	"layanan-kependudukan-api/keramaian"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/penduduk"
	"layanan-kependudukan-api/pengajuan"
	pengajuanDetail "layanan-kependudukan-api/pengajuan_detail"
	"layanan-kependudukan-api/penghasilan"
	pernahMenikah "layanan-kependudukan-api/pernah_menikah"
	"layanan-kependudukan-api/pindah"
	pindahDetail "layanan-kependudukan-api/pindah_detail"
	"layanan-kependudukan-api/position"
	"layanan-kependudukan-api/province"
	"layanan-kependudukan-api/religion"
	"layanan-kependudukan-api/rt"
	"layanan-kependudukan-api/rumah"
	"layanan-kependudukan-api/rw"
	"layanan-kependudukan-api/sktm"
	"layanan-kependudukan-api/sku"
	"layanan-kependudukan-api/sporadik"
	"layanan-kependudukan-api/status"
	"layanan-kependudukan-api/subdistrict"
	"layanan-kependudukan-api/tanah"
	"layanan-kependudukan-api/user"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"

	"google.golang.org/api/option"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// dsn := "host=localhost user=postgres password=asdf1234 dbname=layanan-kependudukan port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	connection := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	app, _, _ := SetupFirebase()

	authService := auth.NewService()

	religionRepository := religion.NewRepsitory(db)
	religionService := religion.NewService(religionRepository)
	religionHandler := handler.NewReligionHandler(religionService, authService)

	statusRepository := status.NewRepsitory(db)
	statusService := status.NewService(statusRepository)
	statusHandler := handler.NewStatusHandler(statusService, authService)

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

	detailPengajuanRepository := pengajuanDetail.NewRepsitory(db)
	detailPengajuanService := pengajuanDetail.NewService(detailPengajuanRepository)

	pendudukRepository := penduduk.NewRepsitory(db)
	pendudukService := penduduk.NewService(pendudukRepository)
	pendudukHandler := handler.NewPendudukHandler(pendudukService, authService)

	keluargaRepository := keluarga.NewRepsitory(db)
	keluargaService := keluarga.NewService(keluargaRepository)
	keluargaHandler := handler.NewKeluargaHandler(keluargaService, pendudukService, authService)

	userRepository := user.NewRepsitory(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, keluargaService, pendudukService, authService)

	subDistrictRepository := subdistrict.NewRepsitory(db)
	subDistrictService := subdistrict.NewService(subDistrictRepository)
	subDistrictHandler := handler.NewSubDistrictHandler(subDistrictService, authService)

	pengajuanRepository := pengajuan.NewRepsitory(db)
	pengajuanService := pengajuan.NewService(pengajuanRepository)
	pengajuanHandler := handler.NewPengajuanHandler(app, pengajuanService, detailPengajuanService, layananService, userService, pendudukService, authService)

	dashboardHandler := handler.NewDashboardHandler(pengajuanService, pendudukService, keluargaService, authService)

	sktmRepository := sktm.NewRepsitory(db)
	sktmService := sktm.NewService(sktmRepository)
	sktmHandler := handler.NewSKTMHandler(sktmService, layananService, *pengajuanHandler, authService)

	skuRepository := sku.NewRepsitory(db)
	skuService := sku.NewService(skuRepository)
	skuHandler := handler.NewSKUHandler(skuService, layananService, *pengajuanHandler, authService)

	domisiliRepository := domisili.NewRepsitory(db)
	domisiliService := domisili.NewService(domisiliRepository)
	domisiliHandler := handler.NewDomisiliHandler(domisiliService, layananService, *pengajuanHandler, authService)

	keramaianRepository := keramaian.NewRepsitory(db)
	keramaianService := keramaian.NewService(keramaianRepository)
	keramaianHandler := handler.NewKeramaianHandler(keramaianService, layananService, *pengajuanHandler, authService)

	kelahiranRepository := kelahiran.NewRepsitory(db)
	kelahiranService := kelahiran.NewService(kelahiranRepository)
	kelahiranHandler := handler.NewKelahiranHandler(kelahiranService, layananService, *pengajuanHandler, authService)

	kematianRepository := kematian.NewRepsitory(db)
	kematianService := kematian.NewService(kematianRepository)
	kematianHandler := handler.NewKematianHandler(kematianService, layananService, *pengajuanHandler, authService)

	berpergianDetailRepository := berpergianDetail.NewRepsitory(db)
	berpergianDetailService := berpergianDetail.NewService(berpergianDetailRepository)
	berpergianDetailHandler := handler.NewBerpergianDetailHandler(berpergianDetailService, authService)

	berpergianRepository := berpergian.NewRepsitory(db)
	berpergianService := berpergian.NewService(berpergianRepository)
	berpergianHandler := handler.NewBerpergianHandler(berpergianService, layananService, *pengajuanHandler, authService)

	pindahDetailRepository := pindahDetail.NewRepsitory(db)
	pindahDetailService := pindahDetail.NewService(pindahDetailRepository)
	pindahDetailHandler := handler.NewPindahDetailHandler(pindahDetailService, authService)

	pindahRepository := pindah.NewRepsitory(db)
	pindahService := pindah.NewService(pindahRepository)
	pindahHandler := handler.NewPindahHandler(pindahService, layananService, *pengajuanHandler, pindahDetailService, authService)

	jandaRepository := janda.NewRepsitory(db)
	jandaService := janda.NewService(jandaRepository)
	jandaHandler := handler.NewJandaHandler(jandaService, layananService, *pengajuanHandler, authService)

	penghasilanRepository := penghasilan.NewRepsitory(db)
	penghasilanService := penghasilan.NewService(penghasilanRepository)
	penghasilanHandler := handler.NewPenghasilanHandler(penghasilanService, layananService, *pengajuanHandler, authService)

	belumMenikahRepository := belumMenikah.NewRepsitory(db)
	belumMenikahService := belumMenikah.NewService(belumMenikahRepository)
	belumMenikahHandler := handler.NewBelumMenikahHandler(belumMenikahService, layananService, *pengajuanHandler, authService)

	pernahMenikahRepository := pernahMenikah.NewRepsitory(db)
	pernahMenikahService := pernahMenikah.NewService(pernahMenikahRepository)
	pernahMenikahHandler := handler.NewPernahMenikahHandler(pernahMenikahService, pendudukService, layananService, *pengajuanHandler, authService)

	kepolisianRepository := kepolisian.NewRepsitory(db)
	kepolisianService := kepolisian.NewService(kepolisianRepository)
	kepolisianHandler := handler.NewKepolisianHandler(kepolisianService, layananService, *pengajuanHandler, authService)

	rumahRepository := rumah.NewRepsitory(db)
	rumahService := rumah.NewService(rumahRepository)
	rumahHandler := handler.NewRumahHandler(rumahService, layananService, *pengajuanHandler, authService)

	tanahRepository := tanah.NewRepsitory(db)
	tanahService := tanah.NewService(tanahRepository)
	tanahHandler := handler.NewTanahHandler(tanahService, layananService, *pengajuanHandler, authService)

	sporadikRepository := sporadik.NewRepsitory(db)
	sporadikService := sporadik.NewService(sporadikRepository)
	sporadikHandler := handler.NewSporadikHandler(sporadikService, layananService, *pengajuanHandler, authService)

	router := gin.Default()

	corsMiddleware := cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS", "CONNECT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-CSRF-Token", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	})
	router.Use(corsMiddleware)
	router.Static("/images", "./images")
	router.Static("/documents", "./documents")
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegiserUser)
	api.POST("/login", userHandler.Login)
	api.POST("/logout", authMiddleware(authService, userService), userHandler.Logout)
	api.POST("/email_checkers", authMiddleware(authService, userService), userHandler.CheckEmailAvailablity)
	api.POST("/users", authMiddleware(authService, userService), userHandler.CreateUser)
	api.POST("/users/:ID", authMiddleware(authService, userService), userHandler.UpdateUser)
	api.GET("/users", authMiddleware(authService, userService), userHandler.GetUsers)
	api.GET("/users/:ID", authMiddleware(authService, userService), userHandler.GetUser)
	api.DELETE("/users/:ID", authMiddleware(authService, userService), userHandler.DeleteUser)

	api.GET("/religions", authMiddleware(authService, userService), religionHandler.GetReligions)
	api.POST("/religions", authMiddleware(authService, userService), religionHandler.CreateReligion)
	api.POST("/religions/:ID", authMiddleware(authService, userService), religionHandler.UpdateReligion)
	api.GET("/religions/:ID", authMiddleware(authService, userService), religionHandler.GetReligion)
	api.DELETE("/religions/:ID", authMiddleware(authService, userService), religionHandler.DeleteReligion)

	api.GET("/status", authMiddleware(authService, userService), statusHandler.GetStatuss)
	api.POST("/status", authMiddleware(authService, userService), statusHandler.CreateStatus)
	api.POST("/status/:ID", authMiddleware(authService, userService), statusHandler.UpdateStatus)
	api.GET("/status/:ID", authMiddleware(authService, userService), statusHandler.GetStatus)
	api.DELETE("/status/:ID", authMiddleware(authService, userService), statusHandler.DeleteStatus)

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

	api.GET("/penduduks", authMiddleware(authService, userService), pendudukHandler.GetPenduduks)
	api.POST("/penduduks", authMiddleware(authService, userService), pendudukHandler.CreatePenduduk)
	api.POST("/penduduks/:ID", authMiddleware(authService, userService), pendudukHandler.UpdatePenduduk)
	api.GET("/penduduks/:ID", authMiddleware(authService, userService), pendudukHandler.GetPenduduk)
	api.DELETE("/penduduks/:ID", authMiddleware(authService, userService), pendudukHandler.DeletePenduduk)

	api.GET("/dashboards", authMiddleware(authService, userService), dashboardHandler.GetDashboard)

	api.GET("/keluargas", authMiddleware(authService, userService), keluargaHandler.GetKeluargas)
	api.POST("/keluargas", authMiddleware(authService, userService), keluargaHandler.CreateKeluarga)
	api.POST("/keluargas/:ID", authMiddleware(authService, userService), keluargaHandler.UpdateKeluarga)
	api.GET("/keluargas/:ID", authMiddleware(authService, userService), keluargaHandler.GetKeluarga)
	api.GET("/keluargas/user", authMiddleware(authService, userService), keluargaHandler.GetKeluargaByUser)
	api.DELETE("/keluargas/:ID", authMiddleware(authService, userService), keluargaHandler.DeleteKeluarga)

	api.GET("/pengajuans", authMiddleware(authService, userService), pengajuanHandler.GetPengajuanUser)
	api.GET("/pengajuans/admin", authMiddleware(authService, userService), pengajuanHandler.GetPengajuanAdmin)
	// api.POST("/pengajuans", authMiddleware(authService, userService), pengajuanHandler.CreatePengajuan)
	api.POST("/pengajuans/:ID", authMiddleware(authService, userService), pengajuanHandler.UpdatePengajuan)
	api.GET("/pengajuans/:ID", authMiddleware(authService, userService), pengajuanHandler.GetPengajuan)
	api.DELETE("/pengajuans/:ID", authMiddleware(authService, userService), pengajuanHandler.DeletePengajuan)

	api.GET("/sktms", authMiddleware(authService, userService), sktmHandler.GetSKTMs)
	api.POST("/sktms", authMiddleware(authService, userService), sktmHandler.CreateSKTM)
	api.POST("/sktms/:ID", authMiddleware(authService, userService), sktmHandler.UpdateSKTM)
	api.GET("/sktms/:ID", authMiddleware(authService, userService), sktmHandler.GetSKTM)
	api.DELETE("/sktms/:ID", authMiddleware(authService, userService), sktmHandler.DeleteSKTM)

	api.GET("/skus", authMiddleware(authService, userService), skuHandler.GetSKUs)
	api.POST("/skus", authMiddleware(authService, userService), skuHandler.CreateSKU)
	api.POST("/skus/:ID", authMiddleware(authService, userService), skuHandler.UpdateSKU)
	api.GET("/skus/:ID", authMiddleware(authService, userService), skuHandler.GetSKU)
	api.DELETE("/skus/:ID", authMiddleware(authService, userService), skuHandler.DeleteSKU)

	api.GET("/domisilis", authMiddleware(authService, userService), domisiliHandler.GetDomisilis)
	api.POST("/domisilis", authMiddleware(authService, userService), domisiliHandler.CreateDomisili)
	api.POST("/domisilis/:ID", authMiddleware(authService, userService), domisiliHandler.UpdateDomisili)
	api.GET("/domisilis/:ID", authMiddleware(authService, userService), domisiliHandler.GetDomisili)
	api.DELETE("/domisilis/:ID", authMiddleware(authService, userService), domisiliHandler.DeleteDomisili)

	api.GET("/keramaians", authMiddleware(authService, userService), keramaianHandler.GetKeramaians)
	api.POST("/keramaians", authMiddleware(authService, userService), keramaianHandler.CreateKeramaian)
	api.POST("/keramaians/:ID", authMiddleware(authService, userService), keramaianHandler.UpdateKeramaian)
	api.GET("/keramaians/:ID", authMiddleware(authService, userService), keramaianHandler.GetKeramaian)
	api.DELETE("/keramaians/:ID", authMiddleware(authService, userService), keramaianHandler.DeleteKeramaian)

	api.GET("/kelahirans", authMiddleware(authService, userService), kelahiranHandler.GetKelahirans)
	api.POST("/kelahirans", authMiddleware(authService, userService), kelahiranHandler.CreateKelahiran)
	api.POST("/kelahirans/:ID", authMiddleware(authService, userService), kelahiranHandler.UpdateKelahiran)
	api.GET("/kelahirans/:ID", authMiddleware(authService, userService), kelahiranHandler.GetKelahiran)
	api.DELETE("/kelahirans/:ID", authMiddleware(authService, userService), kelahiranHandler.DeleteKelahiran)

	api.GET("/kematians", authMiddleware(authService, userService), kematianHandler.GetKematians)
	api.POST("/kematians", authMiddleware(authService, userService), kematianHandler.CreateKematian)
	api.POST("/kematians/:ID", authMiddleware(authService, userService), kematianHandler.UpdateKematian)
	api.GET("/kematians/:ID", authMiddleware(authService, userService), kematianHandler.GetKematian)
	api.DELETE("/kematians/:ID", authMiddleware(authService, userService), kematianHandler.DeleteKematian)

	api.GET("/berpergians", authMiddleware(authService, userService), berpergianHandler.GetBerpergians)
	api.POST("/berpergians", authMiddleware(authService, userService), berpergianHandler.CreateBerpergian)
	api.POST("/berpergians/:ID", authMiddleware(authService, userService), berpergianHandler.UpdateBerpergian)
	api.GET("/berpergians/:ID", authMiddleware(authService, userService), berpergianHandler.GetBerpergian)
	api.DELETE("/berpergians/:ID", authMiddleware(authService, userService), berpergianHandler.DeleteBerpergian)

	api.POST("/berpergiands", authMiddleware(authService, userService), berpergianDetailHandler.CreateBerpergianDetail)

	api.GET("/pindahs", authMiddleware(authService, userService), pindahHandler.GetPindahs)
	api.POST("/pindahs", authMiddleware(authService, userService), pindahHandler.CreatePindah)
	api.POST("/pindahs/:ID", authMiddleware(authService, userService), pindahHandler.UpdatePindah)
	api.GET("/pindahs/:ID", authMiddleware(authService, userService), pindahHandler.GetPindah)
	api.DELETE("/pindahs/:ID", authMiddleware(authService, userService), pindahHandler.DeletePindah)

	api.POST("/pindahds", authMiddleware(authService, userService), pindahDetailHandler.CreatePindahDetail)

	api.GET("/jandas", authMiddleware(authService, userService), jandaHandler.GetJandas)
	api.POST("/jandas", authMiddleware(authService, userService), jandaHandler.CreateJanda)
	api.POST("/jandas/:ID", authMiddleware(authService, userService), jandaHandler.UpdateJanda)
	api.GET("/jandas/:ID", authMiddleware(authService, userService), jandaHandler.GetJanda)
	api.DELETE("/jandas/:ID", authMiddleware(authService, userService), jandaHandler.DeleteJanda)

	api.GET("/penghasilans", authMiddleware(authService, userService), penghasilanHandler.GetPenghasilans)
	api.POST("/penghasilans", authMiddleware(authService, userService), penghasilanHandler.CreatePenghasilan)
	api.POST("/penghasilans/:ID", authMiddleware(authService, userService), penghasilanHandler.UpdatePenghasilan)
	api.GET("/penghasilans/:ID", authMiddleware(authService, userService), penghasilanHandler.GetPenghasilan)
	api.DELETE("/penghasilans/:ID", authMiddleware(authService, userService), penghasilanHandler.DeletePenghasilan)

	api.GET("/belum_menikahs", authMiddleware(authService, userService), belumMenikahHandler.GetBelumMenikahs)
	api.POST("/belum_menikahs", authMiddleware(authService, userService), belumMenikahHandler.CreateBelumMenikah)
	api.POST("/belum_menikahs/:ID", authMiddleware(authService, userService), belumMenikahHandler.UpdateBelumMenikah)
	api.GET("/belum_menikahs/:ID", authMiddleware(authService, userService), belumMenikahHandler.GetBelumMenikah)
	api.DELETE("/belum_menikahs/:ID", authMiddleware(authService, userService), belumMenikahHandler.DeleteBelumMenikah)

	api.GET("/pernah_menikahs", authMiddleware(authService, userService), pernahMenikahHandler.GetPernahMenikahs)
	api.POST("/pernah_menikahs", authMiddleware(authService, userService), pernahMenikahHandler.CreatePernahMenikah)
	api.POST("/pernah_menikahs/:ID", authMiddleware(authService, userService), pernahMenikahHandler.UpdatePernahMenikah)
	api.GET("/pernah_menikahs/:ID", authMiddleware(authService, userService), pernahMenikahHandler.GetPernahMenikah)
	api.DELETE("/pernah_menikahs/:ID", authMiddleware(authService, userService), pernahMenikahHandler.DeletePernahMenikah)

	api.GET("/kepolisians", authMiddleware(authService, userService), kepolisianHandler.GetKepolisians)
	api.POST("/kepolisians", authMiddleware(authService, userService), kepolisianHandler.CreateKepolisian)
	api.POST("/kepolisians/:ID", authMiddleware(authService, userService), kepolisianHandler.UpdateKepolisian)
	api.GET("/kepolisians/:ID", authMiddleware(authService, userService), kepolisianHandler.GetKepolisian)
	api.DELETE("/kepolisians/:ID", authMiddleware(authService, userService), kepolisianHandler.DeleteKepolisian)

	api.GET("/rumahs", authMiddleware(authService, userService), rumahHandler.GetRumahs)
	api.POST("/rumahs", authMiddleware(authService, userService), rumahHandler.CreateRumah)
	api.POST("/rumahs/:ID", authMiddleware(authService, userService), rumahHandler.UpdateRumah)
	api.GET("/rumahs/:ID", authMiddleware(authService, userService), rumahHandler.GetRumah)
	api.DELETE("/rumahs/:ID", authMiddleware(authService, userService), rumahHandler.DeleteRumah)

	api.GET("/tanahs", authMiddleware(authService, userService), tanahHandler.GetTanahs)
	api.POST("/tanahs", authMiddleware(authService, userService), tanahHandler.CreateTanah)
	api.POST("/tanahs/:ID", authMiddleware(authService, userService), tanahHandler.UpdateTanah)
	api.GET("/tanahs/:ID", authMiddleware(authService, userService), tanahHandler.GetTanah)
	api.DELETE("/tanahs/:ID", authMiddleware(authService, userService), tanahHandler.DeleteTanah)

	api.GET("/sporadiks", authMiddleware(authService, userService), sporadikHandler.GetSporadiks)
	api.POST("/sporadiks", authMiddleware(authService, userService), sporadikHandler.CreateSporadik)
	api.POST("/sporadiks/:ID", authMiddleware(authService, userService), sporadikHandler.UpdateSporadik)
	api.GET("/sporadiks/:ID", authMiddleware(authService, userService), sporadikHandler.GetSporadik)
	api.DELETE("/sporadiks/:ID", authMiddleware(authService, userService), sporadikHandler.DeleteSporadik)

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

func SetupFirebase() (*firebase.App, context.Context, *messaging.Client) {

	ctx := context.Background()

	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKeys.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	//Messaging client
	client, _ := app.Messaging(ctx)

	return app, ctx, client
}
