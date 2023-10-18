let ImagenesEnCrear = 1;
let ImagenesEnCrearup = 1;
var ultimaRespuesta = {};
$(document).ready(function() {
    consultamasiva();
});

function plusimageCrear() {
    DivImagenextra = document.getElementById('ImagenesCrearn');
    NuevoDIV = document.createElement("DIV");
    NuevoDIV.setAttribute("class", "inputer");
    NuevoTexto = document.createTextNode('Url imagen ' + ImagenesEnCrear + ' :');
    Nuevoinput = document.createElement("input");
    Nuevoinput.setAttribute("id", "IPImageCreate_" + ImagenesEnCrear);
    ImagenesEnCrear++;
    NuevoDIV.appendChild(NuevoTexto)
    NuevoDIV.appendChild(Nuevoinput)
    DivImagenextra.appendChild(NuevoDIV);
}

function plusimageCrearup() {
    DivImagenextra = document.getElementById('imagenesUpdate');
    NuevoDIV = document.createElement("DIV");
    NuevoDIV.setAttribute("class", "inputer");
    NuevoTexto = document.createTextNode('Url imagen ' + ImagenesEnCrearup + ' :');
    Nuevoinput = document.createElement("input");
    Nuevoinput.setAttribute("id", "IPImageCreateup_" + ImagenesEnCrearup);
    ImagenesEnCrearup++;
    NuevoDIV.appendChild(NuevoTexto)
    NuevoDIV.appendChild(Nuevoinput)
    DivImagenextra.appendChild(NuevoDIV);
}

function crearProducto() {
    var data = {
        "SKU": document.getElementById("IPSKUCreate").value,
        "name": document.getElementById("IPNameCreate").value,
        "brand": document.getElementById("IPBrandCreate").value,
        "size": document.getElementById("IPSizeCreate").value,
        "price": document.getElementById("IPPriceCreate").value,
        "imagenes": []
    };
    var imagenes = [];
    for (i = 0; i < ImagenesEnCrear; i++) {
        var k = document.getElementById("IPImageCreate_" + i).value;
        imagenes.push({
            "url": k,
        });
    }
    data.imagenes = imagenes;
    //console.log(data);
    var s = new sllamado('Create');
    s.consultar(data, "PUT", true);

}

function consultamasiva() {
    var s = new sllamado('Read');
    s.consultar(null, "GET");
    setTimeout(() => { armaTabla() }, 5000);
}

async function armaTabla() {
    if (ultimaRespuesta[1] == undefined) {
        consultamasiva();
        return;
    }
    DivImagenextra = document.getElementById('mass');
    Nuevotab = document.createElement("table");
    Nuevotit = document.createElement("tr");
    NuevoCel1 = document.createElement("td");
    NuevoCel2 = document.createElement("td");
    NuevoCel3 = document.createElement("td");
    NuevoCel4 = document.createElement("td");
    NuevoCel5 = document.createElement("td");
    NuevoCel6 = document.createElement("td");
    NuevoCel7 = document.createElement("td");
    NuevoCel1.innerHTML = "ID"
    NuevoCel2.innerHTML = "SKU"
    NuevoCel3.innerHTML = "NAME"
    NuevoCel4.innerHTML = "BRAND"
    NuevoCel5.innerHTML = "SIZE"
    NuevoCel6.innerHTML = "PRICE"
    NuevoCel7.innerHTML = "IMAGES"
    Nuevotit.appendChild(NuevoCel1);
    Nuevotit.appendChild(NuevoCel2);
    Nuevotit.appendChild(NuevoCel3);
    Nuevotit.appendChild(NuevoCel4);
    Nuevotit.appendChild(NuevoCel5);
    Nuevotit.appendChild(NuevoCel6);
    Nuevotit.appendChild(NuevoCel7);
    Nuevotab.appendChild(Nuevotit);
    for (d = 0; d < +ultimaRespuesta.length; d++) {
        Nuevoalinea = document.createElement("tr");
        NuevoCel1 = document.createElement("td");
        NuevoCel2 = document.createElement("td");
        NuevoCel3 = document.createElement("td");
        NuevoCel4 = document.createElement("td");
        NuevoCel5 = document.createElement("td");
        NuevoCel6 = document.createElement("td");
        NuevoCel7 = document.createElement("td");
        NuevoCel1.innerHTML = ultimaRespuesta[d].identificador
        NuevoCel2.innerHTML = ultimaRespuesta[d].SKU
        NuevoCel3.innerHTML = ultimaRespuesta[d].name
        NuevoCel4.innerHTML = ultimaRespuesta[d].brand
        NuevoCel5.innerHTML = ultimaRespuesta[d].size
        NuevoCel6.innerHTML = ultimaRespuesta[d].price
        Nuevoalinea.appendChild(NuevoCel1);
        Nuevoalinea.appendChild(NuevoCel2);
        Nuevoalinea.appendChild(NuevoCel3);
        Nuevoalinea.appendChild(NuevoCel4);
        Nuevoalinea.appendChild(NuevoCel5);
        Nuevoalinea.appendChild(NuevoCel6);
        for (j = 0; j < ultimaRespuesta[d].imagenes.length; j++) {
            NuevoCel7.innerHTML = NuevoCel7.innerHTML + ultimaRespuesta[d].imagenes[j].url + '<br>'
        }
        Nuevotab.appendChild(Nuevoalinea);
        Nuevoalinea.appendChild(NuevoCel7);
    }
    DivImagenextra.appendChild(Nuevotab);
}

