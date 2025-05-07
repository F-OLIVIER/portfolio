const prevBtn = document.getElementById('prevBtn');
const nextBtn = document.getElementById('nextBtn');
const centerImage = document.querySelector('.center-image');
const images = [
    "public/img/projet3/home.png",
    "public/img/projet3/post.png",
    "public/img/projet3/chat.png",
    "public/img/projet3/group.png",
    "public/img/projet3/notif.png",
    "public/img/projet3/database.png",
    "public/img/projet3/img_app.png",
    "public/img/projet3/img_app_offlinemode.png",
    "public/img/projet3/crossplatform.png"
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
