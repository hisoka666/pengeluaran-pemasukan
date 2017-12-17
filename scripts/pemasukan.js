function getPemasukanBulan() {
    document.getElementById("loading-animation").style.display = "block"
    var tgl = document.getElementById("pemasukan-bulan").value
    // console.log("Tgl adalah : " + tgl)
    var payload = {"nama": tgl}
    sendPost("/get-pemasukan-bulan", JSON.stringify(payload), updatePemasukan)
}

function updatePemasukan(){
    var js = JSON.parse(document.getElementById("server-response").innerHTML)
    document.getElementById("pemasukan-tiap-bulan").innerHTML =  js.tambahan
    document.getElementById("pemasukan-content").innerHTML = js.script
    document.getElementById("loading-animation").style.display = "none"
}

function deleteIncomeBut(){
    console.log("button fired!")
    var link = this.dataset.link
    var tgl = document.getElementById("pemasukan-bulan").value
    var payload = {"link": link, "nama": tgl}
    document.getElementById("modal01-title-header").innerHTML = "Peringatan"
    document.getElementById("modal01-content").innerHTML = "Yakin ingin menghapus entri ini?"
    var but =  document.getElementById("modal01-tombol-tambahan")
    // console.log("Tombol adalah: " + but.innerHTML)
    but.style.display = "block"
    but.dataset.link = JSON.stringify(payload)
    but.onclick = deleteIncome
    document.getElementById("modal01").style.display = "block"
}
function deleteIncome(){
    sendPost("/delete-pemasukan", JSON.stringify(document.getElementById("modal01-tombol-tambahan").dataset.link), updatePemasukanTanggal)
}

function updatePemasukanTanggal(){
    var js = JSON.parse(document.getElementById("server-response").innerHTML)
    document.getElementById("pemasukan-content").innerHTML = js.script
    document.getElementById("pemasukan-tiap bulan").innerHTML = js.tambahan

}