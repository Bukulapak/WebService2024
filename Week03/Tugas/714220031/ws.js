function PostSignUp(nama,npm,prodi,no_telp,alasan_masuk){
    var myHeaders = new Headers();
    myHeaders.append("Login", "ghaida");
    myHeaders.append("Content-Type", "application/json");


var raw = JSON.stringify({
    "nama": nama,
    "npm": npm,
    "prodi": prodi,
    "no_telp": no_telp,
    "alasan_masuk": alasan_masuk
  });
 
  var requestOptions = {
    method: 'POST',
    headers: myHeaders,
    body: raw,
    redirect: 'follow'
  };

fetch("https://eotnalb4eyiywue.m.pipedream.net", requestOptions)
  .then((response) => response.text())
  .then((result) => console.log(result))
  .catch((error) => console.error(error));
}

  function PushButton(){
    nama=document.getElementById("nama").value;
    npm=document.getElementById("npm").value;
    prodi=document.getElementById("prodi").value;
    no_telp=document.getElementById("no_telp").value;
    alasan_masuk=document.getElementById("alasan_masuk").value;
    PostSignUp(nama,npm,prodi,no_telp,alasan_masuk);
}

function GetResponse(result){
document.getElementById("formsignup").innerHTML = result;
}
