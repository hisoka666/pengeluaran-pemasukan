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
    sendPost("/tambah-tukar-jaga", JSON.stringify(payload), updateTukarJaga)
}

function updateTukarJaga(){
    // document.getElementById("refresh-list-tukar-jaga").click()
    var js = JSON.parse(document.getElementById("server-response").innerHTML)
    console.log("isi script adalah: " + js.script)
    document.getElementById("daftar-tukar-jaga").innerHTML = js.script
    document.getElementById("modal01-content").innerHTML = js.modal
    document.getElementById("modal01").style.display = "block"
}
// function refreshListTukarJaga(){
//     // document.getElementById("header-list-tukar-jaga").click
//     servePage("/get-tukar-jaga", "daftar-tukar-jaga")
// }