function PostSignUp(namalengkap,namadepan,tempattgllahir,alamat,notelphp,alamatorangtuawali,emailaddress,noTelpHP){
    var myHeaders = new Headers();
    myHeaders.append("Login", "chapter03");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "namalengkap": namalengkap,
      "namapanggilan": namapanggilan,
      "tempattgllahir": tempattgllahir,
      "alamat": alamat,
      "notelphp": notelphp,
      "alamatorangtuawali": alamatorangtuawali,
      "emailaddress": emailaddress,
      "noTelpHP": noTelpHP
    });
   
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };
   
    fetch("https://eobh0ts7b6p4pfu.m.pipedream.net", requestOptions)
      .then(response => response.text())
      .then(result => GetResponse.log(result))
      .catch(error => console.log('error', error));
    }

    function GetResponse(result){
        document.getElementById("formsignup").innerHTML = result;
  }

  function PushButton(){
    namalengkap=document.getElementById("namalengkap").value;
    namapanggilan=document.getElementById("namapanggilan").value;
    tempattgllahir=document.getElementById("tempattgllahir").value;
    alamat=document.getElementById("alamat").value;
    notelphp=document.getElementById("notelphp").value;
    alamatorangtuawali=document.getElementById("alamatorangtuawali").value;
    emailaddress=document.getElementById("emailaddresst").value;
    noTelpHP=document.getElementById("noTelpHP").value;
    PostSignUp(namalengkap,namapanggilan,tempattgllahir,alamat,notelphp,alamatorangtuawali,emailaddress,noTelpHP);
}