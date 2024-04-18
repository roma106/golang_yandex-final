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
                window.location.href = "http://localhost:8080/calc?username=" + data.username;
                return
            } else if(response.status == 403){
                alert("User with such name already exists");
            }else {
                alert("Something went wrong on server");
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
    }else if(pass1.length < 8 || pass2.length < 8){
        alert("Password must be at least 8 characters long");
    }else if(pass1!=pass2){
        alert("Passwords don't match");
        return false;
    }
    return true
}

let passInput = document.querySelector("#pass");
let passInput2 = document.querySelector("#pass2");
passInput.addEventListener('input', ()=>{
    if (passInput.value.length < 8){
      passInput.style.borderBottom = "2px solid red";
      regButton.style.opacity = "0.5";
    }else{
      passInput.style.borderBottom = "2px solid black";
      regButton.style.opacity = "1";
    }
});
passInput2.addEventListener('input', ()=>{
    if (passInput2.value.length < 8){
      passInput2.style.borderBottom = "2px solid red";
      regButton.style.opacity = "0.5";
    }else{
      passInput2.style.borderBottom = "2px solid black";
      regButton.style.opacity = "1";
    }
});