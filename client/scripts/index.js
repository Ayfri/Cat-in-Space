import {listAppearing} from './animations.js';
import {changeTitle} from './changeTitle.js';

function clickStreamers() {
    const hidden = document.querySelectorAll('.dream-hidden');
    const more = document.querySelector('.add-more');

    more.onclick = () => {
        const isHidden = more.classList.contains('hidden');
        hidden.forEach(item => item.style.display = isHidden ? 'flex' : 'none');
        more.textContent = isHidden ? 'hide...' : 'more...';
        more.classList.toggle('hidden');
    };
}

function cutString() {
    const divs = document.querySelectorAll('.streamer-bio');
    for (let i = 0; i < divs.length; i++) {
        if (divs[i].textContent.length > 60) {
            divs[i].textContent = `${divs[i].textContent.substring(0, 60)}...`;
        }
    }
}

function hideStreamers() {
    const dreamSmp = [...document.querySelectorAll('.dream-smp')];
    const hidden = dreamSmp.slice(-4);
    hidden.forEach(item => item.classList.add('dream-hidden'));
}


window.addEventListener('load', () => {
    const loader = document.querySelector('.loading');
    const waiting = 0.5;
    setTimeout(() => loader.classList.toggle('hide'), waiting * 1250);
});


window.onload = () => {
    cutString();
    hideStreamers();
    clickStreamers();
    changeTitle();
    listAppearing();
};