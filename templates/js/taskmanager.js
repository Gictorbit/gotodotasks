/* DOMContentLoaded */
document.addEventListener("DOMContentLoaded", main);

/* main() FUNCTION */
const serverDomain = document.getElementById('serverDomain').textContent;
const taskManagerURL = `${serverDomain}/v1/todotask`
function main() {

  // Check user Login
  const isAuth = localStorage.getItem('Token');
  if(!!!isAuth){
    location.href = 'login';
  }

  const token = localStorage.getItem('Token');
  console.log(token);
  localStorage.removeItem("todos");
  fetch(`${taskManagerURL}/taskslist`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'token':`${token}`
    },
  }).then(response => {
    if (response.ok) {
      // Handle successful response
      return response.json();
    } else {
      // Handle error response
      throw new Error('Error signing up user');
    }
  }).then(respdata=>{
    const todos = JSON.parse(localStorage.getItem("todos")) || [];
    respdata.tasks.forEach(function (todo) {
      console.log(todo)
      todos.push(todo)
    })
    localStorage.setItem('todos',JSON.stringify(todos));
    addTodo();
  }).catch(error=>{
    console.log(error);
    throw error;
  })

  // theme-switcher
  document
      .getElementById("theme-switcher")
      .addEventListener("click", function () {
        document.querySelector("body").classList.toggle("light");
        const themeImg = this.children[0];
        themeImg.setAttribute(
            "src",
            themeImg.getAttribute("src") === "./assets/images/icon-sun.svg"
                ? "./assets/images/icon-moon.svg"
                : "./assets/images/icon-sun.svg"
        );
      });

  // get alltodos and initialise listeners
  // dragover on .todos container
  document.querySelector(".todos").addEventListener("dragover", function (e) {
    e.preventDefault();
    if (
        !e.target.classList.contains("dragging") &&
        e.target.classList.contains("card")
    ) {
      const draggingCard = document.querySelector(".dragging");
      const cards = [...this.querySelectorAll(".card")];
      const currPos = cards.indexOf(draggingCard);
      const newPos = cards.indexOf(e.target);
      console.log(currPos, newPos);
      if (currPos > newPos) {
        this.insertBefore(draggingCard, e.target);
      } else {
        this.insertBefore(draggingCard, e.target.nextSibling);
      }
      const todos = JSON.parse(localStorage.getItem("todos"));
      const removed = todos.splice(currPos, 1);
      todos.splice(newPos, 0, removed[0]);
      localStorage.setItem("todos", JSON.stringify(todos));
    }
  });
  // add new todos on user input
  const add = document.getElementById("add-btn");
  const txtInput = document.querySelector(".txt-input");
  add.addEventListener("click", function () {
    const item = txtInput.value.trim();
    if (item) {
      txtInput.value = "";
      const todos = JSON.parse(localStorage.getItem("todos")) || [];
      const currentTodo = {
        task: {
          id:0,
          title:item,
          description:'',
          isCompleted: false,
        }
      };
      console.log(currentTodo)
      fetch(`${taskManagerURL}/createtask`,{
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'token': token
        },
        body: JSON.stringify(currentTodo)
      }).then((response) => {
        if (response.ok) {
          return response.json(); // extract the response body as JSON
        }
        throw new Error("Network response was not ok");
      }).then((data) => {
        currentTodo.task.id = data.taskId; // update the currentTodo ID
        todos.push(currentTodo);
        localStorage.setItem("todos", JSON.stringify(todos));
        console.log("created todo:",currentTodo)
        addTodo([currentTodo.task]);
      }).catch((error) => {
        console.error("Error adding todo:", error);
        throw error;
      });
    }
    txtInput.focus();
  });
  // add todo also on enter key event
  txtInput.addEventListener("keydown", function (e) {
    if (e.keyCode === 13) {
      add.click();
    }
  });
  // filter todo - all, active, completed
  document.querySelector(".filter").addEventListener("click", function (e) {
    const id = e.target.id;
    if (id) {
      document.querySelector(".on").classList.remove("on");
      document.getElementById(id).classList.add("on");
      document.querySelector(".todos").className = `todos ${id}`;
    }
  });
  // clear completed
  document
      .getElementById("clear-completed")
      .addEventListener("click", function () {
        deleteIndexes = [];
        document.querySelectorAll(".card.checked").forEach(function (card) {
          deleteIndexes.push(
              [...document.querySelectorAll(".todos .card")].indexOf(card)
          );
          card.classList.add("fall");
          card.addEventListener("animationend", function (e) {
            setTimeout(function () {
              card.remove();
            }, 100);
          });
        });
        removeManyTodo(deleteIndexes);
      });
}

/* stateTodo() FUNCTION TO UPDATE TODO ABOUT COMPLETION */

