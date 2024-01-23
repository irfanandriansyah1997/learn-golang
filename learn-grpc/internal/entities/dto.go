package entities

import "context"

///////////////////////////////////////////////////////////
// Repository Section
///////////////////////////////////////////////////////////

type GenericRepo[payload any, id any] interface {
	Create(context.Context, payload) payload
	FindAll(context.Context) []payload
	FindByID(context.Context, id) (payload, error)
	Delete(context.Context, id) bool
	Update(context.Context, id, payload) payload
}

///////////////////////////////////////////////////////////
// Service Section
///////////////////////////////////////////////////////////

type GenericService[payload any, id any] interface {
	Create(context.Context, payload) payload
	FindAll(context.Context) []payload
	FindByID(context.Context, id) (payload, error)
	Delete(context.Context, id) bool
	Update(context.Context, id, payload) payload
}
