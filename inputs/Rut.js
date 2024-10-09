function changeDniType(e) {
    const input_dni = e.closest('.run-type').querySelector('input[type="text"][data-name="rutDni"]');

    if (e.value === "ch") {
        input_dni.setAttribute("maxlength", 10);
        input_dni.dataset.option = "ch";

    } else {
        input_dni.setAttribute("maxlength", 15);
        input_dni.dataset.option = "ex";
    }
    userFormTyping(input_dni);
};

function RunToPointFormat(rut) {
    // XX.XXX.XXX-X
    let run_number = rut.substring(0, rut.length - 2)
    let run_point = FormateaNumero(run_number);
    let _dv = rut.substring(rut.length - 2, rut.length);

    return run_point + _dv
};