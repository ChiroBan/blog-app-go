package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/pilinux/gorest/database"
)

type Tests struct {
	name string
	server *httptest.Server
	response *Posts
	expectedError error
}

func TestGetPost(t *testing.T) {

	tests := []Tests {
		{
			name: "get-post",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{ "id":1, "title": "test_title", "content": "test_content"}`))
			})),
			response: &Posts{
				Id: 1,
				title: "test_title",
				content: "test_content",
			},
			expectedError: nil,
		},
		{
			name: "get-all-posts",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{ "id":1, "title": "test_title", "content": "test_content"}`,
						`{ "id":2, "title": "test_title_2", "content": "test_content_2"}`))
			})),
			response: &Posts{
				Id: 1,
				title: "test_title",
				content: "test_content",
			},
			{
				Id: 2,
				title: "test_title_2",
				content: "test_content_2",
			}
			expectedError: nil,
		},
		{
			name: "post-post",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{ "id":1, "title": "test_title", "content": "test_content"}`))
			})),
			response: &Posts{
				Id: 1,
				title: "test_title",
				content: "test_content",
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			defer test.server.Close()

			resp, err := GetPost(test.server.URL)

			if !reflect.DeepEqual(resp, test.response) {
				t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
			}
			 if !errors.Is(err, test.expectedError) {
				t.Errorf("Expected error FAILED: expected %v got %v\n", test.expectedError, err)
			 }
		})
	}
}