const prevBtn = document.getElementById('prevBtn');
const nextBtn = document.getElementById('nextBtn');
const centerImage = document.querySelector('.center-image');
const images = [
    "public/img/projet1/Schema_architecture.png",
    "public/img/projet1/Schema_Database.png",
    "public/img/projet1/bot_slashcommand.png",
    "public/img/projet1/bot_commanddata.png",
    "public/img/projet1/bot_message_gvg_on.png",
    "public/img/projet1/bot_message_gvg_off.png",
    "public/img/projet1/site_home.png",
    "public/img/projet1/site_dashboard.png",
    "public/img/projet1/site_charactercard.png",
    "public/img/projet1/site_groupgvg.png",
    "public/img/projet1/site_caserne.png",
    "public/img/projet1/flutter_01.png",
    "public/img/projet1/flutter_04.png",
    "public/img/projet1/flutter_02.png",
    "public/img/projet1/flutter_03.png",
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
