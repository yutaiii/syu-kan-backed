# syu-kan backend
## migration
### local
プロジェクトのルートで以下を実行

#### version
```
migrate -source file://app/backend/tool/migrate -database 'mysql://root:root@tcp(127.0.0.1:3306)/syu_kan' version
```

#### up
```
migrate -source file://app/backend/tool/migrate -database 'mysql://root:root@tcp(127.0.0.1:3306)/syu_kan' up
```

#### down
```
migrate -source file://app/backend/tool/migrate -database 'mysql://root:root@tcp(127.0.0.1:3306)/syu_kan' down
```