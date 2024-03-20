function PostSignUp(nama, namaanabul, jenisusiaanabul, jenisgrooming, notelepon, alamat) {
    var myHeaders = new Headers();
    myHeaders.append("Login", "deviwlndr");
    myHeaders.append("Content-Type", "application/json");

    var raw = JSON.stringify({
        "nama": nama,
        "namaanabul": namaanabul,
        "jenisusiaanabul": jenisusiaanabul,
        "jenisgrooming": jenisgrooming,
        "notelepon": notelepon,
        "alamat": alamat
    });

    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };

    fetch("https://eocne76kfig61v3.m.pipedream.net", requestOptions)
        .then(response => response.text())
        .then(result => GetResponse(result))
        .catch(error => console.error('Error:', error)); // Menampilkan pesan error di konsol
}

function GetResponse(result) {
    document.getElementById("formsignup").innerHTML = result;
}

function PushButton() {
    var nama = document.getElementById("nama").value;
    var namaanabul = document.getElementById("namaanabul").value;
    var jenisusiaanabul = document.getElementById("jenisusiaanabul").value;
    var jenisgrooming = document.getElementById("jenisgrooming").value;
    var notelepon = document.getElementById("notelepon").value;
    var alamat = document.getElementById("alamat").value;
    PostSignUp(nama, namaanabul, jenisusiaanabul, jenisgrooming, notelepon, alamat);
}
