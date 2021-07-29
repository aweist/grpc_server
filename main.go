package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/aweist/grpc_server/blog"
	"github.com/aweist/grpc_server/data"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8080, "This is the port the server will run on ")
)

func main() {
	flag.Parse()
	fmt.Println("gRPC server running on port", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen on port %v: %v", *port, err)
	}

	grpcServer := grpc.NewServer()
	blog.RegisterBlogServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

func newServer() *blogServer {
	server := &blogServer{
		blogPosts: make(map[int]blog.BlogPost),
	}
	server.LoadBlogPosts(data.GetSampleBlogPosts())
	return server
}

type blogServer struct {
	blog.UnimplementedBlogServiceServer
	blogPosts map[int]blog.BlogPost
}

func (bs blogServer) SearchBlog(context.Context, *blog.BlogSearch) (*blog.BlogSearchResults, error) {
	return nil, nil
}

func (bs blogServer) GetPost(c context.Context, blogPostById *blog.BlogPostById) (*blog.BlogPost, error) {
	id := int(blogPostById.GetId())
	blogPost, ok := bs.blogPosts[id]
	if !ok {
		return nil, fmt.Errorf("Could not find blog post with ID %d", id)
	}

	return &blogPost, nil
}

func (bs *blogServer) LoadBlogPosts(blogPosts []data.BlogPost) {
	for _, blogPost := range data.GetSampleBlogPosts() {
		bs.blogPosts[int(blogPost.Id)] = blog.BlogPost{
			Title: blogPost.Title,
			Body:  blogPost.Body,
			Id:    int32(blogPost.Id),
			Type:  blogPost.Type,
		}
	}
}
