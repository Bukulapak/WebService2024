function PostSignUp(namadepan, namabelakang, email, password, number, belimenu) {
    var myHeaders = new Headers();
    myHeaders.append("Login", "chapter03");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
        "namadepan": namadepan,
        "namabelakang": namabelakang,
        "email": email,
        "password": password,
        "number": number,
        "belimenu": belimenu    
    });
   
    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };
   
    fetch("https://eopz56f1hs28mrm.m.pipedream.net", requestOptions)
      .then(response => response.text())
      .then(result => console.log(result))
      .catch(error => console.log('error', error));
}

function PushButton() {
     namadepan = document.getElementById("namadepan").value;
     namabelakang = document.getElementById("namabelakang").value;
     email = document.getElementById("email").value;
     password = document.getElementById("password").value;
     number = document.getElementById("number").value;
     belimenu = document.getElementById("belimenu").value;
    PostSignUp(namadepan, namabelakang, email, password, number, belimenu);
}
