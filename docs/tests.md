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

TODO：Unitテストにおいてデータの永続化処理のテストコードを実装   
E2EテストでカバーしているのでひとまずはOKとしている。

controllerで入出力のテストを実行  
interactorでアプリケーションのコアロジックのテストを実行  
E2Eで総合的なテストを実施するという流れで行なっている。


## モック

mockgenを利用してモックを作成

実行例
```sh
mockgen -source application/port/user_port.go -destination test/units/mock/application/port/mock_user_port.go

mockgen -source application/usecase/user_usecase.go -destination test/units/mock/application/usecase/mock_user_usecase.go
```