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
                console.log("calc page must be loaded!")
                window.location.href = "http://localhost:8080/calc?username=" + data.username;
            }else if(response.status == 403){
                alert("Wrong username or password");
            }else if(response.status == 500){
                alert("Something went wrong on server");
            }else {
                return response.json();
            }
        })
        .then(data => {
            if (data=="go to register") {
                window.location.href = "http://localhost:8080/register";
            }else{
                return
            }
        })
}

