const valRun = {
    // Valida el rut con su cadena completa "XXXXXXXX-X"
    validaRut: function (rutCompleto) {
        if (!/^[0-9]+[-|‐]{1}[0-9kK]{1}$/.test(rutCompleto))
            return false;
        let tmp = rutCompleto.split('-');
        let digv = tmp[1];
        let rut = tmp[0];
        if (digv == 'K') digv = 'k';
        return (valRun.dv(rut) == digv);
    },
    dv: function (T) {
        let M = 0, S = 1;
        for (; T; T = Math.floor(T / 10))
            S = (S + T % 10 * (9 - M++ % 6)) % 11;
        return S ? S - 1 : 'k';
    }
};

function FieldRutChanged(input) {
    // console.log("VALIDANDO RUN: ", input);
    if (valRun.validaRut(input.value)) {
        inputRight(input);
        return true
    } else {
        inputWrong(input);
        return false
    };
};

function ValidateForeignDoc(input) {
    const rex = /^[A-Za-z0-9]{9,15}$/;
    if (rex.test(input.value)) {
        inputRight(input);
        return true
    } else {
        inputWrong(input);
        return false
    }
};

function FieldDniChanged(e) {
    const document_type = e.closest('.run-type').querySelector('input[type="radio"][name="type-dni"]:checked');

    console.log("CAMBIO EN DNI: ", input.value, "TIPO: ", document_type);

    if (ValidateDni(document_type, input)) {
        reportCorrectInput(input)
    } else {
        reportWrongInput(input)
    }
};

function ValidateDni(document_type, input) {
    if (input.value === "") {
        inputRight(input);
        return false
    }

    try {
        if (document_type === "ch") {
            return FieldRutChanged(input);
        } else {
            return ValidateForeignDoc(input);
        }

    } catch (e) {
        return FieldRutChanged(input);
    };
};




