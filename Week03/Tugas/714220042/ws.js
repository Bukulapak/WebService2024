function PostSignUp(namadepan,namabelakang,email,password){
  var myHeaders = new Headers();
  myHeaders.append("Login", "chapter03");
  myHeaders.append("Content-Type", "application/json");
 
  var raw = JSON.stringify({
    "namadepan": namadepan,
    "namabelakang": namabelakang,
    "notlp": notlp,
    "alamat": alamat,
    "jeniskelas": jeniskelas,
    //"tglPendaftaran": tglPendaftaran
  });
 
  var requestOptions = {
    method: 'POST',
    headers: myHeaders,
    body: raw,
    redirect: 'follow'
  };
 
  fetch("https://eoye0d3l4gh6rfj.m.pipedream.net", requestOptions)
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
      booktitle=document.getElementById("jeniskelas").value;
      //tglpinjam=document.getElementById("tglPendaftaran").value;
      PostSignUp(namadepan,namabelakang,notlp,alamat,jeniskelas);
      document.getElementById("formsignup").style.display = 'none';
      document.getElementById("formsignup").style.display = 'block';
  }

