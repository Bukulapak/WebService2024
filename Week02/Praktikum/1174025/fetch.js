const myHeaders = new Headers();
myHeaders.append("Cookie", "connect.sid=s%3Az3QRPcoatnMvXjMzG4VezIvFJ_Usy0-_.e%2FlRGUrLBeQqFjYoSSm%2F1Ehli9qGXKUs6Xpy9xtHrqw");

const requestOptions = {
  method: "GET",
  headers: myHeaders,
  redirect: "follow"
};
hasil = "";

fetch("https://cat-fact.herokuapp.com/facts", requestOptions)
  .then(response => response.text())
  .then(result => tampilkan(result))
  .catch(error => console.log('error', error));

  function tampilkan(result){
    console.log(result);
    iniJson = JSON.parse(result);

    length = iniJson.length;
    console.log(length);

    for (i = 0; i < length; i++) {
        hasil += "<tr>";        
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+iniJson[i].text+"</td>";        
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+iniJson[i].type+"</td>";        
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+iniJson[i].source+"</td>";        
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+iniJson[i].createdAt+"</td>";        
        hasil += "</tr>";        
    }

    document.getElementById("inidata").innerHTML = hasil;
  }



// fetch("https://cat-fact.herokuapp.com/facts", requestOptions)
//   .then(response => response.json())
//   .then(data => showCharacters(data))
//   .catch(error => console.error(error));

// function showCharacters(data) {
//     const tbody = document.querySelector('tbody'); // put data ke element
//     tbody.innerHTML = ''; 
  
//     data.forEach(item => {  // looping
//       const row = document.createElement('tr');
//       row.classList.add('hover:bg-gray-50');
  
//       const textCell = document.createElement('td');
//       textCell.classList.add('px-6', 'py-4', 'font-medium');
//       textCell.textContent = item.text;

//       const textCell2 = document.createElement('td');
//       textCell2.classList.add('px-6', 'py-4', 'font-medium');
//       textCell2.textContent = item.type;

//       //menambahkan data pada masing" row sesuai kolom
//       row.appendChild(textCell);
//       row.appendChild(textCell2);
  
//       tbody.appendChild(row);
//     });
//   }