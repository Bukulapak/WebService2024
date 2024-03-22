function PostSignUp(nama, email, no_telp, BirthDate, agama, alamat) {
  var myHeaders = new Headers();
  myHeaders.append("Login", "chapter03");
  myHeaders.append("Content-Type", "application/json");
  
  var raw = JSON.stringify({
      "nama": nama,
      "email": email,
      "no_telp": no_telp,
      "BirthDate": BirthDate,
      "agama": agama,
      "alamat": alamat
  });
  
  var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
  };
  
  fetch("https://eovc04iuh63k7ku.m.pipedream.net/", requestOptions)
      .then(response => response.text())
      .then(result => console.log(result))
      .catch(error => console.log('error', error));
  } 

  function GetResponse(result){
    document.getElementById("formsignup").innerHTML = result;
  }

  function PushButton() {
  var nama=document.getElementById("nama").value;
  var email=document.getElementById("email").value;
  var no_telp=document.getElementById("no_telp").value;
  var BirthDate=document.getElementById("BirthDate").value;
  var agama=document.getElementById("agama").value;
  var alamat=document.getElementById("alamat").value;
  PostSignUp(nama, email, no_telp, BirthDate, agama, alamat);
  }

  document.getElementById("formsignup").style.display = 'none';
  document.getElementById("formsignup").style.display = 'block';

  
