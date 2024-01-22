package usecase

type GenericUsecase[T any] interface {
	Create(payload T) T
	FindAll() []T
	FindByID(id uint) T
	Delete(id uint)
	Update(id uint, payload T) T
}
