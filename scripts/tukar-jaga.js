function tambahTukarJaga(){
    var payload = {
        "data01": document.getElementById("nama-penukar").value,
        "data02": document.getElementById("jenis-hutang-jaga").value,
        "data03": document.getElementById("tanggal-hutang-jaga").value,
        "data04": document.getElementById("jenis-bayar-jaga").value,
        "data05": document.getElementById("tanggal-bayar-jaga").value,
    }
    // console.log("Payload adalah: " + JSON.stringify(payload))
    document.getElementById("nama-penukar").value = ""
    document.getElementById("jenis-hutang-jaga").value = "1"
    document.getElementById("tanggal-hutang-jaga").value = ""
    document.getElementById("jenis-bayar-jaga").value = "1"
    document.getElementById("tanggal-bayar-jaga").value = ""
    sendPost("/tambah-tukar-jaga", JSON.stringify(payload), pageTukarJaga)
}

function updateTukarJaga(){
    document.getElementById("hal-tukar-jaga").click()
    // document.getElementById("refresh-list-tukar-jaga").click()
    // var js = JSON.parse(document.getElementById("server-response").innerHTML)
    // console.log("isi script adalah: " + js.script)
    // document.getElementById("daftar-tukar-jaga").innerHTML = js.script
    // document.getElementById("modal01-content").innerHTML = js.modal
    // document.getElementById("modal01").style.display = "block"
}
// function refreshListTukarJaga(){
//     // document.getElementById("header-list-tukar-jaga").click
//     servePage("/get-tukar-jaga", "daftar-tukar-jaga")
// }

function editDataJaga(){
    var link = this.dataset.link
    // console.log("Link adalah: " + link)
    var pilihan = this.value
    var ubah = "<select name='' id='ubah-jenis-bayar-jaga' class='w3-select w3-round'>"
    ubah = ubah + "<option value='1'>Pagi</option><option value='2'>Sore</option><option value='3'>Malam</option></select>"
    ubah = ubah + "<input type='date' class='w3-input w3-round' id='ubah-tanggal-bayar-jaga'>"
    switch (pilihan) {
        case "1":
        console.log("Hutang lunas")
        break;
        case "2":
        console.log("Hutang lunas sebagian")
        break;
        case "3":
        // console.log("Ubah tanggal bayar")
        document.getElementById("modal01-title-header").innerHTML = "Ubah Tanggal Bayar Jaga"
        document.getElementById("modal01-content").innerHTML = "Ubah tanggal jaga menjadi:"
        document.getElementById("modal01-content-02").innerHTML = ubah
        document.getElementById("modal01-tombol-tambahan").dataset.link = link
        document.getElementById("modal01-tombol-tambahan").innerHTML = "Ubah"
        document.getElementById("modal01-tombol-tambahan").style.display = "block"
        document.getElementById("modal01").style.display = "block"
        document.getElementById("modal01-tombol-tambahan").addEventListener("click", ubahTanggalBayar)
        break;
        default:
        console.log("Tidak memilih")
    }
}

function ubahTanggalBayar(){
    var payload = {
        "data01" : document.getElementById("ubah-jenis-bayar-jaga").value,
        "data02" : document.getElementById("ubah-tanggal-bayar-jaga").value,
        "data03" : document.getElementById("modal01-tombol-tambahan").dataset.link
    }
    document.getElementById("modal01-tombol-tambahan").style.display = "none"
    document.getElementById("modal01-content-02").innerHTML = ""
    sendPost("/ubah-tanggal-bayar-jaga", JSON.stringify(payload), pageTukarJaga)
    // console.log("payload adalah : " + JSON.stringify(payload))
}
function preHapusDataTukarJaga(){
    var link = this.dataset.link
    console.log("Link adalah: " + link)
}

function tambahTanggalBayarJaga(){
    // var link = this.dataset.link
    var payload = {
        "data01": this.dataset.link,
        "data02": document.getElementById("tambah-jenis-bayar-jaga").value,
        "data03": this.value
    }
    sendPost("/tambah-bayar-jaga", JSON.stringify(payload), pageTukarJaga)
    // console.log("Payload adalah: " + JSON.stringify(payload))
}