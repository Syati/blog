+++
Categories = ["Go"]
Description = "Go 1.22 で追加された net/http の structured_params（Structured Field Values）を、使いどころと一緒に紹介します。"
Tags = ["Go", "net/http", "HTTP"]
comments = true
date = "2026-05-03T08:07:07+00:00"
title = "Go の structured_params 入門（HTTP Structured Field Values）"

+++

Go 1.22 で `net/http` に **structured field values**（RFC 8941）のパーサ/シリアライザとして `structured_params` が追加されました。

HTTPヘッダは歴史的に「カンマ区切り」「セミコロンでパラメータ」など、ヘッダごとに微妙に違う“手書きパース”になりがちでした。Structured Field Values は、そのヘッダ表現を共通化し、**安全にパースできる機械可読な形式**を提供します。

この記事では `structured_params` が何を解決するのか、どんな型があるのか、`net/http` の中でどう使うのかをまとめます。

<!--more-->

## structured field values とは

RFC 8941 で定義された、HTTPヘッダのためのデータ表現です。

- **Item**: 1つの値（文字列/数値/真偽/トークン/バイト列/日付など）+ パラメータ
- **List**: Item の配列
- **Dictionary**: `key=value` の集合

たとえば `Prefer` や `Signature` のように「キーと値」「値にぶら下がるパラメータ」を扱いたいヘッダでは、Structured Field Values を使うとパースが明確になります。

## `net/http` の `structured_params` パッケージ

`structured_params` は、上の Item / List / Dictionary をGoの型で扱うためのパッケージです。

利用する側のメリットはだいたい次の3つです。

1. **手書きパースを減らせる**（split/trim の積み上げをやめられる）
2. **エッジケースに強い**（引用符、エスケープ、空白、重複など）
3. **シリアライズも一貫**（自前で組み立てなくてよい）

## まずは最小例：List をパースする

ここでは「カンマ区切りの値 + パラメータ」を想定して List を読んでみます。

```go
package main

import (
  "fmt"
  "net/http/structured_params"
)

func main() {
  // 例: Item（token）にパラメータが付く
  // "foo" と "bar" の2要素、foo には ;a=1 が付く
  s := "foo;a=1, bar"

  list, err := structured_params.ParseList(s)
  if err != nil {
    panic(err)
  }

  for i, member := range list {
    fmt.Printf("%d: %v\n", i, member)
  }
}
```

ポイントは、`ParseList` の戻りが「単なる `[]string` ではない」ことです。各要素が Item（と、そのパラメータ）として保持されます。

## Item の中身：Bare Item と Parameters

Structured Field Values の Item は「値本体（Bare Item）」と「パラメータ」を持ちます。

- 値本体: token / string / number / boolean / byte sequence / date など
- パラメータ: `;key=value` という形でぶら下がるメタデータ

たとえば `text/html;q=0.9` の `q` は、よくある “品質値” パラメータです。structured field values ではこれも機械的に表現できます。

## Dictionary：`key=value` を扱う

`Dictionary` は `key=value` の集合を扱います（value が省略されたり、true扱いになったりするケースも含む）。

```go
package main

import (
  "fmt"
  "net/http/structured_params"
)

func main() {
  s := "a=1, b=?0, c;foo=\"bar\""

  dict, err := structured_params.ParseDictionary(s)
  if err != nil {
    panic(err)
  }

  fmt.Println(dict)
}
```

ヘッダをアプリケーションで拡張したいとき、Dictionary の形はとても扱いやすいです（後方互換のためにキー追加していく、など）。

## いつ使う？（使いどころの考え方）

`structured_params` は「どんなヘッダでもパースできる魔法の道具」ではなく、**Structured Field Values を採用しているヘッダ**や、**自分で設計する拡張ヘッダ**で真価を発揮します。

使いどころの目安:

- 既存ヘッダが RFC 8941 を参照している（あるいは将来対応したい）
- 値が `key=value` 的で、後から項目追加したい
- “値 + パラメータ” をちゃんと扱いたい（引用符やエスケープ込みで）

逆に、伝統的なフォーマット（例: `Cookie` のような独自構文）を無理に structured field values として扱うのは避けた方がよいです。

## 出力（シリアライズ）も揃える

structured field values は **送信側**でも便利です。自前の文字列連結をやめて、型から正しいフォーマットへ変換できます。

```go
// 例: List/Dictionary を組み立てて Serialize* 系で文字列にする
// （詳細は Go のドキュメントを参照）
```

これにより、微妙な空白・引用符・エスケープを毎回気にしなくて済みます。

## まとめ

- `structured_params` は HTTP Structured Field Values（RFC 8941）を扱うための `net/http` 標準パッケージ
- 手書きパースをやめて、Item / List / Dictionary として安全に扱える
- 「RFC 8941 準拠のヘッダ」や「拡張ヘッダ設計」で特に有用

次は、実際に自作ヘッダ（Dictionary）を導入して、`http.Handler` 側で `ParseDictionary` し、レスポンスで `SerializeDictionary` する例まで書くと運用イメージが湧きやすいです。

