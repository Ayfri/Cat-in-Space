changeTitle = () => {
    const title = document.querySelector("title");
    const name = document.querySelectorAll(".test-name");
    title.innerText = String(name[0].innerText);
}

changeTitle()
