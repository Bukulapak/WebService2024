function PostSignUp(firstname,lastname,email,phone,password){
    var myHeaders = new Headers();
    myHeaders.append("Register", "ValorantRegister");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "firstname": firstname,
      "lastname": lastname,
      "email": email,
      "phone" : phone,
      "password": password
    });
   
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };
   
fetch("https://eownd6iejeqjyrm.m.pipedream.net", requestOptions)
  .then((response) => response.text())
  .then((result) => GetResponse(result))
  .catch((error) => console.log('Error:', error));
}

function GetResponse(result){
  document.getElementById("signupform").innerHTML = result;
}

    function SignupButton(){
        firstname=document.getElementById("firstname").value;
        lastname=document.getElementById("lastname").value;
        email=document.getElementById("email").value;
        phone=document.getElementById("phone").value;
        password=document.getElementById("password").value;
        PostSignUp(firstname,lastname,phone,email,password);
    }
