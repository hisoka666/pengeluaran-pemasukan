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
    // console.log("button fired!")
    var link = this.dataset.link
    var tgl = document.getElementById("pemasukan-bulan").value
    var payload = {"link": link, "nama": tgl}
    document.getElementById("modal01-title-header").innerHTML = "Peringatan"
    document.getElementById("modal01-content").innerHTML = "Yakin ingin menghapus entri ini?"
    var but =  document.getElementById("modal01-tombol-tambahan")
    // console.log("Tombol adalah: " + JSON.stringify(payload))
    but.style.display = "block"
    but.dataset.link = JSON.stringify(payload)
    but.addEventListener("click", deleteIncome)
    document.getElementById("modal01").style.display = "block"
}
function deleteIncome(){
    sendPost("/delete-pemasukan", document.getElementById("modal01-tombol-tambahan").dataset.link, berhasilMenghapusPemasukan)
    
}

function berhasilMenghapusPemasukan(){
    var js = JSON.parse(document.getElementById("server-response").innerHTML)
    document.getElementById("modal01-content").innerHTML = js.script
    document.getElementById("modal01-tombol-tambahan").style.display = "none"
    document.getElementById("modal01").style.display = "block"
    document.getElementById("modal01-tombol-tutup").addEventListener("click", updatePemasukanDelete)
}

function updatePemasukanDelete(){
    // console.log("Button fired")
    var tgl = document.getElementById("pemasukan-bulan").value
    // console.log("Tgl adalah : " + tgl)
    if (tgl == ""){
        // console.log("Tanggal kosong")
        servePage("/pemasukan", "main-content")
    }else {
        // console.log("Tanggal isi")
        getPemasukanBulan()
    }
    // var payload = {"nama": tgl}
    // sendPost("/get-pemasukan-bulan", JSON.stringify(payload), updatePemasukan)
    // // console.log(document.getElementById("server-response").innerHTML)
    // var js = JSON.parse(document.getElementById("server-response").innerHTML)
    // document.getElementById("pemasukan-content").innerHTML = js.script
    // document.getElementById("pemasukan-tiap-bulan").innerHTML = js.tambahan

}