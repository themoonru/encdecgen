encdecgen
=========

Пакет для генерации кодеров-декодеров msgPack на основе ast дерева.

Пример использования
====================
``` 
// Построить ast дерево
parser := ast_parser.New()
data, err := parser.ParseFile("source.go")
if err != nil {
    fmt.Println(err)
    os.Exit(1)
}

// Сгенерировать кодеры-декодеры на основе одного из шаблонов
gen := generator.New(template.MsgpackEncDecTmpl)
res, err := gen.Generate(data)
if err != nil {
    fmt.Println(err)
    os.Exit(1)
}

fmt.Println(res)
```

Contributing
============
1. Склонируйте репозиторий 
2. Создайте новую ветку и внесите изменения
3. Напишите тесты
4. Перед тем как сделать push, запустите тесты ```go test ./...``` и убедитесь, что они проходят
5. Подайте pull-request