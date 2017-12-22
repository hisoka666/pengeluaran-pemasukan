function getJenisPiutang(){
    if (this.value == "1"){
        document.getElementById("tanggal-piutang").style.display = "block"
        document.getElementById("jumlah-piutang").style.display = "none"
        document.getElementById("pilihan-piutang").style.display = "block"
    } if (this.value == "2") {
        document.getElementById("jumlah-piutang").style.display = "block"
        document.getElementById("tanggal-piutang").style.display = "none"
        document.getElementById("pilihan-piutang").style.display = "block"
    }
}

function tambahPiutang(){
    var payload = {
        "namajenis":document.getElementById("nama-jenis").value,
        "tglpiu":document.getElementById("tanggal-jaga").value,
        "jml":document.getElementById("jumlah").value
    }
    console.log("Payload adalah: " + JSON.stringify(payload))
    sendPost("/tambah-piutang", JSON.stringify(payload), updatePiutang)
}

function updatePiutang(){
    servePage("/get-piutang", "daftar-piutang")
}