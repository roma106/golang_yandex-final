

// Валидация поля ввода

function validateExpression(input) {
  // Используем регулярное выражение для проверки
  var regex = /^[0-9+\-*/.() ]*$/;
  return regex.test(input) && input.length > 2;
}

let exprInput = document.querySelector(".expr-input");
exprInput.addEventListener('input', ()=>{
    if (!validateExpression(exprInput.value)){
      document.querySelector(".input-container").style.borderBottom = "2px solid red";
      sendButton.style.opacity = "0.5";
    }else{
      document.querySelector(".input-container").style.borderBottom = "2px solid black";
      sendButton.style.opacity = "1";
    }
});


// отправка выражения на сервер

let sendButton = document.querySelector(".send-btn")

sendButton.addEventListener('click', sendData);


function sendData() {

  if (!validateExpression(exprInput.value)){
    alert("Invalid expression");
    return
  }
  let data = {
    username: window.location.search.split("=")[1],
    expression: document.querySelector(".expr-input").value,
    time: document.querySelector(".expr-time").value
  };
  console.log(localStorage.getItem('token'));

  // Отправка данных на сервер
  fetch('http://localhost:8080/new-expr', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Access-Control-Allow-Methods': 'POST',
      'Access-Control-Allow-Origin': '*',
      'Authorization': localStorage.getItem('token'),
    },
    body: JSON.stringify(data),
  })
  .then(response => {
    // Обработка ответа от сервера
    if (response.status == 200) {
      window.location.reload();
    }else if(response.status==401){
      alert("Your Authorization token(jwt) has expired. Please log in again.");
      window.location.href = "http://localhost:8080/login"; 
    }else {
      alert("Failed to handle data from server");
    }
  })
  .catch(error => {
    alert("Failed to send data to server");
  });
}
