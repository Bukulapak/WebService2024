function PostSignUp(namapemesan, jenislaptop, email, telepon, deskripsi) {
  var myHeaders = new Headers();
  myHeaders.append('Login', 'chapter03');
  myHeaders.append('Content-Type', 'application/json');

  var raw = JSON.stringify({
    namapemesan: namapemesan,
    jenislaptop: jenislaptop,
    email: email,
    telepon: telepon,
    deskripsi: deskripsi,
  });

  var requestOptions = {
    method: 'POST',
    headers: myHeaders,
    body: raw,
    redirect: 'follow',
  };

  fetch('https://eom1mrrovab0a22.m.pipedream.net', requestOptions)
    .then((response) => response.text())
    .then((result) => GetResponse(result))
    .catch((error) => console.error('error', error));
}

function GetResponse(result) {
  document.getElementById('formsignup').innerHTML = result;
}

function PushButton() {
  var namapemesan = document.getElementById('namapemesan').value;
  var jenislaptop = document.getElementById('jenislaptop').value;
  var email = document.getElementById('email').value;
  var telepon = document.getElementById('telepon').value;
  var deskripsi = document.getElementById('deskripsi').value;
  PostSignUp(namapemesan, jenislaptop, email, telepon, deskripsi);
}