function TraeProducto() {
    var data = {
        "SKU": document.getElementById("IPSKUASK").value,
    };
    var s = new sConsul("Ask");
    s.consultar(data, "POST", true);

}

function MODProducto() {
    var data = {
        "SKU": document.getElementById("IPSKUUpdate").value,
        "name": document.getElementById("IPNameUpdate").value,
        "brand": document.getElementById("IPBrandUpdate").value,
        "size": document.getElementById("IPSizeUpdate").value,
        "price": document.getElementById("IPPriceUpdate").value,
        "imagenes": []
    };
    var imagenes = [];
    for (i = 1; i < +ImagenesEnCrearup; i++) {
        var k = document.getElementById("IPImageCreateup_" + i).value;
        imagenes.push({
            "url": k,
        });
    }
    data.imagenes = imagenes;
    //console.log(data);
    var s = new sllamado('Update');
    s.consultar(data, "POST", true);

}

function DelProducto() {
    var data = {
        "SKU": document.getElementById("IPSKUDelete").value,
    };
    var s = new sConsul("Delete");
    s.consultar(data, "DELETE", true);
    location.reload()
}


function sConsul(link) {
    this.source = 'http://localhost:18080/' + link;
    this.callback = null;
    this.extra = null;
    this.consultar = function(datas, metod, c) {
        var data = ""
        var that = this;
        $.ajax({
            url: this.source,
            data: JSON.stringify(datas),
            method: metod,
            success: function(data) {

                Respuesta = JSON.parse(data);
                document.getElementById("IPSKUASK").value = Respuesta.SKU;
                document.getElementById("IPNameASK").value = Respuesta.name;
                document.getElementById("IPBrandASK").value = Respuesta.brand;
                document.getElementById("IPSizeASK").value = Respuesta.size;
                document.getElementById("IPPriceASK").value = Respuesta.price;
                document.getElementById("IPSKUUpdate").value = Respuesta.SKU;
                document.getElementById("IPNameUpdate").value = Respuesta.name;
                document.getElementById("IPBrandUpdate").value = Respuesta.brand;
                document.getElementById("IPSizeUpdate").value = Respuesta.size;
                document.getElementById("IPPriceUpdate").value = Respuesta.price;
                document.getElementById('imagenesASK').innerHTML = 'URL imagenes';
                ImagenesEnCrearup = 1;
                document.getElementById('imagenesUpdate').innerHTML = 'URL imagenes';
                ImagenesEnCrear = 1;
                for (j = 0; j < Respuesta.imagenes.length; j++) {
                    document.getElementById('imagenesASK').innerHTML = document.getElementById('imagenesASK').innerHTML + '<br>' + Respuesta.imagenes[j].url + '<br>'
                    plusimageCrearup()
                    document.getElementById('IPImageCreateup_' + (j + 1)).value = Respuesta.imagenes[j].url

                }
            },
            error: function(data) {
                //alert("comando fallo")
                if (data.estado != undefined) {
                    alert(data.estado)
                }

            },
            async: true
        });
    };
}

function sllamado(link) {
    this.source = 'http://localhost:18080/' + link;
    this.callback = null;
    this.extra = null;
    this.consultar = function(datas, metod, c) {
        var data = ""
        var that = this;
        $.ajax({
            url: this.source,
            data: JSON.stringify(datas),
            method: metod,
            success: function(data) {
                ultimaRespuesta = JSON.parse(data);
            },
            error: function(data) {
                //alert("comando fallo")
                if (data.estado != undefined) {
                    alert(data.estado)
                }

            },
            async: true
        });
    };
}
//document.getElementById("").value