package model

type Tests struct {
	name string
	query string
	response *Posts
	expectedError error
}

func TestDbPost(t *testing.T) {

	tests := []Tests {
		{
			name: "get-post",
			query: db.query(`select title, content from posts where id=$1`),
			response: &Posts{
				Id: 1,
				title: "test_title",
				content: "test_content",
			},
			expectedError: nil,
		},
		{
			name: "get-all-posts",
			query: db.query(`select id, title, content from posts;`),
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
			name: "create-post",
			query: db.query(`insert into posts(title, content) values($1, $2);`),
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

			resp, err := TestDbPost(test.server.URL)

			if !reflect.DeepEqual(resp, test.response) {
				t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
			}
			if !errors.Is(err, test.expectedError) {
				t.Errorf("Expected error FAILED: expected %v got %v\n", test.expectedError, err)
			}
		})
	}
}