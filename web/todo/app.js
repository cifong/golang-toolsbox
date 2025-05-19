document.addEventListener('DOMContentLoaded', () => {
  const form = document.getElementById('todo-form');
  const input = document.getElementById('task-input');
  const select = document.getElementById('type-select');

  function loadTodos() {
    const todos = JSON.parse(localStorage.getItem('todos') || '[]');
    document.querySelectorAll('.quadrant ul').forEach(ul => ul.innerHTML = '');
    todos.forEach(todo => {
      const li = document.createElement('li');
      li.textContent = todo.text;
      li.onclick = () => {
        if (confirm('刪除這項任務？')) {
          deleteTodo(todo.id);
        }
      };
      document.querySelector(`.quadrant[data-type="${todo.type}"] ul`).appendChild(li);
    });
  }

  function saveTodo(text, type) {
    const todos = JSON.parse(localStorage.getItem('todos') || '[]');
    todos.push({ id: Date.now(), text, type });
    localStorage.setItem('todos', JSON.stringify(todos));
    loadTodos();
  }

  function deleteTodo(id) {
    let todos = JSON.parse(localStorage.getItem('todos') || '[]');
    todos = todos.filter(t => t.id !== id);
    localStorage.setItem('todos', JSON.stringify(todos));
    loadTodos();
  }

  form.onsubmit = (e) => {
    e.preventDefault();
    if (input.value.trim() !== '') {
      saveTodo(input.value.trim(), select.value);
      input.value = '';
    }
  };

  loadTodos();
});
