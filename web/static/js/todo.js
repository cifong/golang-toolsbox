const quadrants = [
  { id: 'q1', label: '重要且緊急' },
  { id: 'q2', label: '重要不緊急' },
  { id: 'q3', label: '不重要但緊急' },
  { id: 'q4', label: '不重要不緊急' }
];

function loadTasks() {
  const data = localStorage.getItem('todo-tasks');
  return data ? JSON.parse(data) : { q1: [], q2: [], q3: [], q4: [] };
}

function saveTasks(tasks) {
  localStorage.setItem('todo-tasks', JSON.stringify(tasks));
}

function renderTasks(tasks) {
  quadrants.forEach(({ id }) => {
    const ul = document.querySelector(`#${id} ul`);
    ul.innerHTML = '';
    tasks[id].forEach((task, index) => {
      const li = document.createElement('li');
      li.textContent = task;
      li.className = 'px-3 py-2 bg-blue-100 rounded hover:bg-blue-200 cursor-pointer transition';
      li.dataset.index = index;
      li.dataset.quadrant = id;
      ul.appendChild(li);
    });
  });
}

function createQuadrantSection({ id, label }) {
  const div = document.createElement('div');
  div.className = 'bg-white p-4 rounded shadow';
  div.id = id;

  const h3 = document.createElement('h3');
  h3.className = 'text-xl font-semibold mb-2';
  h3.textContent = label;

  const ul = document.createElement('ul');
  ul.className = 'space-y-2';

  const form = document.createElement('form');
  form.className = 'mt-4 flex';
  form.dataset.quadrant = id;

  const input = document.createElement('input');
  input.type = 'text';
  input.placeholder = '新增事項';
  input.className = 'flex-1 border border-gray-300 rounded-l px-3 py-2 focus:outline-none';

  const btn = document.createElement('button');
  btn.type = 'submit';
  btn.textContent = '新增';
  btn.className = 'bg-blue-600 text-white px-4 py-2 rounded-r hover:bg-blue-700 transition';

  form.appendChild(input);
  form.appendChild(btn);

  div.appendChild(h3);
  div.appendChild(ul);
  div.appendChild(form);

  return div;
}

// 主初始化與事件統一處理
function initTodoApp() {
  const app = document.getElementById('todo-app');
  const tasks = loadTasks();

  // 渲染 UI 結構
  quadrants.forEach(q => {
    app.appendChild(createQuadrantSection(q));
  });

  // 渲染初始任務
  renderTasks(tasks);

  // 表單事件統一註冊
  app.addEventListener('submit', (e) => {
    e.preventDefault();
    const form = e.target.closest('form');
    const input = form.querySelector('input');
    const quadrant = form.dataset.quadrant;
    const task = input.value.trim();

    if (task) {
      tasks[quadrant].push(task);
      saveTasks(tasks);
      renderTasks(tasks);
      input.value = '';
    }
  });

  // 點擊任務事件統一處理
  app.addEventListener('click', (e) => {
    if (e.target.tagName === 'LI') {
      const index = e.target.dataset.index;
      const quadrant = e.target.dataset.quadrant;
      if (confirm('刪除此項目？')) {
        tasks[quadrant].splice(index, 1);
        saveTasks(tasks);
        renderTasks(tasks);
      }
    }
  });
}

document.addEventListener('DOMContentLoaded', initTodoApp);
