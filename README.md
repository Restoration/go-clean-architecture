# go-clean-architecture

## 概要

Go 言語によるRESTFul APIの実装。  
  
使用技術  
実装詳細はリンクを参照

- [クリーンアーキテクチャ](https://github.com/Restoration/go-clean-architecture/blob/main/docs/clean_architecture.md)
- [分散トレーシング](https://github.com/Restoration/go-clean-architecture/blob/main/docs/distributed_tracing.md)
- [Unitテスト＆E2Eテスト](https://github.com/Restoration/go-clean-architecture/blob/main/docs/tests.md)
- [データベースシャーディング](https://github.com/Restoration/go-clean-architecture/blob/main/docs/database_sharding.md)

TODO

- 負荷テスト

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
http://localhost:8080/v1/users
```

モニタリング（jaeger）
```
http://localhost:16686
```

## Author

[RyotArch](https://www.developer-ryota.com/)

## License

MIT License
