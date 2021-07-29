package data

import "github.com/aweist/grpc_server/blog"

type BlogPost struct {
	Title string
	Body  string
	Id    int
	Type  blog.BlogPost_Type
}

func GetSampleBlogPosts() []BlogPost {
	return []BlogPost{
		{
			Title: "First Blog Post",
			Body:  "This is the body of my first blog post",
			Id:    0,
			Type:  blog.BlogPost_TEXT,
		},
		{
			Title: "Second Blog Post",
			Body:  "This is the body of my second blog post",
			Id:    1,
			Type:  blog.BlogPost_TEXT,
		},
		{
			Title: "Third Blog Post",
			Body:  "This is the body of my third blog post",
			Id:    2,
			Type:  blog.BlogPost_TEXT,
		},
		{
			Title: "Fourth Blog Post",
			Body:  "This is a picture for my fourth blog post",
			Id:    3,
			Type:  blog.BlogPost_PICTURE,
		},
		{
			Title: "Fifth Blog Post",
			Body:  "This is a picture for my fifth blog post",
			Id:    4,
			Type:  blog.BlogPost_PICTURE,
		},
		{
			Title: "Sixth Blog Post",
			Body:  "This is a video for my sixth blog post",
			Id:    5,
			Type:  blog.BlogPost_VIDEO,
		},
	}
}
