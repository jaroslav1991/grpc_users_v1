package handlers

import (
	"errors"
	"fmt"
	"grpc_users_v1/internal/grpc/pb"
	"grpc_users_v1/internal/users"
	"io"
	"log"
	http "net/http"
	"text/template"
	"time"
)

func MainPageHandler(post pb.PostsClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templateData := map[string]any{}
		templateData["name"] = "anonymous"

		result, _ := r.Cookie("jwt")
		if result != nil {
			authClaims, err := users.ParseToken(result.Value)
			if err != nil {
				renderError(w, err)
				return
			}

			if err == nil {
				templateData["userId"] = authClaims.UserId
				templateData["name"] = authClaims.UserName

				resp, err := post.GetPosts(r.Context(), &pb.GetPostsRequest{UserId: authClaims.UserId})
				if err != nil {
					log.Println(err)
				} else {
					templateData["posts"] = resp.Posts
				}
			}
		}

		renderTemplate(w, "index", templateData)
	}
}

func renderTemplate(w io.Writer, templateName string, data any) {
	tmpl := template.New("templates")
	_, err := tmpl.ParseGlob("./internal/templates/*.html")
	if err != nil {
		log.Println(err)
		return
	}

	err = tmpl.ExecuteTemplate(w, templateName, data)
	if err != nil {
		log.Println(err)
		return
	}
}

func renderError(w io.Writer, err error) {
	renderTemplate(w, "error-500", map[string]interface{}{
		"error": err.Error(),
	})
}

func redirectToErrorPage(w http.ResponseWriter, value string) {
	w.Header().Add("Location", value)
	w.WriteHeader(http.StatusFound)
}

func redirectSuccess(w http.ResponseWriter, value string) {
	w.Header().Add("Location", value)
	w.WriteHeader(http.StatusFound)
}

func Registration(client pb.UsersClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if err := r.ParseForm(); err != nil {
				log.Println(err)
			}

			resp, err := client.SignUp(r.Context(), &pb.SignUpRequest{
				Name:         r.Form.Get("name"),
				Email:        r.Form.Get("email"),
				PasswordHash: r.Form.Get("password"),
			})
			if err != nil {
				log.Println(err)
				return
			}

			if resp.Error != nil && resp.Error.Code == 101 {
				redirectToErrorPage(w, "/error-password")
				return
			}

			if resp.Error != nil && resp.Error.Code == 102 {
				redirectToErrorPage(w, "/error-email")
				return
			}

			if resp.Error != nil && resp.Error.Code == 100 {
				redirectToErrorPage(w, "/error-signup")
				return
			}

			if resp.Error != nil && resp.Error.Code == 103 {
				redirectToErrorPage(w, "/error-signup")
				return
			} else {
				log.Println(resp)
				redirectSuccess(w, "/")
				return
			}
		}

		renderTemplate(w, "registration", nil)
	}

}

func SignIn(client pb.UsersClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if err := r.ParseForm(); err != nil {
				log.Println(err)
			}
			resp, err := client.SignIn(r.Context(), &pb.SignInRequest{
				Email:        r.Form.Get("email"),
				PasswordHash: r.Form.Get("password"),
			})
			if err != nil {
				log.Println(err)
			}
			if resp.Error != nil && resp.Error.Code == 102 {
				w.Header().Add("Location", "/error-signin")
				w.WriteHeader(http.StatusFound)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     "jwt",
				Value:    resp.Token.Token,
				Path:     "/",
				Expires:  time.Now().Add(10 * time.Hour),
				Secure:   true,
				HttpOnly: true,
			})

			redirectSuccess(w, "/")
		}

		renderTemplate(w, "signin", nil)
	}
}

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:    "jwt",
			Value:   "",
			Path:    "/",
			Expires: time.Now().Add(-10 * time.Hour),
		})
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func CreatePost(post pb.PostsClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			if err := r.ParseForm(); err != nil {
				log.Println(err)
				renderError(w, err)
				return
			}

			var userId int64

			jwtCookie, _ := r.Cookie("jwt")
			if jwtCookie != nil {
				authClaims, err := users.ParseToken(jwtCookie.Value)
				if err != nil {
					log.Println(err)
				}

				userId = authClaims.UserId
				fmt.Println(authClaims.UserId)
			}

			if userId == 0 {
				renderError(w, errors.New("нужна авторизация"))
				return
			}

			_, err := post.CreatePost(r.Context(), &pb.CreatePostRequest{
				Title:   r.Form.Get("title"),
				Message: r.Form.Get("message"),
				UserId:  userId,
			})
			if err != nil {
				w.Header().Add("Location", "/error-create-post")
				w.WriteHeader(http.StatusFound)
				return

			}

			redirectSuccess(w, "/")
			//w.Header().Add("Location", "/")
			//w.WriteHeader(http.StatusFound)
			return
		}

		renderTemplate(w, "create-post", nil)
	}
}

func ErrorSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "error-signin", nil)
	}
}

func ErrorDomain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validDomain := "@mail.ru, @gmail.com, @yandex.ru, @mail.com"

		renderTemplate(w, "error-signup", map[string]interface{}{
			"domain": validDomain,
		})
	}
}

func ErrorCreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "error-create-post", nil)
	}
}

func ErrorPassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "error-password", nil)
	}
}

func ErrorEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "error-email", nil)
	}
}
