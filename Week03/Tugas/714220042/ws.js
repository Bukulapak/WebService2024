function PostSignUp(namadepan,namabelakang,notlp,alamat,jeniskelas,jeniskelamin,tingkat){
  var myHeaders = new Headers();
  myHeaders.append("Login", "chapter03");
  myHeaders.append("Content-Type", "application/json");
 
  var raw = JSON.stringify({
    "namadepan": namadepan,
    "namabelakang": namabelakang,
    "notlp": notlp,
    "alamat": alamat,
    "jeniskelas": jeniskelas,
    "jeniskelamin": jeniskelamin,
    "tingkat": tingkat,
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
      jeniskelamin=document.getElementById("jeniskelamin").value;
      tingkat=document.getElementById("tingkat").value;
      //tglpendaftaran=document.getElementById("tglPendaftaran").value;
      PostSignUp(namadepan,namabelakang,notlp,alamat,jeniskelas,jeniskelamin,tingkat);
      document.getElementById("formsignup").style.display = 'none';
      document.getElementById("formsignup").style.display = 'block';
  }

