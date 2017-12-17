function pagePengeluaranRefresh(){
    servePage("/pengeluaran-refresh", "pengeluaran-content")
    // console.log("this clicked")
}

function getPerbulan() {
    var kur = this.getAttribute("data-link-kursor")
    var tgl = this.innerHTML
    // console.log("Kursor adalah: " + kur)
    // console.log("Tanggal adalah: " + tgl)
    var payload = {"link" : kur, "nama": tgl}
    // console.log("Payload adalah: " + payload.link)
    sendPost("/getbulan", JSON.stringify(payload), displayPerbulan)
}
function displayPerbulan() {
    // console.log(document.getElementById("server-response").innerHTML)
    js = JSON.parse(document.getElementById("server-response").innerHTML)
    document.getElementById("pengeluaran-content").innerHTML = js.script
    // var jml = parseInt(js.data.total)
    // document.getElementById("total-pengeluaran").innerHTML = "Rp " + toLocaleString(jml)
    document.getElementById("server-response").innerHTML = ""
}

function deletePemasukanBut() {
    var link = this.dataset.link
    // console.log("Link adalah: " + link)
    var but = document.getElementById("modal01-tombol-tambahan")
    // var index = document.getElementsByClassName("table-item-pengeluaran").findIndex(this)
    // console.log("Index adalah: " + index)
    // var payload = {"link": this.dataset.link}
    document.getElementById("modal01-title-header").innerHTML = "Peringatan"
    document.getElementById("modal01-content").innerHTML = "Yakin ingin menghapus entri ini?"
    but.innerHTML = "Hapus"
    but.dataset.link = link
    // console.log("Index adalah: " + this.findIndex)
    but.dataset.index = this.index
    but.onclick = deletePemasukan
    but.style.display="block"
    document.getElementById("modal01").style.display="block"
    // sendPost("/delete-pengeluaran", payload, updateListPengeluaran)
}
function deletePemasukan() {
    // console.log("button fired")
    var payload = {"link": document.getElementById("modal01-tombol-tambahan").dataset.link}
    // console.log("Link adalah : " + JSON.stringify(payload))
    document.getElementById("modal01-tombol-tambahan").style.display = "none"
    // console.log("Link adalah: " + payload.link)
    sendPost("/delete-pengeluaran", JSON.stringify(payload), berhasilMenghapus)
    // servePage("/pengeluaran-refresh", "pengeluaran-content")    
    // document.getElementById("pengeluaran-tombol-bulan-ini").click()
}
function berhasilMenghapus(){
    // document.getElementById("pengeluaran-tombol-bulan-ini").click()
    // pagePengeluaranRefresh()
    document.getElementById("modal01-content").innerHTML = document.getElementById("server-response").innerHTML
    document.getElementById("modal01").style.display = "block"
    document.getElementById("modal01-tombol-tutup").addEventListener("click", pagePengeluaranRefresh)
    // servePage("/pengeluaran-refresh", "pengeluaran-content")
}
function updateListPengeluaran(){
    var js = JSON.parse(document.getElementById("server-response").innerHTML)

}