function PostSignUp(namadepan,namabelakang,email,password){
    var myHeaders = new Headers();
    myHeaders.append("Login", "chapter03");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "namadepan": namadepan,
      "namabelakang": namabelakang,
      "email" : email,
      "notlp": notlp,
      "Ppemain": Ppemain,
    });
   
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };
   
    fetch("https://eop32fgoudfaaoh.m.pipedream.net", requestOptions)
      .then(response => response.text())
      .then(result => GetResponse(result))
      .catch(error => console.log('error', error));
    }

    function GetResponse(result){
        document.getElementById("formsignup").innerHTML = result;
    }

    function PushButton(){
        namadepan=document.getElementById("namadepan").value;
        namabelakang=document.getElementById("namabelakang").value;
        email=document.getElementById("email").value;
        notlp=document.getElementById("notlp").value;
        Ppemain=document.getElementById("Ppemain").value;
        PostSignUp(namadepan,namabelakang,notlp,Ppemain,);
        document.getElementById("formsignup").style.display = 'none';
        document.getElementById("formsignup").style.display = 'block';
    }

