// Code generated by goctl. DO NOT EDIT!
// Source: book.proto

package bookclient

import (
	"context"

	"go-zero-admin-server/service/book/rpc/book"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	IdsReq         = book.IdsReq
	BookInfoReply  = book.BookInfoReply
	TypesReq       = book.TypesReq
	BookAddReq     = book.BookAddReq
	TypeInfoReply  = book.TypeInfoReply
	TypesInfoReply = book.TypesInfoReply
	IsExistReply   = book.IsExistReply
	NameReq        = book.NameReq
	TypeUpdateReq  = book.TypeUpdateReq
	TypeAddReq     = book.TypeAddReq
	BooksReq       = book.BooksReq
	BookUpdateReq  = book.BookUpdateReq
	BooksInfoReply = book.BooksInfoReply
	EmptyReq       = book.EmptyReq
	IdReq          = book.IdReq
	IsSuccessReply = book.IsSuccessReply

	Book interface {
		GetBookById(ctx context.Context, in *IdReq) (*BookInfoReply, error)
		GetBookByNameLike(ctx context.Context, in *NameReq) (*BooksInfoReply, error)
		IsExistBookById(ctx context.Context, in *IdReq) (*IsExistReply, error)
		IsExistBookByName(ctx context.Context, in *NameReq) (*IsExistReply, error)
		GetBooks(ctx context.Context, in *BooksReq) (*BooksInfoReply, error)
		AddBook(ctx context.Context, in *BookAddReq) (*IsSuccessReply, error)
		UpdateBook(ctx context.Context, in *BookUpdateReq) (*IsSuccessReply, error)
		DelBook(ctx context.Context, in *IdReq) (*IsSuccessReply, error)
		DelSomeBook(ctx context.Context, in *IdsReq) (*IsSuccessReply, error)
		GetTypeById(ctx context.Context, in *IdReq) (*TypeInfoReply, error)
		GetAllTypes(ctx context.Context, in *EmptyReq) (*TypesInfoReply, error)
		GetTypeByNameLike(ctx context.Context, in *NameReq) (*TypesInfoReply, error)
		IsExistTypeById(ctx context.Context, in *IdReq) (*IsExistReply, error)
		IsExistTypeByName(ctx context.Context, in *NameReq) (*IsExistReply, error)
		GetTypes(ctx context.Context, in *TypesReq) (*TypesInfoReply, error)
		AddType(ctx context.Context, in *TypeAddReq) (*IsSuccessReply, error)
		UpdateType(ctx context.Context, in *TypeUpdateReq) (*IsSuccessReply, error)
		DelType(ctx context.Context, in *IdReq) (*IsSuccessReply, error)
	}

	defaultBook struct {
		cli zrpc.Client
	}
)

func NewBook(cli zrpc.Client) Book {
	return &defaultBook{
		cli: cli,
	}
}

func (m *defaultBook) GetBookById(ctx context.Context, in *IdReq) (*BookInfoReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.GetBookById(ctx, in)
}

func (m *defaultBook) GetBookByNameLike(ctx context.Context, in *NameReq) (*BooksInfoReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.GetBookByNameLike(ctx, in)
}

func (m *defaultBook) IsExistBookById(ctx context.Context, in *IdReq) (*IsExistReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.IsExistBookById(ctx, in)
}

func (m *defaultBook) IsExistBookByName(ctx context.Context, in *NameReq) (*IsExistReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.IsExistBookByName(ctx, in)
}

func (m *defaultBook) GetBooks(ctx context.Context, in *BooksReq) (*BooksInfoReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.GetBooks(ctx, in)
}

func (m *defaultBook) AddBook(ctx context.Context, in *BookAddReq) (*IsSuccessReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.AddBook(ctx, in)
}

func (m *defaultBook) UpdateBook(ctx context.Context, in *BookUpdateReq) (*IsSuccessReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.UpdateBook(ctx, in)
}

func (m *defaultBook) DelBook(ctx context.Context, in *IdReq) (*IsSuccessReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.DelBook(ctx, in)
}

func (m *defaultBook) DelSomeBook(ctx context.Context, in *IdsReq) (*IsSuccessReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.DelSomeBook(ctx, in)
}

func (m *defaultBook) GetTypeById(ctx context.Context, in *IdReq) (*TypeInfoReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.GetTypeById(ctx, in)
}

func (m *defaultBook) GetAllTypes(ctx context.Context, in *EmptyReq) (*TypesInfoReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.GetAllTypes(ctx, in)
}

func (m *defaultBook) GetTypeByNameLike(ctx context.Context, in *NameReq) (*TypesInfoReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.GetTypeByNameLike(ctx, in)
}

func (m *defaultBook) IsExistTypeById(ctx context.Context, in *IdReq) (*IsExistReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.IsExistTypeById(ctx, in)
}

func (m *defaultBook) IsExistTypeByName(ctx context.Context, in *NameReq) (*IsExistReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.IsExistTypeByName(ctx, in)
}

func (m *defaultBook) GetTypes(ctx context.Context, in *TypesReq) (*TypesInfoReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.GetTypes(ctx, in)
}

func (m *defaultBook) AddType(ctx context.Context, in *TypeAddReq) (*IsSuccessReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.AddType(ctx, in)
}

func (m *defaultBook) UpdateType(ctx context.Context, in *TypeUpdateReq) (*IsSuccessReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.UpdateType(ctx, in)
}

func (m *defaultBook) DelType(ctx context.Context, in *IdReq) (*IsSuccessReply, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.DelType(ctx, in)
}
