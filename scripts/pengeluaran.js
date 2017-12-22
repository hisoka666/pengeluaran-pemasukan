function pagePengeluaranRefresh(){
    servePage("/pengeluaran-refresh", "pengeluaran-content")
    // console.log("this clicked")
}

function getPerbulan() {
    var kur = this.getAttribute("data-link-kursor")
    var tgl = this.innerHTML
    var payload = {"link" : kur, "nama": tgl}
    sendPost("/getbulan", JSON.stringify(payload), displayPerbulan)
}
function displayPerbulan() {
    js = JSON.parse(document.getElementById("server-response").innerHTML)
    document.getElementById("pengeluaran-content").innerHTML = js.script
    document.getElementById("server-response").innerHTML = ""
}

function deletePengeluaranBut() {
    var link = this.dataset.link
    var but = document.getElementById("modal01-tombol-tambahan")
    document.getElementById("modal01-title-header").innerHTML = "Peringatan"
    document.getElementById("modal01-content").innerHTML = "Yakin ingin menghapus entri ini?"
    but.innerHTML = "Hapus"
    but.dataset.link = link
    but.dataset.index = this.index
    but.onclick = deletePengeluaran
    but.style.display="block"
    document.getElementById("modal01").style.display="block"
}
function deletePengeluaran() {
    var payload = {"link": document.getElementById("modal01-tombol-tambahan").dataset.link}
    document.getElementById("modal01-tombol-tambahan").style.display = "none"
    sendPost("/delete-pengeluaran", JSON.stringify(payload), berhasilMenghapus)
}
function berhasilMenghapus(){
    document.getElementById("modal01-content").innerHTML = document.getElementById("server-response").innerHTML
    document.getElementById("modal01").style.display = "block"
    document.getElementById("modal01-tombol-tutup").addEventListener("click", pagePengeluaranRefresh)
}
function updateListPengeluaran(){
    var js = JSON.parse(document.getElementById("server-response").innerHTML)

}