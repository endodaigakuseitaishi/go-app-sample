package controllers

import (
	"go-todo-sample/app/models"
	"go-todo-sample/app/services"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		// エラーがあればログインしていない
		generateHTML(w, "hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		// セッション情報からuserを取得し、ユーザーに紐づくtodo一覧を返す
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		content := r.PostFormValue("content")
		if err := services.CheckContentLength(content); err != nil {
			errorMessage := err.Error()

			// エラーメッセージをテンプレートに渡して表示
			data := struct {
				ErrorMessage string
			}{
				ErrorMessage: errorMessage,
			}

			generateHTML(w, data, "layout", "private_navbar", "todo_new")
			return
		}

		if err := user.CreateTodo(content); err != nil {
			log.Fatalln(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}

// requestごとにURLからIDを取得して処理を振り分ける
func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Fatalln(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		content := r.PostFormValue("content")
		t := &models.Todo{ID: id, Content: content, UserId: user.ID}
		if err := t.UpdateTodo(); err != nil {
			log.Fatalln(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Fatalln(err)
		}
		if err := t.DeleteTodo(); err != nil {
			log.Fatalln(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}