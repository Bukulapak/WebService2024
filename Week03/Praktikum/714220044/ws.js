function PostSignUp(namadepan,namabelakang,email,password){
    var myHeaders = new Headers();
    myHeaders.append("Login", "chapter03");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "namadepan": namadepan,
      "namabelakang": namabelakang,
      "email": email,
      "password": password
    });
   
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };
   
    fetch("https://eou4fr6ul056dfh.m.pipedream.net", requestOptions)
  .then((response) => response.text())
  .then((result) => GetResponse(result))
  .catch((error) => console.error(error));      
    }

    function PushButton(){
        namadepan=document.getElementById("namadepan").value;
        namabelakang=document.getElementById("namabelakang").value;
        email=document.getElementById("email").value;
        password=document.getElementById("password").value;
        PostSignUp(namadepan,namabelakang,email,password);
  }
  function GetResponse(result){
    document.getElementById("formsignup").innerHTML = result;
}