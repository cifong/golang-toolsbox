// 載入系統資訊
fetch('/api/v1/system/info')
  .then(res => res.json())
  .then(data => {
    const infoDiv = document.getElementById('system-info');
    infoDiv.innerHTML = `
      <p>作業系統：${data.os}</p>
      <p>架構：${data.arch}</p>
      <p>Go 版本：${data.version}</p>
      <p>CPU Usage：${data.cpu_usage}</p>
      <p>Used Memory：${data.used_memory}</p>
      <p>Total Memory：${data.total_memory}</p>
    `;
  });

// 關機按鈕事件
document.getElementById('shutdown-button').addEventListener('click', () => {
  if (confirm("你確定要關機嗎？")) {
    fetch('/api/v1/system/shutdown', {
      method: 'POST'
    })
    .then(res => res.json())
    .then(data => {
      alert(data.message || data.error);
    });
  }
});
