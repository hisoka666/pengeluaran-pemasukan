function tambahHutang(){
    var payload = {
        "data01": document.getElementById("nama-penghutang").value,
        "data02": document.getElementById("jenis-hutang").value,
        "data03": document.getElementById("jumlah").value
    }
    // console.log("payload adalah: " + JSON.stringify(payload))
    sendPost("/tambah-hutang", JSON.stringify(payload), updateHutang)
}

function updateHutang(){
    alert("Update hutang")
}