package main

import (
	"encoding/json"
	"go-swag/model"
	"log"
	"net/http"
	"strconv"

	_ "go-swag/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @Summary ブログ記事リストを取得
// @Description 全てのブログ記事のリストを取得します。
// @Tags blog
// @Accept json
// @Produce json
// @Success 200 {array} model.BlogPost
// @Router /api/v1/blog/posts [get]
func GetBlogPosts(w http.ResponseWriter, r *http.Request) {
	// ダミーデータ
	posts := []model.BlogPost{
		{ID: 1, Title: "ブログ記事1", Content: "ブログ記事の内容1"},
		{ID: 2, Title: "ブログ記事2", Content: "ブログ記事の内容2"},
		// 他のブログ記事
	}
	json.NewEncoder(w).Encode(posts)
}

// @Summary 特定のブログ記事を取得
// @Description 指定されたIDのブログ記事を取得します。
// @Tags blog
// @Accept json
// @Produce json
// @Param id path int true "記事ID"
// @Success 200 {object} model.BlogPost
// @Router /api/v1/blog/posts/{id} [get]
func GetBlogPost(w http.ResponseWriter, r *http.Request) {
	// ダミーのブログ記事データ
	post := model.BlogPost{ID: 1, Title: "サンプルブログ記事", Content: "これはサンプルのブログ記事の内容です。"}

	// URLからIDを取得
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "不正な記事ID", http.StatusBadRequest)
		return
	}

	// 実際のアプリケーションでは、IDに基づいてデータベースから記事を取得します。
	// ここではダミーデータを返すだけです。
	post.ID = id

	json.NewEncoder(w).Encode(post)
}

// @Summary 新しいブログ記事を作成
// @Description 新しいブログ記事を作成します。
// @Tags blog
// @Accept json
// @Produce json
// @Param blogPost body model.BlogPost true "ブログ記事"
// @Success 201 {object} model.BlogPost
// @Router /api/v1/blog/posts [post]
func CreateBlogPost(w http.ResponseWriter, r *http.Request) {
	var newPost model.BlogPost

	// リクエストボディからブログ記事を読み込む
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, "リクエストの解析に失敗しました", http.StatusBadRequest)
		return
	}

	// 実際のアプリケーションでは、ここでデータベースに記事を保存します。
	// ここでは、送信された記事をそのまま返しています。
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}

// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
// @title Blog API
// @version 1.0
// @description This is a blog API.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := mux.NewRouter()
	// Swaggerドキュメントの設定
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	api := r.PathPrefix("/api/v1").Subrouter()
	// ブログ関連のエンドポイント
	api.HandleFunc("/blog/posts", GetBlogPosts).Methods("GET")
	api.HandleFunc("/blog/posts/{id}", GetBlogPost).Methods("GET")
	api.HandleFunc("/blog/posts", CreateBlogPost).Methods("POST")

	log.Println("server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", api))
}
