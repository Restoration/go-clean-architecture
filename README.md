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

[ca](https://github.com/Restoration/go-clean-architecture/blob/main/docs/images/ca.jpg)

クリーンアーキテクチャにおいて肝心な図はこちら。
ドーナツ状になっている図はあくまで抽象的な部分を示しており、さらに細分化した図はこちらになる。

[ca2](https://github.com/Restoration/go-clean-architecture/blob/main/docs/images/ca-2.jpg)

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

## 各種レイヤーとしての役割

改めて戻るがドーナッツ状になっているこちらの図は大きなレイヤーとして意味になる。
Presentation, Application, Infrastructure, Entity となる。

基本的にレイヤーは抽象と依存するように実装を行う。
そのため、レイヤーは基本的に抽象と実態の両方を持つようにし、レイヤーとの境界には抽象を挟むイメージ。

レイヤーには実態を DI(Dependency Injection)を繰り返していき、実態を入れていくイメージである。

### Presentation

プレゼンテーション層と呼ばれるこのレイヤーは表示および、外部からの操作と外部への出力を示すレイヤーとなる。
ユーザー、クライアントとのやり取りに使用される。

##### Controller

外部からのリクエストを受け取る。
主にバリデーションや外部からの値をアプリケーションが扱いやすいように整形する。
ドメイン変換もここで行う。

##### Response

Viewの役割を持つ。
レスポンスとして整形し、値を返す場所。
外部とやり取りする際はここが最終的な値になる。

### Application

アプリケーション層と呼ばれるこのレイヤーはアプリケーション本体を指し示す。
コアとなるロジックはこちらに記載をしていく。

##### UseCase

ビジネスロジックを扱うInteractorの抽象。

##### Interactor

UseCase で定義した実態。

アプリケーションロジックを担うサービス層でコアの部分。
DI されてきた Repository やGateWayなどはここで扱う。

##### Port

Port は Input Boundary と Output Boundary の役割を持っており、interface として定義している。
GateWay や Repository の抽象。
アプリケーションロジックに対する

### Infrastructure

インフラ層は外部との連携。
クリーンアーキテクチャにおいて外部I/Oやフレームワーク依存部分を担う層。
外部サービス、システム外のリソースなど、アプリケーション外部への接続やフレームワーク依存部分を担当する。

プレゼンテーションとインフラストラクチャーの明確な違いはプレゼンテーションはユーザーとのやりとりを行い、インフラストラクチャーは外部リソースとのやりとりを行う。

##### Gateway

Port で定義した抽象の実態。

Gateway は外部システムへの接続を担う役割を持っている。
外部 API や外部サービスへのリクエスト/レスポンス処理を行う。
外部サービスがビジネスロジックに直接影響を与えないようにする。

##### Repository

Port で定義した抽象の実態。

Repository は内部データ管理をする。
データベースへのクエリ、CRUD 操作などはこの Repository 内で行う。
内部データ（データベースやストレージ）とのやり取りを抽象化する。
アプリケーションが永続化層の具体的な実装に依存しないようにする。

クリーンアーキテクチャでは、内側（Domain／UseCase）から外側（Infra, Presentation）へは依存しないように設計する。

#### Driver

クライアント系のモジュール。
データベースへの接続やHTTPクライアントなど、外部とやり取りするための処理はここに格納。


#### db

マイグレーションやシードなどのSQLなどはこちらに格納。

#### DAO

Domain Access Objectの略。ORMのマッピングはこちらで定義する。

#### Router

コントローラーを呼び出し、APIのルーティングを定義。


### Entity

一意の識別子を持ち、ライフサイクルを通じて同一性（同じオブジェクトであること）を保証するオブジェクトを指します。

「ユーザー」や「注文」など、実世界の概念や業務上の概念をプログラム内で表現したものです。
データが変化しても、そのオブジェクトが同一であること（同一性）をIDで保証します。

DDDでは単なるデータの集まりではなく、そのドメイン（業務や問題領域）固有のルールや振る舞い（メソッド）を持ちます。
（例：ユーザー名の変更時に何かしらのバリデーションを行う、支払い完了時に在庫を減らすなど。）

アプリケーションの中心であるドメイン層（もっとも内側の層）に位置する

DAOとの違い。
データベースのテーブル定義やORMの具体的なアノテーションなどは、できるだけドメイン層から切り離しておくのが望ましいとされる。
それはドメインがインフラ技術に引きずられないように防ぐためです。

#### Domain

ビジネスロジックを定義したドメイン。
DAOからドメインへと変換するロジックもこちらで持つ。

## まとめ

クリーンアーキテクチャで実装されたプロジェクトは何度も見てきたが、会社によってはレイヤーを減らして工数削減を狙ったりするプロジェクトもあった。
ここまでいったようにレイヤーが細分化されていることは、ファイル数が多くなる。
ソースを追うのが大変になるケースも無きにしも非ずなので、必ずしもクリーンアーキテクチャを完璧に使うというのが正攻法とは限らない。

## Author

[RyotArch](https://www.developer-ryota.com/)

## License

MIT License
