# ルビ振りAPIサンプル
* Yahoo デベロッパーネットワークで掲載されているルビ振り API ( [ルビ振り（V2） - Yahoo!デベロッパーネットワーク](https://developer.yahoo.co.jp/webapi/jlp/furigana/v2/furigana.html) ) を使用したサンプルアプリケーション

## 開発環境
* VSCode
* Golang 1.17.1

## Go 言語のダウンロード
* [Downloads - The Go Programming Language](https://golang.org/dl/)

## VSCode 拡張パックのインストール
* Go - Go Team at Google をインストールする。
    - 拡張パックに関する情報 [https://marketplace.visualstudio.com/items?itemName=golang.Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)

## 事前準備
* GoLang 1.17.1 をダウンロードして `D:\\GoLang\\go1.17.1` ディレクトリに解凍する

## ビルド方法
* Shift + Ctrl + B でビルド。
* `./build` 配下に `app.exe` バイナリファイルが生成される
* Shift + Ctrl + P でコマンドパレットを表示、タスクの実行を選択
* `02. Copy text files` を実行する

## 使い方
* APIキーの設定
    - `app.exe` と同じ階層に `key.txt` ファイルを作成し、そこに API キー文字列を入力・保存する。
* 解析対象文字列の設定
    - `app.exe` と同じ階層に `list.txt` ファイルを作成し、そこにフリガナをつけたい文字列を入力する。
    - 解析は 1 行毎に実行される。
* デバックログの出力
    - `app.exe` と同じ階層に `debug.log` ファイルを作成するとアプリケーションログが出力される。
