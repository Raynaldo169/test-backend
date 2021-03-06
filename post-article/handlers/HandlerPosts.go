package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"post-article/connection"
	"post-article/structs"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Have a nice day, Raynaldo!")
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var article structs.Posts
	article.Created_date = time.Now()
	article.Updated_date = time.Now()

	json.Unmarshal(payloads, &article)

	validate := validator.New()
	err := validate.Struct(article)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
		http.Error(w, "The article cannot be created!", http.StatusUnprocessableEntity)
	} else {
		connection.DB.Create(&article)
		res := structs.Result{Code: 200, Data: article, Message: "Success create article"}
		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit := vars["limit"]
	offset := vars["offset"]
	articles := []structs.Posts{}

	connection.DB.
		Limit(limit).
		Offset(offset).
		Order("updated_date").
		Find(&articles)

	res := structs.Result{Code: 200, Data: articles, Message: "Success get articles"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID := vars["id"]

	var article structs.Posts

	connection.DB.First(&article, articleID)
	sql := "select * from posts where id='" + articleID + "'"
	connection.DB.Exec(sql)
	if article.Title == "" {
		http.Error(w, "Article not found!", http.StatusNotFound)
	} else {
		res := structs.Result{Code: 200, Data: article, Message: "Success get article"}
		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var articleUpdates structs.Posts
	articleUpdates.Updated_date = time.Now()

	json.Unmarshal(payloads, &articleUpdates)

	var article structs.Posts

	connection.DB.First(&article, articleID)
	sql := "select * from posts where id='" + articleID + "'"
	connection.DB.Exec(sql)
	if article.Title == "" {
		http.Error(w, "Article not found!", http.StatusNotFound)
	} else {
		validate := validator.New()
		err := validate.Struct(articleUpdates)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err)
			}
			http.Error(w, "Can't update article!", http.StatusUnprocessableEntity)
		} else {
			connection.DB.Model(&article).Updates(&articleUpdates)
			res := structs.Result{Code: 200, Data: article, Message: "Success update article"}
			result, err := json.Marshal(res)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID := vars["id"]

	var article structs.Posts
	connection.DB.First(&article, articleID)
	sql := "select * from posts where id='" + articleID + "'"
	connection.DB.Exec(sql)
	if article.Title == "" {
		http.Error(w, "Article not found!", http.StatusNotFound)
	} else {
		connection.DB.Delete(&article)

		res := structs.Result{Code: 200, Message: "Success delete article"}
		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}
