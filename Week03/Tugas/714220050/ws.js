function SendSuperheroData(superheroData) {
    const myHeaders = new Headers();
    myHeaders.append("Login", "divacute");
    myHeaders.append("Content-Type", "application/json");

    const raw = JSON.stringify(superheroData);

    const requestOptions = {
        method: "POST",
        headers: myHeaders,
        body: raw,
        redirect: "follow"
    };

    fetch("https://eo6zlnf1tx5w96y.m.pipedream.net", requestOptions)
        .then(response => response.text())
        .then(result => console.log(result))
        .catch(error => console.error(error));
}

function GetSuperheroFormData() {
    const namaSuperhero = document.getElementById("nama-superhero").value;
    const kodeRahasia = document.getElementById("kode-rahasia").value;
    const gelarKehormatan = document.getElementById("gelar-kehormatan").value;

    const kekuatanUtamaInputs = document.querySelectorAll("input[name='kekuatan-utama']:checked");
    const kekuatanUtama = Array.from(kekuatanUtamaInputs).map(input => input.value);
    

    const pemberitahuan = document.querySelector("input[name='pemberitahuan']:checked").value;

    const superheroData = {
        namaSuperhero: namaSuperhero,
        kodeRahasia: kodeRahasia,
        gelarKehormatan: gelarKehormatan,
        kekuatanUtama: kekuatanUtama,
        pemberitahuan: pemberitahuan
    };

    SendSuperheroData(superheroData);
    
}
function GetResponse(result){
    document.getElementById("formsignup").innerHTML = result
}
