# 簡易的なTodoアプリ
- [ ] 基本的にはMVCアーキテクチャ
- [ ] modelに関するvalidationテストを実装
- [ ] ログイン情報をsessionで管理することでtodoのSQL実行前に確認

## 工夫した点
- [ ] ログイン状態のセッション情報の管理をAuth0等のIDaasを使わずに１から作成
  - 次はauth0使いたい
  https://auth0.com/signup?place=header&type=button&text=sign%20up

## 実装上困っている点
- [ ] model層とservice層の使い分け
  - [ ] model層
    - データモデルの定義と属性の管理
    - データアクセス(CRUD処理、外部APIとのやり取り)
  - [ ] service層
    - ビジネスロジックの実装
      - [ ] データのバリデーション
      - [ ] データの変換
      - [ ] 外部APIへのアクセス
      - [ ] データの処理
    - ドメインルールの実装
    - コントローラとのインターフェース

- [ ] testを書くとき、DBモックの作成の仕方がよくわからなかった
  - テスト用にDB接続を立ててもいいけどなぁ、サーバー負荷、CI回す時遅くなりそう....
  - やり方調べよう
  - [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)とか使えそう

## 今後取り組みたいこと
- [ ] クリーンアーキテクチャで作成してみたい
  - https://zenn.dev/88888888_kota/articles/1a5caac8d743b8
- [ ] Gin触ってみたい
  ``` go
  package main

    import (
    "fmt"
    "net/http"
    )

    func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
    }

    func main() {
    http.HandleFunc("/", helloHandler)
    http.ListenAndServe(":8080", nil)
    }
  ```
  
  これが
  
  ``` gin
  package main

    import (
    "github.com/gin-gonic/gin"
    )

    func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
      c.String(http.StatusOK, "Hello, World!")
    })

    router.Run(":8080")
    }
  ```
  
  こう書けるらしい
  
  - [ ] フロントはNext.js使ってリッチにしていきたい
  - [ ] AWSにデプロイしていきたい
