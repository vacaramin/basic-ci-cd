const API_BASE_URL = "http://api.vacaramin.me";

async function fetchUsers() {
    try {
        const response = await fetch(`${API_BASE_URL}/users`);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const users = await response.json();
        displayUsers(users);
    } catch (error) {
        console.error("There was a problem fetching the user data:", error);
    }
}

function displayUsers(users) {
    const usersList = document.getElementById('usersList');
    // Clear the list first
    usersList.innerHTML = '';

    // Create a table element
    const table = document.createElement('table');
    table.classList.add('user-table'); // Add a class for styling

    // Create a table header
    const thead = document.createElement('thead');
    thead.innerHTML = `<tr>
        <th>S#</th>
        <th>Name</th>
        <th>Age</th>
    </tr>`;
    table.appendChild(thead);

    // Create a table body
    const tbody = document.createElement('tbody');
    
    let count= 1; 
    users.forEach(user => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${count}</td>
            <td>${user.name}</td>
            <td>${user.age}</td>
        `;
        count += 1;
        tbody.appendChild(row);
    });

    table.appendChild(tbody);
    usersList.appendChild(table);
}

// Call fetchUsers to populate the data on load
fetchUsers();
