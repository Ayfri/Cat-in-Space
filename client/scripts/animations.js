function isVisibleInViewport(element) {
    if (element.offsetWidth || element.offsetHeight || element.getClientRects().length) {
        const rect = element.getBoundingClientRect();
        return rect.bottom > 0 && rect.right > 0 && rect.left < (window.innerWidth || document.documentElement.clientWidth) && rect.top < (window.innerHeight || document.documentElement.clientHeight);
    }
    return false;
}

export function listAppearing() {
    function animateVisibleElements() {
        const list = document.querySelectorAll('.streamer');

        list.forEach(item => {
            if (isVisibleInViewport(item)) item.classList.add('animate');
        });
    }

    document.addEventListener('scroll', animateVisibleElements);
    animateVisibleElements();
}