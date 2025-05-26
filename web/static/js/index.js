document.addEventListener('click', function (e) {
  const action = e.target.getAttribute('data-action');
  if (!action) return;

  switch (action) {
    case 'get-system-info':
      fetch('/api/v1/system/info')
        .then(res => res.json())
        .then(data => {
          document.getElementById('sysinfo').innerHTML = `
            <p>作業系統：${data.os}</p>
            <p>架構：${data.arch}</p>
            <p>Go 版本：${data.version}</p>
            <p>CPU Usage：${roundToTwo(data.cpu_usage)}%</p>
            <p>Used Memory：${bytesToGB(data.used_memory)}GB</p>
            <p>Total Memory：${bytesToGB(data.total_memory)}GB</p>
          `;
        })
        .catch(err => {
          document.getElementById('sysinfo').textContent = '取得失敗：' + err.message;
        });
      break;

    case 'shutdown':
      if (confirm("確定要關機嗎？這會立即執行系統關機！")) {
        fetch('/api/v1/system/shutdown', { method: 'POST' })
          .then(res => res.json())
          .then(data => alert(data.message))
          .catch(err => alert("關機失敗：" + err.message));
      }
      break;

    case 'goto-todo':
      window.location.href = "/todo";
      break;
  }
});

let socket = null;

function connectWebSocket() {
  socket = new WebSocket("ws://localhost:8080/ws/v1/system/info");

  socket.onopen = () => {
    console.log("✅ WebSocket 已連線");
  };

  socket.onmessage = (event) => {
    const data = JSON.parse(event.data);
    document.getElementById('sysinfo').innerHTML = `
      <p>作業系統：${data.os}</p>
      <p>架構：${data.arch}</p>
      <p>Go 版本：${data.version}</p>
      <p>CPU Usage：${roundToTwo(data.cpu_usage)}%</p>
      <p>Used Memory：${bytesToGB(data.used_memory)}GB</p>
      <p>Total Memory：${bytesToGB(data.total_memory)}GB</p>
    `;
  };

  socket.onerror = (err) => {
    console.error("❌ WebSocket 錯誤", err);
  };

  socket.onclose = () => {
    console.log("🔌 WebSocket 已關閉");
  };
}

function disconnectWebSocket() {
  if (socket) {
    socket.close();
    socket = null;
  }
}

document.getElementById('realtimeToggle').addEventListener('change', function () {
  if (this.checked) {
    connectWebSocket();
  } else {
    disconnectWebSocket();
  }
});

function roundToTwo(num) {
  return Number((num).toFixed(2));
}
function bytesToGB(bytes) {
  return Number((bytes / (1024 ** 3)).toFixed(2));
}