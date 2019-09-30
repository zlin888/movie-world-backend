// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"hello-beego-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/movie/all", &controllers.MovieController{}, "get:GetMovies")
	beego.Router("/api/movie/add", &controllers.MovieController{}, "put:AddMovie")
	beego.Router("/api/movie/delete/:id", &controllers.MovieController{}, "delete:DeleteMovie")
	beego.Router("/api/movie/update/:id", &controllers.MovieController{}, "post:UpdateMovie")
}
