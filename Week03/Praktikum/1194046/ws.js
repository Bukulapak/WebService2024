
function PostSignUp(namadepan,namabelakang,email,password){
    const myHeaders = new Headers();
    myHeaders.append("Login", "akusiapa");
    myHeaders.append("Content-Type", "application/json");
    
    var raw = JSON.stringify({
        "namadepan": namadepan,
        "namabelakang": namabelakang,
        "email": email,
        "password": password
      });
    
    const requestOptions = {
      method: "POST",
      headers: myHeaders,
      body: raw,
      redirect: "follow"
    };
    
    fetch("https://eom1gff5vapqnja.m.pipedream.net", requestOptions)
    .then((response) => response.text())
    .then((result) => GetResponse(result))
    .catch((error) => console.error(error));
  }
  function GetResponse(result){
      // document.getElementById("formsignup").style.display = 'none';
      document.getElementById("formsignup").innerHTML = result;
  }
  
  function PushButton(){
      namadepan=document.getElementById("namadepan").value;
      namabelakang=document.getElementById("namabelakang").value;
      email=document.getElementById("email").value;
      password=document.getElementById("password").value;
      PostSignUp(namadepan,namabelakang,email,password);
  }