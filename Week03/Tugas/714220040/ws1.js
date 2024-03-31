function PostSignUp(nama, no_telp, email, birthDate, alamat) {
    var myHeaders = new Headers();
    myHeaders.append("Login", "chapter03");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
        "nama": nama,
        "no_telp": no_telp,
        "email": email,
        "birthDate": birthDate,
        "alamat": alamat
    });
   
    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };
   
    fetch("https://eowinodzalilvvr.m.pipedream.net", requestOptions)
        .then(response => response.text())
        .then(result => GetResponse(result)) // call GetResponse function with result
        .catch(error => console.log('error', error));
}

function GetResponse(result) {
    document.getElementById("formsignup").innerHTML = result;
}

function PushButton() {
    var nama = document.getElementById("nama").value;
    var no_telp = document.getElementById("no_telp").value;
    var email = document.getElementById("email").value;
    var birthDate = document.getElementById("birthDate").value;
    var alamat = document.getElementById("alamat").value;
    PostSignUp(nama, no_telp, email, birthDate, alamat); // fix the order of parameters
}
