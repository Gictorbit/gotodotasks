const todoForm = document.querySelector('#todo-form');
const titleInput = document.querySelector('#title-input');
const descriptionInput = document.querySelector('#description-input');
const dateInput = document.querySelector('#date-input');
const todoCards = document.querySelector('#todo-cards');

let tasks = [];

function addTask(event) {
    event.preventDefault();
    const task = {
        id: Date.now(),
        title: titleInput.value,
        description: descriptionInput.value,
        dueDate: dateInput.value,
        completed: false,
    };
    tasks.push(task);
    renderTasks();
    todoForm.reset();
}

function renderTasks() {
    todoCards.innerHTML = '';
    tasks.forEach(task => {
        const card = document.createElement('div');
        card.classList.add('todo-card');
        const h3 = document.createElement('h3');
        h3.innerText = task.title;
        card.appendChild(h3);
        const p1 = document.createElement('p');
        p1.innerText = task.description;
        card.appendChild(p1);
        const p2 = document.createElement('p');
        p2.innerText = `Due Date: ${task.dueDate}`;
        card.appendChild(p2);
        const checkbox = document.createElement('input');
        checkbox.type = 'checkbox';
        checkbox.checked = task.completed;
        checkbox.addEventListener('change', () => {
            task.completed = !task.completed;
            renderTasks();
        });
        card.appendChild(checkbox);
        const label = document.createElement('label');
        label.innerText = 'Completed';
        card.appendChild(label);
        todoCards.appendChild(card);
    });
}

todoForm.addEventListener('submit', addTask);
renderTasks();
