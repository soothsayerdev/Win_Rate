// Create a starfield animation
const canvas = document.createElement('canvas');
document.body.appendChild(canvas);
canvas.style.position = 'fixed';
canvas.style.top = '0';
canvas.style.left = '0';
canvas.style.zIndex = '-1';

const ctx = canvas.getContext('2d');
let stars = [];

function resizeCanvas() {
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;
}

resizeCanvas();
window.addEventListener('resize', resizeCanvas);

function createStar() {
    return {
        x: Math.random() * canvas.width,
        y: Math.random() * canvas.height,
        size: Math.random() * 2 + 1,
        speedX: Math.random() * 0.5 - 0.25,
        speedY: Math.random() * 0.5 - 0.25
    };
}

function drawStar(star) {
    ctx.beginPath();
    ctx.arc(star.x, star.y, star.size, 0, Math.PI * 2);
    ctx.fillStyle = '#ffffff';
    ctx.fill();
}

function updateStars() {
    for (let i = 0; i < stars.length; i++) {
        let star = stars[i];
        star.x += star.speedX;
        star.y += star.speedY;

        if (star.x > canvas.width) star.x = 0;
        if (star.y > canvas.height) star.y = 0;
        if (star.x < 0) star.x = canvas.width;
        if (star.y < 0) star.y = canvas.height;

        drawStar(star);
    }
}

function animate() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    if (stars.length < 200) {
        stars.push(createStar());
    }

    updateStars();
    requestAnimationFrame(animate);
}

animate();

// Sound effects for lightsabers (when hovering over buttons)
// const lightsaberSound = new Audio('https://www.soundjay.com/button/beep-07.wav');

document.querySelectorAll('.btn').forEach(button => {
    button.addEventListener('mouseenter', () => {
        lightsaberSound.play();
    });
});