function stateTodo(index, completed) {
  const todos = JSON.parse(localStorage.getItem("todos"));
  const token = localStorage.getItem('Token');
  todos[index].isCompleted = completed;
  localStorage.setItem("todos", JSON.stringify(todos));
  const updateTodo = {
    task: {
      id: todos[index].id,
      title: todos[index].title,
      description: todos[index].description,
      isCompleted: completed,
    }
  };
  console.log("todo task update",updateTodo)

  fetch(`${taskManagerURL}/updatetask`,{
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'token': token
    },
    body: JSON.stringify(updateTodo)
  }).then((response) => {
    if (response.ok) {
      return response.json(); // extract the response body as JSON
    }
    throw new Error("Network response was not ok");
  }).catch((error) => {
    console.error("Error adding todo:", error);
    throw error;
  });
}

/* removeManyTodo() FUNCTION TO REMOVE ONE TODO */

function removeTodo(index,todoId) {
  // Check user Login
  const token = localStorage.getItem('Token');
  if(!!!token){
    location.href = 'login';
  }
  console.log("remove ID",todoId)
  fetch(`${taskManagerURL}/deletetask/${todoId}`,{
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
      'token': token
    },
  }).then(response=>{
    if(response.ok){
      const todos = JSON.parse(localStorage.getItem("todos"));
      todos.splice(index, 1);
      localStorage.setItem("todos", JSON.stringify(todos));
    }
  }).catch(error=>{
    console.log(error);
    throw error;
  })
}

/* removeManyTodo FUNCTION TO REMOVE MANY TODOS */

function removeManyTodo(indexes) {
  let todos = JSON.parse(localStorage.getItem("todos"));
  indexes.forEach(index => {
    removeTodo(index,todos[index].id);
  });
  todos = todos.filter(function (todo, index) {
    return !indexes.includes(index);
  });
  localStorage.setItem("todos", JSON.stringify(todos));
}

/* addTodo() FUNCTION TO LIST/CREATE TODOS AND ADD EVENT LISTENERS */

function addTodo(todos = JSON.parse(localStorage.getItem("todos"))) {
  if (!todos) {
    return null;
  }
  const itemsLeft = document.getElementById("items-left");
  // create cards
  todos.forEach(function (task) {
    const card = document.createElement("li");
    const cbContainer = document.createElement("div");
    const cbInput = document.createElement("input");
    const check = document.createElement("span");
    const item = document.createElement("p");
    const button = document.createElement("button");
    const img = document.createElement("img");
    // Add classes
    card.classList.add("card");
    button.classList.add("clear");
    cbContainer.classList.add("cb-container");
    cbInput.classList.add("cb-input");
    item.classList.add("item");
    check.classList.add("check");
    button.classList.add("clear");
    // Set attributes
    card.setAttribute("draggable", true);
    img.setAttribute("src", "assets/images/icon-cross.svg");
    img.setAttribute("alt", "Clear it");
    cbInput.setAttribute("type", "checkbox");
    // set todo item for card
    item.textContent = task.title;
    // if completed -> add respective class / attribute
    if (task.isCompleted) {
      card.classList.add("checked");
      cbInput.setAttribute("checked", "checked");
    }
    // Add drag listener to card
    card.addEventListener("dragstart", function () {
      this.classList.add("dragging");
    });
    card.addEventListener("dragend", function () {
      this.classList.remove("dragging");
    });
    // Add click listener to checkbox
    cbInput.addEventListener("click", function () {
      const correspondingCard = this.parentElement.parentElement;
      const checked = this.checked;
      stateTodo(
          [...document.querySelectorAll(".todos .card")].indexOf(
              correspondingCard
          ),
          checked
      );
      checked
          ? correspondingCard.classList.add("checked")
          : correspondingCard.classList.remove("checked");
      itemsLeft.textContent = document.querySelectorAll(
          ".todos .card:not(.checked)"
      ).length;
    });
    // Add click listener to clear button
    button.addEventListener("click", function () {
      const correspondingCard = this.parentElement;
      correspondingCard.classList.add("fall");
      removeTodo(
          [...document.querySelectorAll(".todos .card")].indexOf(
              correspondingCard
          ),
          task.id
      );
      correspondingCard.addEventListener("animationend", function () {
        setTimeout(function () {
          correspondingCard.remove();
          itemsLeft.textContent = document.querySelectorAll(
              ".todos .card:not(.checked)"
          ).length;
        }, 100);
      });
    });
    // parent.appendChild(child)
    button.appendChild(img);
    cbContainer.appendChild(cbInput);
    cbContainer.appendChild(check);
    card.appendChild(cbContainer);
    card.appendChild(item);
    card.appendChild(button);
    document.querySelector(".todos").appendChild(card);
  });
  // Update itemsLeft
  itemsLeft.textContent = document.querySelectorAll(
      ".todos .card:not(.checked)"
  ).length;
}
