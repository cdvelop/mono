// convierte fecha 2021-03-11 a Lunes 11 de Marzo 2021
function DayWordChange(tag) {

    let fd = tag.closest('fieldset');
    // console.log("FD: ",fd);
    if (tag.dataset.date_target != "") {
        fd.classList.add("foka");
        
        tag.parentNode.dataset.spanish = DateToWords(tag.dataset.date_target,"long");
        // console.log("fecha traducida: ",daySelectedInWords);
    }else{
        fd.classList.remove("foka");
        tag.parentNode.dataset.spanish = ""
    }
};