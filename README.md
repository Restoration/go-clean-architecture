# go-clean-architecture

## 概要

Go 言語によるクリーンアーキテクチャを実装。クリーンアーキテクチャの概念についてを記載。

## 立ち上げ方法

事前に Docker をインストールしください。

```
$ make init
$ docker-compose -d
$ make maigrate_u
$ make seed
```

## クリーンアーキテクチャ

クリーンアーキテクチャにおいて肝心な図はこちら。
ドーナツ状になっている図はあくまで抽象的な部分を示しており、さらに細分化した図はこちらになる。

この図の中で出てくる`<I>`とは interface を指し示しており、`<DS>`とは DataStructure を意味している。

以降の説明では interface のことを抽象として扱い、DataStructure のことを実態として説明する。

### SOLID 原則

前提として、クリーンアーキテクチャを実装する上でまず知っておかなければならない原則がある。
それが SOLID 原則である。SOLID 原則は以下の５つの原則の頭文字を取ったものになる。

- 単一責任の原則 (single-responsibility principle)
- 開放閉鎖の原則（open/closed principle）
- リスコフの置換原則（Liskov substitution principle）
- インターフェース分離の原則 (interface segregation principle)
- 依存性逆転の原則（dependency inversion principle）

SOLID 原則をアプリケーションのアーキテクトに反映させたものがクリーンアーキテクチャになる。

### 各種レイヤーとしての役割

改めて戻るがドーナッツ状になっているこちらの図は大きなレイヤーとして意味になる。
Presentation, Application, Infrastructure, Entity となる。

基本的にレイヤーは抽象と依存するように実装を行う。
そのため、レイヤーは基本的に抽象と実態の両方を持つようにし、レイヤーとの境界には抽象を挟むイメージ。

レイヤーには実態を DI(Dependency Injection)を繰り返していき、実態を入れていくイメージである。

#### Presentation

プレゼンテーション層と呼ばれるこのレイヤーは表示および、外部からの操作と外部への出力を示すレイヤーとなる。
家庭用ゲーム機でいうところのコントローラーとテレビ画面にあたる。

##### Controller

##### Response

#### Application

アプリケーション層と呼ばれるこのレイヤーはアプリケーション本体を指し示す。
コアとなるロジックはこちらに記載をしていく。

##### UseCase

##### Port

Port は Input Boundary と Output Boundary の役割を持っており、interface として定義している。

##### Gateway

Gateway は外部システムへの接続を担う役割を持っている。
外部 API や外部サービスへのリクエスト/レスポンス処理を行う。
外部サービスがビジネスロジックに直接影響を与えないようにする。

##### Repository

Repository は内部データ管理をする。
データベースへのクエリ、CRUD 操作などはこの Repository 内で行う。
内部データ（データベースやストレージ）とのやり取りを抽象化する。
アプリケーションが永続化層の具体的な実装に依存しないようにする。

## まとめ

クリーンアーキテクチャで実装されたプロジェクトは何度も見てきたが、会社によってはレイヤーを減らして工数削減を狙ったりするプロジェクトもあった。
ここまでいったようにレイヤーが細分化されていることは、ファイル数が多くなる。
ソースを追うのが大変になるケースも無きにしも非ずなので、必ずしもクリーンアーキテクチャを完璧に使うというのが正攻法とは限らない。

## Author

[RyotArch](https://www.developer-ryota.com/)

## License

MIT License
