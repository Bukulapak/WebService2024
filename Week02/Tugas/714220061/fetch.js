const requestOptions = {
  method: "GET",
  redirect: "follow",
};

function loadQt() {
  fetch("https://animechan.xyz/api/quotes", requestOptions)
    .then((response) => response.json())
    .then((result) => showQuotes(result))
    .catch((error) => console.error(error));
}

function showQuotes(quotes) {
  console.log(quotes);
  dataqt = "";

  for (i = 0; i < quotes.length; i++) {
    dataqt += "<tr>";
    dataqt += "<td>" + quotes[i].id + "</td>";
    dataqt += "<td>" + quotes[i].quote + "</td>";
    dataqt += "<td>" + quotes[i].anime + "</td>";
    dataqt += "<td>" + quotes[i].character + "</td>";
    dataqt += "</tr>";
  }

  document.getElementById("thisQuote").innerHTML = dataqt;
}
function generateNewQt() {
  loadQt();
}
loadQt();

function loadImg() {
  fetch("https://nekos.best/api/v2/neko?amount=3", requestOptions)
    .then((response) => response.json())
    .then((result) => shoWaifu(result))
    .catch((error) => console.error(error));
}


function shoWaifu(result) {
  console.log(result);
  datawf = "";
  getdataWf = result.results;

  for (i = 0; i < getdataWf.length; i++) {
    datawf += "<tr>";
    datawf +=
      "<td><a class='text-break' href='" +
      getdataWf[i].artist_href +
      "'>" +
      getdataWf[i].artist_href +
      "</a></td>";
    datawf += "<td>" + getdataWf[i].artist_name + "</td>";
    datawf +=
      "<td><a class='text-break' href='" +
      getdataWf[i].source_url +
      "'>" +
      getdataWf[i].source_url +
      "</a></td>";
    datawf +=
      "<td><img src='" +
      getdataWf[i].url +
      "' alt='Neko Image' class='nekoimg'></td>";
    datawf += "</tr>";
  }
  document.getElementById("thisWaifu").innerHTML = datawf;
}
function generateNewImg() {
  loadImg();
}
loadImg();

// Nyoba fitur

document.addEventListener("DOMContentLoaded", function() {
  var audioControl = document.getElementById("audioControl");
  var bgmp3 = document.getElementById("bgmp3");

  audioControl.addEventListener("click", function() {
    if (bgmp3.paused) {
      bgmp3.play();
      audioControl.innerHTML = '<img class="float-end" style="position: relative; top: 10px; right: 10px;" src="https://i.ibb.co/JB73jMv/unmute.png" alt="">';
    } else {
      bgmp3.pause();
      audioControl.innerHTML = '<img class="float-end" style="position: relative; top: 10px; right: 10px;" src="https://i.ibb.co/Qjp9VKB/mute.png" alt="">';
    }
  });

  bgmp3.addEventListener("ended", function() {
    bgmp3.currentTime = 0;
    bgmp3.play();
  });
});
