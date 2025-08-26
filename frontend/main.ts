const API_URL = "http://localhost:8080";

// Elements
const usernameInput = document.getElementById("username") as HTMLInputElement;
const emailInput = document.getElementById("email") as HTMLInputElement;
const createUserBtn = document.getElementById("createUserBtn")!;
const titleInput = document.getElementById("title") as HTMLInputElement;
const authorInput = document.getElementById("author") as HTMLInputElement;
const userIdInput = document.getElementById("userId") as HTMLInputElement;
const createBookBtn = document.getElementById("createBookBtn")!;
const getUsersBtn = document.getElementById("getUsersBtn")!;
const getBooksBtn = document.getElementById("getBooksBtn")!;
const usersList = document.getElementById("usersList")!;
const booksList = document.getElementById("booksList")!;

// Create User
createUserBtn.addEventListener("click", async () => {
    const name = usernameInput.value;
    const email = emailInput.value;
    if (!name || !email) return alert("Enter name & email");

    const res = await fetch(`${API_URL}/users`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, email })
    });

    const data = await res.json();
    alert(`User created: ${data.name} (ID: ${data.id})`);
});

// Create Book
createBookBtn.addEventListener("click", async () => {
    const title = titleInput.value;
    const author = authorInput.value;
    const userId = parseInt(userIdInput.value);

    if (!title || !author || isNaN(userId)) return alert("Fill all book fields");

    const res = await fetch(`${API_URL}/books`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title, author, user_id: userId })
    });

    const data = await res.json();
    alert(`Book created: ${data.title} for user ${data.user_id}`);
});

// Get Users
getUsersBtn.addEventListener("click", async () => {
    const res = await fetch(`${API_URL}/users`);
    const users = await res.json();

    usersList.innerHTML = "";
    users.forEach((u: any) => {
        const li = document.createElement("li");
        li.className = "border rounded p-2";
        li.textContent = `ID: ${u.id}, Name: ${u.name}, Email: ${u.email}`;
        usersList.appendChild(li);
    });
});

// Get Books
getBooksBtn.addEventListener("click", async () => {
    const res = await fetch(`${API_URL}/books`);
    const books = await res.json();

    booksList.innerHTML = "";
    books.forEach((b: any) => {
        const li = document.createElement("li");
        li.className = "border rounded p-2";
        li.textContent = `ID: ${b.id}, Title: ${b.title}, Author: ${b.author}, UserID: ${b.user_id}`;
        booksList.appendChild(li);
    });
});
