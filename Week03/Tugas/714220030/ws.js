function PostSignUp(username,ideuy,email,password,c_password){
    var myHeaders = new Headers();
    myHeaders.append("Login", "chapter03");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "username": username,
      "ideuy": ideuy,
      "email": email,
      "password": password,
      "c_password": c_password
    });
   
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };

    
   
    fetch("https://eobt8as1w4ptl3n.m.pipedream.net", requestOptions)
      .then(response => response.text())
      .then(result => GetResponse(result))
      .catch(error => console.log('error', error));
    }

    

    function PushButton(){
        username=document.getElementById("username").value;
        ideuy=document.getElementById("ideuy").value;
        email=document.getElementById("email").value;
        password=document.getElementById("password").value;
        c_password=document.getElementById("c_password").value;
        PostSignUp(username,ideuy,email,password,c_password);
  }

  function GetResponse(result){
    document.getElementById("formsignup").innerHTML = result + '<br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br>';
}