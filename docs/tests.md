# Unitテスト＆E2Eテスト

## 概要

APIの品質担保のためにテストコードを実装する。

テスト観点
- 正常系
- 準正常系
- 異常系

準正常系はマストではない。


Unitテスト実行
```sh
make unit_test
```

E2Eテスト実行
```sh
make e2e_test:
```


## 実装

.env.testに依存するため、必ずhelperの読み込みをすること

```go
ctx, _ := helpers.Initialize(envPath)
```


## モック

```sh
mockgen -source application/port/user_port.go -destination test/units/mock/application/port/mock_user_port.go
```