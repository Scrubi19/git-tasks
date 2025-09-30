package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

type StringStruct struct {
	data unsafe.Pointer
	Len  int
}

type SliceStruct struct {
	data unsafe.Pointer
	Size int
	Cap  int
}

func main() {
	s := "Hello, Мир"

	// Пример формирования строки из байтов
	bytes := []byte{72, 101, 108, 108, 111}
	s3 := string(bytes)
	fmt.Printf("Вывод строки из байтов: %s\n", s3) // "Hello"

	fmt.Printf("Длина в байтах: %d\n", len(s))                         // 13 байт
	fmt.Printf("Количество символов: %d\n", utf8.RuneCountInString(s)) // 10 символов

	// Преобразования
	runes := []rune(s)   // string → []rune
	str := string(runes) // []rune → string

	fmt.Printf("string в []rune: %d\n", len(runes))
	fmt.Printf("Обратно в string: %s\n", str)

	// Сравнение размера в памяти
	fmt.Printf("Размер string: %d байт\n", unsafe.Sizeof(s))     // 16 байт
	fmt.Printf("Размер rune: %d байт\n", unsafe.Sizeof('A'))     // 4 байта
	fmt.Printf("Размер []rune: %d байт\n", unsafe.Sizeof(runes)) // 24 байта (Потому что размер самой структуры 8 байт)
	fmt.Printf("Длина в байтах []rune: %d байт\n", len(runes)*4) // 40 байт
}
