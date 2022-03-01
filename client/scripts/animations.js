export function listAppearing() {
    const elements = document.querySelectorAll('.streamer');
    const options = {
        root: null,
        rootMargin: '0px',
        threshold: 0.1
    };
    const observer = new IntersectionObserver(function(entries, observer) {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.classList.add('appear');
                observer.unobserve(entry.target);
            }
        });
    }, options);
    elements.forEach(element => {
        observer.observe(element);
    });
}