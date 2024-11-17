package models

type Error struct {
	Message string `json:"message"`
}

type AlertMessage struct {
	Message string `json:"message"`
}

const (
	RequiredRefreshMessage = "Required reshresh"
	NoAccessMessage        = "You have no access this page"
	TokenInvalid           = "Invalid token claims"

	WrongLogin           = "Login xato"
	WrongPassword        = "Parol xato"
	WrongLoginOrPassword = "Login yoki Parol xato"
	NotEqualConfirm      = "Parol bir xil emas"
	WeakPassword         = "Parol uzunligi 8 dan 30 gacha bo'lishi kerak"
	WrongDateMessage     = "Sanada xatolik mavjud"
	WrongInfoMessage     = "Ma'lumotlarda xatolik mavjud"
	AlreadyAdded         = "Allaqachon qo'shilgan"

	NotFoundMessage   = "Ma'lumot topilmadi"
	NotCreatedMessage = "Yaratilmadi"
	NotUpdatedMessage = "O'zgartirilmadi"
	NotDeletedMessage = "O'chirilmadi"
	NotAddedMessage   = "Qo'shilmadi"
	InternalMessage   = "Xatolik yuzaga keldi"
)
