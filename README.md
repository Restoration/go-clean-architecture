# go-clean-architecture

## 概要

Go 言語によるRESTFul APIの実装。
使用技術は以下

- DDD(クリーンアーキテクチャ)
- 分散トレーシング
- ユニットテスト


## 立ち上げ方法

事前に Docker をインストールしください。

```
$ make init
$ docker-compose -d
$ make maigrate_up
$ make seed
```

以下でレスポンス確認

```
http://localhost:8080/users
```

モニタリング（jaeger）
```
http://localhost:16686
```

## ドキュメント
- [クリーンアーキテクチャ]()
- [分散トレーシング]()

## Author

[RyotArch](https://www.developer-ryota.com/)

## License

MIT License
