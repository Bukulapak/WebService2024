
function PostSignUp(fname,lname,email,password){
    const myHeaders = new Headers();
    myHeaders.append("Login", "akusiapa");
    myHeaders.append("Content-Type", "application/json");
    
    var raw = JSON.stringify({
        "namadepan": fname,
        "namabelakang": lname,
        "email": email,
        "password": password
      });
    
    const requestOptions = {
      method: "POST",
      headers: myHeaders,
      body: raw,
      redirect: "follow"
    };
    
    fetch("https://eoggappx4mtwf6v.m.pipedream.net", requestOptions)
      .then((response) => response.text())
      .then((result) => GetResponse(result))
      .catch((error) => console.error(error));
    }
    function GetResponse(result){
        // document.getElementById("formsignup").style.display = 'none';
        document.getElementById("formsignup").innerHTML = result;
    }
    
    function PushButton(){
        namadepan=document.getElementById("fname").value;
        namabelakang=document.getElementById("lname").value;
        email=document.getElementById("email").value;
        password=document.getElementById("password").value;
        PostSignUp(namadepan,namabelakang,email,password);
    }
    
    