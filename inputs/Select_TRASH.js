// agrega un elemento option a todos los select del documento según nombre
function addOptSelect(name, value) {
    let selects = document.querySelectorAll("select[name='" + name + "']");
    // console.log("selects: ", selects);
    for (let s = 0; s < selects.length; s++) {
        selects[s].insertAdjacentHTML("beforeend", value);
    }
};