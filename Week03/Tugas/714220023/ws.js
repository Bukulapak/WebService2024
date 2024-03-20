function PostSignUp(NamaLengkap,Alamattinggal,Nomortelepon,Kewarganegaraan,NomorVisa){
    var myHeaders = new Headers();
    myHeaders.append("login", "akuserli");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "NamaLengkap": NamaLengkap,
      "Alamattinggal": Alamattinggal,
      "Nomortelepon":Nomortelepon,
      "Kewarganegaraan": Kewarganegaraan,
      "NomorVisa": NomorVisa
    });
   
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };
   
    fetch("https://eoeibxdmb5yp66z.m.pipedream.net", requestOptions)
      .then(response => response.text())
      .then(result => GetResponse(result))
      .catch(error => console.log('error', error));
    }

    function GetResponse(result){
        document.getElementById("formsignup").innerHTML = result;
 }
    function PushButton(){
        NamaLengkap=document.getElementById("Nama Lengkap").value;
        Alamattinggal=document.getElementById("Alamat tinggal").value;
        Nomortelepon=document.getElementById("Nomor telepon").value;
        Kewarganegaraan=document.getElementById("Kewarganegaraan").value;
        NomorVisa=document.getElementById("Nomor Visa").value;
        PostSignUp(NamaLengkap,Alamattinggal,Nomortelepon,Kewarganegaraan,NomorVisa);
  }