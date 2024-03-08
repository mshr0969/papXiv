# papXiv
Paper Management System based on arXiv

## feature
- arXiv APIを用いて論文を検索できる(TODO)
- 論文の情報のCRUD操作

## Preparation
```bash
cd path/to/papXiv/gateway
```

### Start the Database (terminal 1)
Dockerを使用してMySQLを起動する
```bash
make up
```

### run Go API server (terminal 2)
```bash
go run internal/cmd/main.go
```

## Example
### 論文を登録
```bash
curl -X POST -H "Content-Type: application/json" -d '{"subject":"physics", "url":"https://arxiv.org/hogehoge", "published":"2024/03/01", "id":"2d1423b3-15ff-498e-b978-241f2b87de9e", "title": "test title"}' localhost:8080/papers
```

### 論文情報の詳細取得
```bash
curl localhost:8080/paper/2d1423b3-15ff-498e-b978-241f2b87de9e
```
Response example:
```
{"created_at":"2024-03-08T23:35:02+09:00","id":"2d1423b3-15ff-498e-b978-241f2b87de9e","published":"2024/03/01","subject":"physics","title":"test title","updated_at":"2024-03-08T23:35:02+09:00","url":"https://arxiv.org/hogehoge"}%
```

### 全ての論文を取得
```bash
curl localhost:8080/papers
```
Response example:
```
{"papers":[{"id":"2d1423b3-15ff-498e-b978-241f2b87de9e","title":"test title"}],"total":1}
```

### 論文情報の更新
subject, url, published, titleは全て任意のため、更新したいもののみリクエストする
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"subject":"math", "url":"https://arxiv.org/hogehoge", "published":"2024/03/07", "title": "updated title"}' localhost:8080/paper/2d1423b3-15ff-498e-b978-241f2b87de9e
```

### 論文情報の削除
```bash
curl -X DELETE localhost:8080/paper/2d1423b3-15ff-498e-b978-241f2b87de9e
```
