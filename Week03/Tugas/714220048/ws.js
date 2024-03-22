function PostSignUp(namadepan,namabelakang,email,password){
    var myHeaders = new Headers();
    myHeaders.append("Login", "chapter03");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "namadepan": namadepan,
      "namabelakang": namabelakang,
      "notlp": notlp,
      "alamat": alamat,
      "kelas": kelas,
      "matapelajaran": matapelajaran
    });
   
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };
   
    fetch("https://eotdq6p91bvzl9v.m.pipedream.net", requestOptions)
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
        kelas=document.getElementById("kelas").value;
        matapelajaran=document.getElementById("matapelajaran").value;
        PostSignUp(namadepan,namabelakang,notlp,alamat,kelas,matapelajaran);
        document.getElementById("formsignup").style.display = 'none';
        document.getElementById("formsignup").style.display = 'block';
    }

