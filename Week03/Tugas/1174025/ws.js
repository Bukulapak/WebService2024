
function PostSignUp(username,email,address,phone,password){
    const myHeaders = new Headers();
    myHeaders.append("Login", "akusiapa");
    myHeaders.append("Content-Type", "application/json");
    
    var raw = JSON.stringify({
        "username": username,
        "email": email,
        "address": address,
        "phone": phone,
        "password": password
      });
    
    const requestOptions = {
      method: "POST",
      headers: myHeaders,
      body: raw,
      redirect: "follow"
    };
    
    fetch("https://eoewbv47unkcq3u.m.pipedream.net", requestOptions)
      .then((response) => response.text())
      .then((result) => GetResponse(result))
      .catch((error) => console.error(error));
    }
    function GetResponse(result){
        // document.getElementById("formsignup").style.display = 'none';
        document.getElementById("formsignup").innerHTML = result;
    }
    
    function PushButton(){
        username=document.getElementById("username").value;
        email=document.getElementById("email").value;
        address=document.getElementById("address").value;
        phone=document.getElementById("phone").value;
        password=document.getElementById("password").value;
        PostSignUp(username,email,address,phone, password);
    }
    
    