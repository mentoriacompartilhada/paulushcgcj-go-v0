module github.com/mentoriacompartilhada/paulushcgcj-go-v0

go 1.16

require (
	github.com/andybalholm/brotli v1.0.3 // indirect
	github.com/gofiber/fiber/v2 v2.18.0 // indirect
	github.com/klauspost/compress v1.13.5 // indirect
	github.com/valyala/fasthttp v1.30.0 // indirect
	golang.org/x/sys v0.0.0-20210908233432-aa78b53d3365 // indirect
)

replace(
	github.com/mentoriacompartilhada/paulushcgcj-go-v0/controllers => "./src/controllers"
	github.com/mentoriacompartilhada/paulushcgcj-go-v0/models => "./src/models"
	github.com/mentoriacompartilhada/paulushcgcj-go-v0/services => "./src/services"
)