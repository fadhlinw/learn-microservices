package handler

import (
	"context"
	"log"

	dto "learn-microservices-server/src/app/dto/books"
	booksProto "learn-microservices-server/src/app/proto/books"
)

func (c *Handler) Book(ctx context.Context, req *booksProto.BookReq) (*booksProto.BookResp, error) {

	data := dto.BookReqDTO{
		Subject: req.Subject,
	}

	resp, err := c.useCases.BookUC.GetBooksBySubject(ctx, &data)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	dataResp, err := dto.TransformDTOToProto(resp)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dataResp, nil
}