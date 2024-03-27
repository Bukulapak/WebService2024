function PostSignUp(namalengkap,nomortelepon,haritanggal,waktu,acara){
    var myHeaders = new Headers();
    myHeaders.append("Login", "chapter03");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "namalengkap": namalengkap,
      "nomortelepon": nomortelepon,
      "haritanggal": haritanggal,
      "waktu": waktu,
      "acara": acara
    });
   
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };
   
    fetch("https://eovags4f749tjr7.m.pipedream.net", requestOptions)
      .then(response => response.text())
      .then(result => GetResponse(result))
      .catch(error => console.log('error', error));
    }
    
    function GetResponse(result){
      document.getElementById("formreserve").innerHTML = result;
  }

    function PushButton(){
        namalengkap=document.getElementById("namalengkap").value;
        nomortelepon=document.getElementById("nomortelepon").value;
        haritanggal=document.getElementById("haritanggal").value;
        waktu=document.getElementById("waktu").value;
        acara=document.getElementById("acara").value;
        PostSignUp(namalengkap,nomortelepon,haritanggal,waktu,acara);
  }
  document.getElementById("formreserve").style.display = 'none';
  document.getElementById("formreserve").style.display = 'block';
