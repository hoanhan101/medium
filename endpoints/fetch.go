package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hoanhan101/isocial/models/socialmedia"
)

// FetchPosts fetches all posts for a given user.
func FetchPosts(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement a real fetching mechanism. Below are just mock data.
	vars := mux.Vars(r)
	username := "hoanhan"
	if vars["username"] == username {
		mockPosts := make([]socialmedia.Post, 3)

		post1 := socialmedia.NewPost(
			"hoanhan",
			socialmedia.Moods["neutral"],
			"Caption 1",
			"Body 1",
			"URL 1",
			"Image URI 1",
			"Thumbnail URI 1",
			[]string{"key 1", "key 2", "key 3"},
		)

		post2 := socialmedia.NewPost(
			"hoanhan",
			socialmedia.Moods["happy"],
			"Caption 2",
			"Body 2",
			"URL 2",
			"Image URI 2",
			"Thumbnail URI 2",
			[]string{"key 1", "key 2", "key 3"},
		)

		post3 := socialmedia.NewPost(
			"hoanhan",
			socialmedia.Moods["sad"],
			"Caption 3",
			"Body 3",
			"URL 3",
			"Image URI 3",
			"Thumbnail URI 3",
			[]string{"key 1", "key 2", "key 3"},
		)

		mockPosts = append(mockPosts, *post1)
		mockPosts = append(mockPosts, *post2)
		mockPosts = append(mockPosts, *post3)

		json.NewEncoder(w).Encode(mockPosts)
	} else {
		json.NewEncoder(w).Encode(nil)
	}
}
