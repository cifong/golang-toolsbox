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

function renderTasks() {
  const tasks = loadTasks();
  quadrants.forEach(({ id }) => {
    const ul = document.querySelector(`#${id} ul`);
    ul.innerHTML = '';
    tasks[id].forEach((task, index) => {
      const li = document.createElement('li');
      li.textContent = task;
      li.onclick = () => {
        if (confirm("刪除此項目？")) {
          tasks[id].splice(index, 1);
          saveTasks(tasks);
          renderTasks();
        }
      };
      ul.appendChild(li);
    });
  });
}

function createQuadrantSection({ id, label }) {
  const div = document.createElement('div');
  div.className = 'quadrant';
  div.id = id;

  const h3 = document.createElement('h3');
  h3.textContent = label;

  const ul = document.createElement('ul');

  const form = document.createElement('form');
  form.onsubmit = (e) => {
    e.preventDefault();
    const input = form.querySelector('input');
    const task = input.value.trim();
    if (task) {
      const tasks = loadTasks();
      tasks[id].push(task);
      saveTasks(tasks);
      input.value = '';
      renderTasks();
    }
  };

  const input = document.createElement('input');
  input.type = 'text';
  input.placeholder = '新增事項';

  form.appendChild(input);
  div.appendChild(h3);
  div.appendChild(ul);
  div.appendChild(form);

  return div;
}

function initTodoApp() {
  const app = document.getElementById('todo-app');
  app.innerHTML = '';
  quadrants.forEach(q => {
    app.appendChild(createQuadrantSection(q));
  });
  renderTasks();
}

document.addEventListener('DOMContentLoaded', initTodoApp);
