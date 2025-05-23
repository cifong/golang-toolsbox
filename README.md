# 🧰 golang-toolsbox

一個使用 Golang 開發的系統工具箱，具備下列功能：

- 📊 顯示系統資訊（CPU、記憶體、磁碟使用率）
- ⏻ 遠端關機功能（透過 API 觸發）
- ✅ 四象限 ToDoList 工具（資料儲存在瀏覽器端 LocalStorage）

> 架構遵循 [golang-standards/project-layout](https://github.com/golang-standards/project-layout)，具備可測試性與模組化。

---

## 專案架構（project-layout）  

golang-toolsbox/  
├── cmd/ # 主要執行點 (server)  
├── internal/  
│ ├── handler/ # HTTP handler (page/system)
│ ├── router/ # Gin 路由設定  
│ ├── system/ # 系統資訊、關機邏輯  
├── web/  
│ ├── static/ # 存放css js   
│ └── templates/ # html模板  
├── go.mod  
├── README.md  

## 環境
- Go 1.24.3
- Git
- VSCode

---

## 🚀 執行方式

### 1️⃣ 安裝依賴
- go mod tidy
### 2️⃣ 執行伺服器
- go run cmd/server/main.go

### 伺服器啟動後：

- 系統資訊： http://localhost:8080/api/v1/info

- 關機： POST http://localhost:8080/api/v1/shutdown

- ToDo 工具頁面： http://localhost:8080/todo/