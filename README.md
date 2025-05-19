# 專案架構（project-layout）  
   golang-toolsbox/
├── cmd/                  # 主執行點  
│   └── server/           # Web Server 入口  
│       └── main.go  
├── internal/             # 主要邏輯（封裝在 internal）  
│   ├── systeminfo/       # 系統資訊  
│   ├── shutdown/         # 關機邏輯  
│   └── todolist/         # ToDoList 模組  
├── api/                  # handler 與路由設定  
├── web/                  # HTML/JS 前端（或分離部署） 
├── test/                 # 測試  
├── scripts/              # 啟動與部署腳本  
├── go.mod  
└── README.md  

# 環境
- Go 1.24.3
- Git
- VSCode