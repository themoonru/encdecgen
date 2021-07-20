package test

import "code.citik.ru/gobase/msgpackgen/internal/test/other_package_for_test"

//использование типов из других пакетов
type OtherPackageScalarTypes struct {
	Age  other_package_for_test.Age  `msgpack_index:"1"`
	Name other_package_for_test.Name `msgpack_index:"2"`
	Flag other_package_for_test.Flag `msgpack_index:"3"`
}

type OtherPackagePersonStruct struct {
	Person other_package_for_test.Person `msgpack_index:"1"`
}

type OtherPackageCarStruct struct {
	Auto other_package_for_test.Car `msgpack_index:"1"`
}