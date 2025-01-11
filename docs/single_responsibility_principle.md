# 単一責任原則

[Wikipedia](https://ja.wikipedia.org/wiki/%E5%8D%98%E4%B8%80%E8%B2%AC%E4%BB%BB%E3%81%AE%E5%8E%9F%E5%89%87)

> 単一責任の原則 (たんいつせきにんのげんそく、英: single-responsibility principle) は、プログラミングに関する原則であり、モジュール、クラスまたは関数は、単一の機能について責任を持ち、その機能をカプセル化するべきであるという原則である。
> モジュール、クラスまたは関数が提供するサービスは、その責任と一致している必要がある。

## ポイント

## 責任を明確化

1 つの構造体や関数が、複数の責任を持たないように設計します。
例えば、「データベースの操作」と「HTTP レスポンスの生成」を 1 つの構造体で行わない。

## 変更理由を 1 つに限定

あるコードを変更する理由が複数ある場合、そのコードは単一責任原則に反しています。

## コードの分離

関数やモジュールを小さな責任ごとに分け、他の部分に依存しすぎないようにします。

NG

```go
type UserHandler struct {
    db *sql.DB
}

// データベースの操作とHTTPレスポンスの生成を同じ関数で行っている
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    // データベースにユーザーを保存
    _, err := h.db.Exec("INSERT INTO users (name) VALUES ('John')")
    if err != nil {
        http.Error(w, "Failed to save user", http.StatusInternalServerError)
        return
    }

    // HTTPレスポンスを生成
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User created"))
}
```

データベースの操作と HTTP レスポンスの生成を同じ関数で行っている
これはデータベース操作と HTTP レスポンス生成の責任を分離させるべきである。

Good

```go
// データベースの操作を担当する構造体
type UserRepository struct {
    db *sql.DB
}

func (repo *UserRepository) CreateUser(name string) error {
    _, err := repo.db.Exec("INSERT INTO users (name) VALUES (?)", name)
    return err
}

// HTTPリクエストの処理を担当する構造体
type UserHandler struct {
    repo *UserRepository
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        http.Error(w, "Name is required", http.StatusBadRequest)
        return
    }

    // UserRepositoryに処理を委譲
    err := h.repo.CreateUser(name)
    if err != nil {
        http.Error(w, "Failed to save user", http.StatusInternalServerError)
        return
    }

    // HTTPレスポンスを生成
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User created"))
}
```

- UserRepository:データベース操作の責任を持つ。
- UserHandler:HTTP リクエストの処理とレスポンス生成の責任を持つ。

分離させたことによりそれぞれが 1 つの責任に集中しており、変更の影響が限定的になります。
