function myFunc(){
    var xhttp = new XMLHttpRequest()
    xhttp.onreadystatechange = function (){
        if (this.readyState == 4 && this.status == 200) {
            // console.log(this.responseText)
            var js = JSON.parse(this.responseText)
            // console.log(js.data.toLocaleString())
            document.getElementById("main-content").innerHTML = js.script
            // document.getElementById("balance-sisa").innerHTML = "Rp " + js.data.toLocaleString()
            // document.getElementById("balance-hidden").value = js.data
            // console.log(document.getElementById("balance-hidden").value)
            // stopReloading()
            document.getElementById("but-pemasukan").addEventListener("click", function(event){
                event.preventDefault()
                inputPemasukan()
            })
            document.getElementById("but-pengeluaran").addEventListener("click", function(event){
                event.preventDefault()
                inputPengeluaran()
            })
        }
    }
    xhttp.open("GET", "/frontpage", true)
    // xhttp.open("GET", "templates/main-content.html?t=" + Math.random(), true)
    xhttp.send()
}
function stopReloading() {
    var list = document.body.getElementsByTagName("a")
    console.log("panjang list adalah: " + list.length)
    for (i=0;i<list.length-1;i++) {
        console.log("list no: " + i)
        console.log("Isi dari list adalah: " + list[i].innerHTML)
        list[i].addEventListener("click", function(e){
            e.preventDefault()
            console.log("nomer list adalah : " + i)
            console.log("Isi dari list adalah: " + list[i].innerHTML)
        })
    }
}
function inputPemasukan(){
    document.getElementById("loading-animation").style.display="block"
    var nama = document.getElementById("input-nama").value
    var jumlah = document.getElementById("input-jumlah").value
    document.getElementById("input-nama").value = ""
    document.getElementById("input-jumlah").value = ""
    payload = {"nama": nama, "pemasukan": jumlah}
    sendPost("/add-pemasukan", JSON.stringify(payload), pemasukanFunction)
    document.getElementById("modal01").style.display="block"
}
function inputPengeluaran(){
    document.getElementById("loading-animation").style.display="block"
    var nama = document.getElementById("input-nama").value
    var jumlah = document.getElementById("input-jumlah").value
    document.getElementById("input-nama").value = ""
    document.getElementById("input-jumlah").value = ""
    payload = {"nama": nama, "pengeluaran": jumlah}
    sendPost("/add-pengeluaran", JSON.stringify(payload), updateBalance)
}
function pemasukanFunction() {
    js = JSON.parse(document.getElementById("server-response").innerHTML)
    document.getElementById("modal01-content").innerHTML = js.servermsg
    document.getElementById("server-response").innerHTML = ""
    document.getElementById("loading-animation").style.display="none"
}
function updateBalance(){
    js = JSON.parse(document.getElementById("server-response").innerHTML)
    document.getElementById("balance-sisa").innerHTML = js.pemasukan
    document.getElementById("modal01-content").innerHTML = js.servermsg
    document.getElementById("loading-animation").style.display = "none"
    document.getElementById("modal01").style.display="block"
    document.getElementById("balance-hidden").value = js.pengeluaran
    document.getElementById("server-response").innerHTML = ""
}
function createServConn (targetId){
    var xhttp = new XMLHttpRequest()
    xhttp.onreadystatechange = function (){
        if (this.readyState == 4 && this.status == 200) {
            document.getElementById(targetId).innerHTML = this.responseText
            document.getElementById("loading-animation").style.display="none"
        }
    }

    return xhttp
}

function sendPost(url, content, somefunc){
    document.getElementById("loading-animation").style.display="block"
    var xhttp = new XMLHttpRequest()
    xhttp.onreadystatechange = function (){
        if (this.readyState == 4 && this.status == 200) {
            document.getElementById("server-response").innerHTML = this.responseText
            document.getElementById("loading-animation").style.display="none"
            somefunc()
        } else if (this.readyState == 4 && this.status != 200) {
            console.log(this.responseText)
            document.getElementById("loading-animation").style.display="none"
            document.getElementById("modal01-title-header").innerHTML = "Peringatan!"
            document.getElementById("modal01-content").innerHTML = this.responseText
            document.getElementById("modal01").style.display="block"
        }
    }
    xhttp.open("POST", url, true)
    xhttp.send(content)
}
function servePage(url, target){
    document.getElementById("loading-animation").style.display="block"
    var xhttp = createServConn(target)
    xhttp.open("GET", url, true)
    xhttp.send()
}
function pagePengeluaran(){
    servePage("/pengeluaran", "main-content")
    // console.log("Jumalh element " + document.getElementById("pengeluaran-content").childElementCount())
}
function changeToRupiah() {
    var jml = this.innerHTML
    console.log("Text adalah: " + jml)
    this.innerHTML = "Rp " + toLocaleString(parseInt(jml))
}
function pagePemasukan(){
    servePage("/pemasukan", "main-content")
}

function pageHutang(){
    servePage("/hutang", "main-content")
}

function pagePiutang(){
    servePage("/piutang", "main-content")
}

function pagePengaturan(){
    servePage("/pengaturan", "main-content")
}