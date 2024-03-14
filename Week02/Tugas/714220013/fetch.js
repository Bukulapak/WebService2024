const requestOptions = {
    method: "GET",
    redirect: "follow"
};

fetch("https://emojihub.yurace.pro/api/all", requestOptions)
    .then(response => response.json())
    .then(result => tampilkan(result))
    .catch(error => console.error(error));

function tampilkan(result) {
    console.log(result);
    let tableBody = document.getElementById("inidata");

    result.forEach(photo => {
        let row = `<tr>
            <td>${photo.name}</td>
            <td>${photo.category}</td>
            <td>${photo.group}</td>
            <td>${photo.htmlCode}</td>
        </tr>`;

        tableBody.innerHTML += row;
    });
}
