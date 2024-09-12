package mocking

import entity "belajar-golang-unit-test/mocking/Entity"

// jika kita ingin query ke db pakai go-lang itu disarankan untuk dibuat interface nya dulu jangan langsung
// dibuat didalam function supaya kedepannya itu bisa dilakukan mock testing, jadi nanti setiap function
// akan menjadi kontrak dari interface nya
type CategoryRepository interface {
	FindById(id string) *entity.Category
}
