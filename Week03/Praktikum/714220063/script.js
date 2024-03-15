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

    function GetResponse(result){
        document.getElementById("formsignup").innerHTML = result;
    }
   
    fetch("https://eo63zywwzy46zgr.m.pipedream.net", requestOptions)
      .then(response => response.text())
      .then(result => GetResponse(result))
      .catch(error => console.log('error', error));
    }

    function PushButton(){
        namadepan=document.getElementById("namadepan").value;
        namabelakang=document.getElementById("namabelakang").value;
        email=document.getElementById("email").value;
        password=document.getElementById("password").value;
        PostSignUp(namadepan,namabelakang,email,password);
        console.log("Kepencet bang")
    }
    