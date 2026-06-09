+++
Categories = ["Zsh"]
Description = "fzf を使って Entire CLI のセッションやチェックポイントをインタラクティブに操作できる zsh プラグイン「entire-fzf」を作りました。"
Tags = ["Zsh", "fzf", "entire", "CLI"]
comments = true
date = "2026-06-10T07:00:00+09:00"
title = "fzf で Entire セッションを選択・操作する zsh プラグイン entire-fzf を作った"
logo = "shell"
+++

Claude Code などの AI エージェントのセッション履歴やチェックポイントを Git ライクに記録・管理できる CLI ツール **Entire**。
非常に便利なのですが、セッションやチェックポイントの切り替え、レジューム、各チェックポイントの AI による変更理由の解説などをコマンドラインから ID を指定して実行するのは少々面倒です。

そこで、fzf を使ってインタラクティブに Entire セッションやチェックポイントを選択・操作できる Zsh プラグイン **entire-fzf** を作成しました。

https://github.com/Syati/entire-fzf

<!--more-->

![Demo](https://github.com/Syati/entire-fzf/raw/main/demo/demo.gif)

## 必要なもの

このプラグインを利用するには、以下のコマンドがあらかじめインストールされている必要があります。

* `zsh`
* `entire` (AI セッション記録・管理 CLI)
* `fzf` (コマンドラインファジーファインダー)
* `jq` (JSON パーサー)
* `git`

## インストール方法

### プラグインマネージャーを使わない場合

リポジトリをクローンして、`.zshrc` でプラグインスクリプトを読み込みます。

~~~bash
$ git clone https://github.com/Syati/entire-fzf.git ~/.zsh/entire-fzf
~~~

`.zshrc` に以下を追記します：

~~~bash
# ~/.zshrc
source ~/.zsh/entire-fzf/entire-fzf.plugin.zsh
~~~

### Sheldon を使う場合

[Sheldon](https://github.com/rossmacarthur/sheldon) をお使いの場合は、`~/.config/sheldon/plugins.toml` に以下を追加します：

~~~toml
[plugins.entire-fzf]
github = "Syati/entire-fzf"
~~~

Sheldon はプラグイン名 `entire-fzf` から自動的に `entire-fzf.plugin.zsh` をマッチングしてロードしてくれます。

設定を反映するには、シェルを再起動するか以下を実行します：

~~~bash
$ eval "$(sheldon source)"
~~~

## 使い方（提供されるコマンド）

プラグインをロードすると、以下の 2 つの便利なインタラクティブコマンドが使えるようになります。

### 1. `etf` (Entire Session Picker)

`etf` を実行すると、記録されている Entire セッション一覧が fzf でインタラクティブに表示されます。
セッションを選択すると、そのセッションに対して以下のいずれかのアクションを選択して実行できます：

* **resume**: セッションを再開する
* **explain latest**: 最新のチェックポイントの説明（AI による変更の意図などの解説）を表示する
* **pick checkpoint & explain**: チェックポイント一覧から個別に選択して、その説明を表示する
* **info**: セッションの詳細情報を表示する
* **stop**: アクティブなセッションを停止する
* **clean**: セッション履歴をクリーンアップする

### 2. `etfc` (Active Session Checkpoint Picker)

`etfc` を実行すると、カレントのワークツリーでアクティブになっている Entire セッションのチェックポイント（AI がコミット・チェックポイントを作成した履歴）一覧が fzf で表示されます。
チェックポイントを選択するだけで、即座にそのチェックポイントに対する AI 生成の解説（explain）を確認できます。

## 開発とCI

プラグインスクリプトのシンタックスチェックは以下で行うことができます：

~~~bash
$ zsh -n entire-fzf.plugin.zsh
~~~

Entire での開発効率が圧倒的に向上するので、ぜひ使ってみてください！
