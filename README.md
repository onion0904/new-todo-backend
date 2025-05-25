## コードを作成する上で
### 気づいたところ
- 調べながら作るには「どう動かそうか」「どう構成するか」などが分かっていないとできないので言語化するのがとても難しかったです
- Statusをフックを利用して3パターンのみ入れられるように指定したかったけど、エラー解決が間に合わなかった。

### 工夫点
- 分割できる所は分割した。
- 以前作ったTodoListを基に大まかなところは変えずに作った。

### 技術の選定理由
- ginとgormは有名だからとりあえず使ってみたかった。
- MySQLは本([APIを作りながら進むGo 中級者への道](https://techbookfest.org/product/jXDAEU1dR53kbZkgtDm9zx?productVariantID=dvjtgpjw8VDTXNqKaanTVi&utm_source=pocket_saves))で学んだため用いた。

## 作成したアプリケーションを起動し動作確認する手順
### インストール
```
git clone https://github.com/catechlounge/onion0904-todo-api
```

### 起動方法

```
docker-compose up
go run main.go
```

### 動作確認
- 追加する
  - TitleにはTodo名を入れる
  - Statusには未着手、進行中、完了の三つを入れる
  - Priorityには数字が入れられる。小さいほど優先度が高くなる
```
curl -X POST "http://localhost:8080/Todo/add?Title=Todo1&Status=進行中&Priority=5"
```

- 一覧の取得
  - Titleの指定なしで一覧の取得
  - Titleの指定で部分一致するTodoの取得
```
curl -X GET "http://localhost:8080/Todo/list?Title=任意の文字列"
```

- 優先順位順で取得
```
curl -X GET "http://localhost:8080/Todo/list/sorted"
```

- Todoのアップデート
  - IDを指定して指定されたTodoのTitle、Status、Priorityをアップデート
```
curl -X PUT "http://localhost:8080/Todo/update?ID=11&Title=ああ&Status=完了&Priority=5"
```

- Todoの削除
  - 指定されたIDのTodoを削除
```
curl -X DELETE "http://localhost:8080/Todo/delete?id=1"
```