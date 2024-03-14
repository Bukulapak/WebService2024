const requestOptions = {
    method: "GET",
    redirect: "follow"
};

fetch("https://reqres.in/api/users", requestOptions)
    .then(response => response.json())
    .then(result => tampilkan(result))
    .catch(error => console.error(error));

function tampilkan(result) {
    console.log(result);
    let hasil = '<table border="1">';
    hasil += "<tr><th>ID</th><th>Email</th><th>Nama Depan</th><th>Nama Belakang</th><th>Avatar</th></tr>";

    result.data.forEach(user => {
        hasil += "<tr>";
        hasil += "<td>" + user.id + "</td>";
        hasil += "<td>" + user.email + "</td>";
        hasil += "<td>" + user.first_name + "</td>";
        hasil += "<td>" + user.last_name + "</td>";
        hasil += "<td><img src='" + user.avatar + "' alt='Avatar'></td>";
        hasil += "</tr>";
    });

    hasil += "</table>";

    document.getElementById("inidata").innerHTML = hasil;
}
