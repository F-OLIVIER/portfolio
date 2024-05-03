const prevBtn = document.getElementById('prevBtn');
const nextBtn = document.getElementById('nextBtn');
const centerImage = document.querySelector('.center-image');
const images = [
    "public/img/projet1/data.png",
    "public/img/projet1/listcommand.png",
    "public/img/projet1/Schema_DatabaseGvG.png",
    "public/img/projet1/connexion.png",
    "public/img/projet1/home.png",
    "public/img/projet1/charactercard.png",
    "public/img/projet1/caserne.png",
    "public/img/projet1/creategroup.png",
    "public/img/projet1/administration.png"
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
