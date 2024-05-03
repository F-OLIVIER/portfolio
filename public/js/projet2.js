const prevBtn = document.getElementById('prevBtn');
const nextBtn = document.getElementById('nextBtn');
const centerImage = document.querySelector('.center-image');
const images = [
    "public/img/projet2/startGame.jpg",
    "public/img/projet2/level1.png",
    "public/img/projet2/level2.png",
    "public/img/projet2/level3.png",
    "public/img/projet2/level4.png"
];
let currentIndex = 0;

prevBtn.addEventListener('click', () => {
    currentIndex = (currentIndex - 1 + images.length) % images.length;
    updateGallery();
});

nextBtn.addEventListener('click', () => {
    currentIndex = (currentIndex + 1) % images.length;
    updateGallery();
});

centerImage.addEventListener('click', () => {
    toggleEnlarge();
});

function updateGallery() {
    const newImage = document.createElement('img');
    newImage.src = images[currentIndex];
    centerImage.innerHTML = '';
    centerImage.appendChild(newImage);
}

function toggleEnlarge() {
    const image = centerImage.querySelector('img');
    image.classList.toggle('enlarged');
}