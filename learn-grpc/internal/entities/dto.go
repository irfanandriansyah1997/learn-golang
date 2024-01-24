package entities

import "context"

///////////////////////////////////////////////////////////
// Repository Section
///////////////////////////////////////////////////////////

type GenericRepo[arg, payload, id any] interface {
	Create(context.Context, arg) payload
	FindAll(context.Context) []payload
	FindByID(context.Context, id) (*payload, error)
	Delete(context.Context, id) bool
	Update(context.Context, id, arg) *payload
}

///////////////////////////////////////////////////////////
// Service Section
///////////////////////////////////////////////////////////

type GenericService[arg, payload, id any] interface {
	Create(context.Context, arg) payload
	FindAll(context.Context) []payload
	FindByID(context.Context, id) (payload, error)
	Delete(context.Context, id) bool
	Update(context.Context, id, arg) payload
}
