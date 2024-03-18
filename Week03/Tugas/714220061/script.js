const images = [
    "assets/img/GITat-kb0AAbNmi.jpg",
    "assets/img/F_iGS6-a8AAvCJx.jpg",
    "assets/img/FPAo9ZLVQAYYa8W.jpg",
    "assets/img/FUZqGHIVUAEyHMM.jpg",
    "assets/img/Fs8cvd0aAAEMYB5.jpg",
    "assets/img/F76OHK6akAA9aPu.jpg"
];

const imageContainer = document.getElementById('imageContainer');

let currentIndex = 0;

function changeBackground() {
    const imageUrl = images[currentIndex];
    imageContainer.style.backgroundImage = `url('${imageUrl}')`;

    var box = document.getElementById('imageContainer');
    if (box.style.opacity === '1') {
        fadeOut(box);
    } else {
        fadeIn(box);
    }
    currentIndex = (currentIndex + 1) % images.length;
}

function fadeIn(element) {
    var opacity = 0;
    var interval = setInterval(function () {
        if (opacity < 1) {
            opacity += 0.1;
            element.style.opacity = opacity;
        } else {
            clearInterval(interval);
        }
    }, 70);
}

function fadeOut(element) {
    var opacity = 1;
    var interval = setInterval(function () {
        if (opacity > 0) {
            opacity -= 0.1;
            element.style.opacity = opacity;
        } else {
            clearInterval(interval);
        }
    }, 70);
}

setInterval(changeBackground, 5000);