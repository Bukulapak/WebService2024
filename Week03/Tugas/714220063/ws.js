const myHeaders = new Headers();
myHeaders.append("Login", "swww");
myHeaders.append("Content-Type", "application/json");

const raw = JSON.stringify({
  "npm": 714220063,
  "nama": "Andika",
  "no_hp": 62089999999,
  "status_akun": "true"
});

const requestOptions = {
  method: "POST",
  headers: myHeaders,
  body: raw,
  redirect: "follow"
};

fetch("https://eo63zywwzy46zgr.m.pipedream.net", requestOptions)
  .then((response) => response.text())
  .then((result) => console.log(result))
  .catch((error) => console.error(error));

function PostSignUp(firstName,lastName,email,password,c_password){
    var myHeaders = new Headers();
    myHeaders.append("Login", "chapter03");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "firstName": firstName,
      "lastName": lastName,
      "email": email,
      "password": password,
      "c_password": c_password,
    });
   
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };
   
    function GetResponse(result){
        document.getElementById("formsignup").innerHTML = '<div class="ss">' + result + '</div>';
    
    }

    fetch("https://eo63zywwzy46zgr.m.pipedream.net", requestOptions)
        .then(response => response.text())
        .then(result => GetResponse(result))
        .catch(error => console.log('error', error));
    }

    function PushButton(){
        var firstName = document.getElementById("firstName").value;
        var lastName = document.getElementById("lastName").value;
        var email = document.getElementById("email").value;
        var password = document.getElementById("password").value;
        var c_password = document.getElementById("c_password").value;
        PostSignUp(firstName, lastName, email, password, c_password);
        
    }
    
    
    