const regButton = document.querySelector(".register-button");
regButton.addEventListener('click', Reg)

function Reg() {
    if (!Validate()){
        return
    }
    let data = {
        username: document.querySelector("#username").value,
        password: document.querySelector("#pass").value
    };

    fetch('http://localhost:8080/auth-reg', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Methods': 'POST',
            'Access-Control-Allow-Origin': '*',
        },
        body: JSON.stringify(data),
    }).then(response => {
            // Обработка ответа от сервера
            if (response.status == 200) {
                console.log("successfully registered")
            } else {
                alert("Failed to handle data from server");
            }
        })
        .catch(error => {
            alert("Failed to send data to server");
        });
}

function Validate(){
    var pass1 = document.querySelector("#pass").value;
    var pass2 = document.querySelector("#pass2").value;
    var username = document.querySelector("#username").value
    if (username == "" || username==null || pass1==null || pass2==null || pass1=="" || pass2==""){
        alert("Please enter valid data");
        return false;
    }else if(pass1!=pass2){
        alert("Passwords don't match");
        return false;
    }
    return true
}