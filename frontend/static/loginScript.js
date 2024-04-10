const logButton = document.querySelector(".login-button");
logButton.addEventListener('click', Log)


function Log() {
    let data = {
        username: document.querySelector("#username").value,
        password: document.querySelector("#pass").value
    };

    fetch('http://localhost:8080/auth-login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Methods': 'POST',
            'Access-Control-Allow-Origin': '*',
        },
        body: JSON.stringify(data),
    })
        .then(response => {
            // Обработка ответа от сервера
            if (response.status == 200) {
                console.log("successfully logged")
            } else {
                alert("Failed to handle data from server");
            }
        })
        .catch(error => {
            alert("Failed to send data to server");
        });
}
