package main

import (
	"fmt"
	"unsafe"
)

// p := *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(tmp)) + unsafe.Sizeof(s[0])*2))

// Порядок работы с арифметикой указателей:
// 1. Приведение указателя tmp в тип unsafe.Pointer - unsafe.Pointer(tmp)
// 2. Приведение типа unsafe.Pointer к типу uintptr - uintptr(unsafe.Pointer(tmp))
// TOGO: проверить на структуре: (получение начального адреса в виде uintptr)
// тип данных позволяющий вместить в себя все что угодно
// 3. Получаем размер одного элемента(для слайса) - unsafe.Sizeof(s[0])
// 4. Получаем значение на которое необходимо сделать сдвиг - unsafe.Sizeof(s[0])*2
// 5. Прибавляем к значению uintptr сдвиг - uintptr(unsafe.Pointer(tmp)) + unsafe.Sizeof(s[0])*2
// 6. Приводим тип uintptr снова в тип unsafe.Pointer - unsafe.Pointer(uintptr(unsafe.Pointer(tmp)) + unsafe.Sizeof(s[0])*2)
// 7. Оборачиваем приведение в скобки чтобы полностью все преобразования были сделаны
// - (unsafe.Pointer(uintptr(unsafe.Pointer(tmp)) + unsafe.Sizeof(s[0])*2))
// 8. Приводим тип unsafe.Pointer к необходимому нам типу указателя int32
// - (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(tmp)) + unsafe.Sizeof(s[0])*2))
// 9. Разыменовываем полученный указатель в значение этого типа
// - *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(tmp)) + unsafe.Sizeof(s[0])*2))

type SetTypes struct {
	First  bool
	Second int32
	Third  string
	Fourth bool
}

func main() {
	settypes := &SetTypes{First: true, Second: 777, Third: "Hello World !!!", Fourth: false}

	p := unsafe.Pointer(settypes)
	p1 := *(*bool)(unsafe.Pointer(uintptr(p)))
	fmt.Println(p1)

	fmt.Println(*(*int32)(unsafe.Pointer(uintptr(p) + unsafe.Sizeof(int32(0)))))
	fmt.Println(*(*string)(unsafe.Pointer(uintptr(p) + unsafe.Sizeof(int32(0))*2)))
	fmt.Println(*(*bool)(unsafe.Pointer(uintptr(p) + unsafe.Sizeof(int32(0))*3)))
}
