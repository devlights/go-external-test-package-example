# go-external-test-package-example
This is example using golang external test package.

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/devlights/go-external-test-package-example)

## internal and external tests

The following tests exist in the hello package (internal test package):

- pkg/hello/hello_test.go
  - TestMakeMessage

The following tests exist in the hello_test package (external test package):

- pkg/hello/hello_external_test.go
  - TestSay

## run

```sh
$ go test -v -cover ./...
=== RUN   TestMakeMessage
=== RUN   TestMakeMessage/not_empty
=== RUN   TestMakeMessage/empty
--- PASS: TestMakeMessage (0.00s)
    --- PASS: TestMakeMessage/not_empty (0.00s)
    --- PASS: TestMakeMessage/empty (0.00s)
=== RUN   TestSay
=== RUN   TestSay/not_empty
=== RUN   TestSay/empty
--- PASS: TestSay (0.00s)
    --- PASS: TestSay/not_empty (0.00s)
    --- PASS: TestSay/empty (0.00s)
=== RUN   ExampleSay
--- PASS: ExampleSay (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/devlights/go-external-test-package-example/pkg/hello 0.004s  coverage: 100.0% of statements
```

## doc

```sh
$ go doc -all ./pkg/hello/
package hello // import "github.com/devlights/go-external-test-package-example/pkg/hello"

Package hello contains functions to make greeing message.

FUNCTIONS

func Say(message string) string
    Say makes the greeting message.

    if parameter is empty, return empty.

```

## references

- 書籍「プログラミング言語Go」第11章 テスト
  - ここに外部テストパッケージの件も含めて記載されています
- [外部テストパッケージの利用ケース](https://qiita.com/hogedigo/items/5f491994647aa4a8a905)
- [5 advanced testing techniques in Go](https://segment.com/blog/5-advanced-testing-techniques-in-go/)
- [Go Fridayこぼれ話：非公開（unexported）な機能を使ったテスト](https://engineering.mercari.com/blog/entry/2018-08-08-080000/)
