function PostSignUp(namadepan,namabelakang,email,password){
    var myHeaders = new Headers();
    myHeaders.append("Login", "chapter03");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "namadepan": namadepan,
      "namabelakang": namabelakang,
      "notlp": notlp,
      "alamat": alamat,
      "booktitle": booktitle,
      "tglpinjam": tglpinjam
    });
   
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };
   
    fetch("https://eovnuau298i8ogh.m.pipedream.net", requestOptions)
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
        notlp=document.getElementById("notlp").value;
        alamat=document.getElementById("alamat").value;
        booktitle=document.getElementById("booktitle").value;
        tglpinjam=document.getElementById("tglpinjam").value;
        PostSignUp(namadepan,namabelakang,notlp,alamat,booktitle,tglpinjam);
        document.getElementById("formsignup").style.display = 'none';
        document.getElementById("formsignup").style.display = 'block';
    }

